package main

import (
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
	case "GET":
		p := ("./website/portal.html")
		http.ServeFile(w, r, p)

	case "POST":
		uname := r.FormValue("uname")
		pwd := r.FormValue("pwd")
		redirectTarget := "/"

		if uname == "admin" && pwd == "root" {
			redirectTarget = "http://www.google.com"
			http.Redirect(w, r, redirectTarget, http.StatusMovedPermanently)
		} else {
			http.Redirect(w, r, redirectTarget, http.StatusNotFound)
		}
	}
}

func main() {

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", GetIndexHandler).Methods("GET")
	r.HandleFunc("/portal.html", FormHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
