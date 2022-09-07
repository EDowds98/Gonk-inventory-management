package auth

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"webserver/templates"

	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
)

// make a new cookie
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// helper function for checking cookies
func getUserName(req *http.Request) (userName string) {
	if cookie, err := req.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return
}

func setSession(userName string, resp http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/", //come back to this
		}
		http.SetCookie(resp, cookie)
	}
}

/*func clearSession(resp http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/", //come back to this
		MaxAge: -1,
	}
	http.SetCookie(resp, cookie)
} */

// global variable needed to declare template variable
var tmpl *template.Template

/* This function reads the form data from the portal login page and
 * redirects depending on the values
 */
func FormHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, _ = template.ParseFiles("../website/template.html")
	// load env variables (will later be from db)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("problem with loading environment variables, %s", err)
	}

	if r.Method == "POST" {

		uname := os.Getenv("UNAME")
		pwd := os.Getenv("PWORD")
		fmt.Println("In post")
		redirectTarget := "/portal"
		if uname == r.FormValue("uname") && pwd == r.FormValue("pwd") {
			fmt.Println("just before redirect")
			setSession(uname, w)
			redirectTarget = "/internal"
		}
		http.Redirect(w, r, redirectTarget, http.StatusTemporaryRedirect)
	}

}

func InternalHandler(resp http.ResponseWriter, req *http.Request) {
	userName := getUserName(req)

	if userName == os.Getenv("UNAME") {
		fmt.Println("here")
		if err := tmpl.Execute(resp, templates.LoginContent); err != nil {
			log.Fatalf("internal portal templating error, %s", err)
		}
	}

}
