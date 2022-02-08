package main

import (
	"fmt"
	"log"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":
		uname := r.FormValue("uname")
		pwd := r.FormValue("pwd")

		if uname == "admin" && pwd == "root" {
			http.Redirect(w, r, "./website/login-success.html", http.StatusFound)
		} else {
			http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusFound)
		}

	case "GET":
		p := ("./website/login-success.html")
		http.ServeFile(w, r, p)
	}
}

func FormPresenter(w http.ResponseWriter, r *http.Request) {

	p := ("login-success.html")
	http.ServeFile(w, r, p)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w)
}

func main() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/login-success", FormHandler)
	http.HandleFunc("/portal", FormPresenter)

	css := http.FileServer(http.Dir("./website"))
	http.Handle("/website/", http.StripPrefix("/website/", css))
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
