package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	p := ("./website/index.html")
	// set header
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":

		r.ParseForm()
		uname := r.FormValue("uname")
		pwd := r.FormValue("pwd")

		fmt.Printf(uname + pwd)
		if uname == "admin" && pwd == "root" {
			http.Redirect(w, r, "https://www.google.com", http.StatusFound)
		} else {
			http.Redirect(w, r, "https://wikipedia.com", http.StatusFound)
		}

	case "GET":
		p := ("./website/portal.html")
		http.ServeFile(w, r, p)
	}
}

func main() {

	r := mux.NewRouter().StrictSlash(false)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./website/")))

	r.HandleFunc("/", GetIndexHandler).Methods("GET")
	r.HandleFunc("/portal.html", FormHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
