package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age uint16
	Money int
	Hobbies [] string
}

type Contacts struct {
	Email string
	Telegram string
}


func home(w http.ResponseWriter, r *http.Request) {
	admin := User{"Boris Mukhin", 27, 15, []string{"Reading books", "Playing video games"}}
	templ, _ := template.ParseFiles("templates/home.html")
	templ.Execute(w, admin)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	contact_info := Contacts{"admin@gmail.com", "Admin"}
	templ, _ := template.ParseFiles("templates/contacts.html")
	templ.Execute(w, contact_info)
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contacts/", contacts)
	http.ListenAndServe(":8181", nil)
}

func main() {
	handleRequest()
}
