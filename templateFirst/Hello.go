package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

func main() {
	http.HandleFunc("/", sayhello)
	http.ListenAndServe(":8888", nil)
}
func sayhello(w http.ResponseWriter, r *http.Request) {
	user := User{Username: "wayne"}
	tpl, _ := template.ParseFiles("templates/index.html")
	tpl.Execute(w, &user)
}
