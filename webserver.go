/*
 * Written by Eric Dowds ed67@hw.ac.uk
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/* This is needed instead of a static port because heroku
 * randomises ports when it starts
 */
func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		//return "", fmt.Errorf("$PORT not set")
		port = "9000"
	}
	return ":" + port, nil
}

/* This function reads the form data from the portal login page and
 * redirects depending on the values
 */
func FormHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":
		uname := r.FormValue("uname")
		pwd := r.FormValue("pwd")

		if uname == "admin" && pwd == "root" {
			fmt.Println("just before redirect")
			http.Redirect(w, r, "/login-success", http.StatusFound)
		} else {
			http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusFound)
		}

	case "GET":
		fmt.Println("GET case")
		p := "./website/login-success.html"
		http.ServeFile(w, r, p)

	}
}

func FormPresenter(w http.ResponseWriter, r *http.Request) {
	p := ("./website/portal.html")
	fmt.Println("at formPresenter")
	http.ServeFile(w, r, p)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := ("./website/index.html")
	http.ServeFile(w, r, p)
}

func main() {

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/login-success", FormHandler)
	http.HandleFunc("/portal", FormPresenter)

	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
