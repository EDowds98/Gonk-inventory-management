/*
 * Written by Eric Dowds ed67@hw.ac.uk
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

/* This is our global variable to save our POST data from the
 * ESP8266
 */

type message struct {
	Shelf   int    `json:"Shelf"`
	Module1 []bool `json:"1"`
	Module2 []bool `json:"2"`
	Module3 []bool `json:"3"`
	Module4 []bool `json:"4"`
	Module5 []bool `json:"5"`
	Module6 []bool `json:"6"`
	Module7 []bool `json:"7"`
	Module8 []bool `json:"8"`
}

/*
 * This is an instance of the struct and we'll fill this with JSON
 * and send it to the js program on the front end
 */

var ESPJson message

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

// simple static handlers go here

func FormPresenter(w http.ResponseWriter, r *http.Request) {
	p := ("./website/portal.html")
	fmt.Println("at formPresenter")
	http.ServeFile(w, r, p)

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	p := ("./website/about-us.html")
	http.ServeFile(w, r, p)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := ("./website/index.html")
	http.ServeFile(w, r, p)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	p := ("./website/contact-us.html")
	http.ServeFile(w, r, p)
}

//this function I hope will handle data sent from the ESP8266

func ESPHandler(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&ESPJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(ESPJson)
}

func SendToJS(w http.ResponseWriter, r *http.Request) {
	userJson, err := json.Marshal(&ESPJson)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}

func main() {

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/login-success", FormHandler)
	http.HandleFunc("/portal", FormPresenter)
	http.HandleFunc("/about-us", AboutHandler)
	http.HandleFunc("/contact-us", ContactHandler)

	http.HandleFunc("/ESP-requests", ESPHandler)

	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	js := http.FileServer(http.Dir("./js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
