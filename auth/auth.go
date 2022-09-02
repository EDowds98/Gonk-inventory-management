package auth

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"webserver/sessions"
	"webserver/templates"

	"github.com/joho/godotenv"
)

/* This function reads the form data from the portal login page and
 * redirects depending on the values
 */
func FormHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := sessions.Store.Get(r, "session")

	// global variable needed to declare template variable
	var tmpl *template.Template

	tmpl, _ = template.ParseFiles("./website/template.html")
	// load env variables (will later be from db)
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("problem with loading environment variables")
	}

	fmt.Println(session.Values["loggedin"])
	if session.Values["loggedin"] != "true" {
		if r.Method == "POST" {

			uname := os.Getenv("PWORD")
			pwd := os.Getenv("UNAME")
			fmt.Println("In post")
			if uname == r.FormValue("uname") && pwd == r.FormValue("pwd") {
				session.Values["loggedin"] = "true"
				session.Save(r, w)
				fmt.Println("just before redirect")
				http.Redirect(w, r, "www.google.com", http.StatusTemporaryRedirect)
				return
			}

		} else if r.Method == "GET" {
			fmt.Println("GET case")
			tmpl.Execute(w, templates.LoginContent)
			/*	if err != nil {
				log.Fatalf("templating failed in FormHandler method: %v", err)
			} */
		} else {
			fmt.Println("not logged in")
			http.Redirect(w, r, "www.youtube.com", http.StatusTemporaryRedirect)
		}

	}
}
