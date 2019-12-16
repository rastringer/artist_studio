package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
}

func displayHome(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "home", nil)
}

func displayAbout(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "about", nil)
}

func displayWorks(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "works", nil)
}

func displayContact(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "contact", nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", displayHome)
	http.HandleFunc("/about", displayAbout)
	http.HandleFunc("/works", displayWorks)
	http.HandleFunc("/contact", displayContact)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	server.ListenAndServe()
}
