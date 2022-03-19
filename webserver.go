/*
 * Written by Eric Dowds ed67@hw.ac.uk
 */

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

/*
 * This is our global variable to save our POST data from the
 * ESP8266
 */

type message struct {
	//	Shelf   string `json:"Shelf"`
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

// global variable needed to declare template variable
var tmpl *template.Template

/* This is needed instead of a static port because heroku
 * randomises ports when it starts
 */
func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
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
		err := tmpl.Execute(w, LoginContent)
		if err != nil {
			log.Fatalf("templating failed in FormHandler method: %v", err)
		}

	}
}

// simple static handlers go here

func FormPresenter(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, PortalContent)
	if err != nil {
		log.Fatalf("templating brokennnn: %v", err)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, AboutUsContent)
	if err != nil {
		log.Fatalf("templating brokennnn: %v", err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, IndexContent)
	if err != nil {
		log.Fatalf("templating brokennnn: %v", err)
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, ContactContent)
	if err != nil {
		log.Fatalf("templating brokennnn: %v", err)
	}
}

func OurTechHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, OurTechContent)
	if err != nil {
		log.Fatalf("templating brokennnn: %v", err)
	}
}

//this function handles data sent from the ESP8266

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

	log.Println("Here is the data sent from the ESP packaged into a struct: ")
	log.Println(ESPJson)
}

func SendToJS(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(&ESPJson)

	if err != nil {
		log.Println("fatal json error!")
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	//	w.Write(userJson) // TODO: this doesn't work investigate
}

func main() {

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = template.ParseFiles("./website/template.html")
	if err != nil {
		log.Fatalf("bizarre templating error: %v: ", err)
	}
	// static page handlers
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/login-success", FormHandler)
	http.HandleFunc("/portal", FormPresenter)
	http.HandleFunc("/about-us", AboutHandler)
	http.HandleFunc("/contact-us", ContactHandler)
	http.HandleFunc("/our-tech", OurTechHandler)

	// getting the data from the ESP onto the site
	http.HandleFunc("/ESP-requests", ESPHandler)
	http.HandleFunc("/SendToJS", SendToJS)

	// static fileservers for js and css and images
	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	js := http.FileServer(http.Dir("./js"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	images := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", images))

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
