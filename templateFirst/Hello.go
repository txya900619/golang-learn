package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}
type indexViewModelIf struct {
	Title string
	User  User
}
type Post struct {
	User User
	Body string
}
type indexViewModelRange struct {
	Title string
	User  User
	Posts []Post
}

func main() {
	http.HandleFunc("/", sayhello)
	http.HandleFunc("/if/", testIf)
	http.HandleFunc("/range/", testRange)
	http.ListenAndServe(":8888", nil)
}
func sayhello(w http.ResponseWriter, r *http.Request) {
	user := User{Username: "wayne"}
	tpl, _ := template.ParseFiles("templates/index.html")
	tpl.Execute(w, &user)
}
func testIf(w http.ResponseWriter, r *http.Request) {
	user := User{Username: "wayne"}
	v := indexViewModelIf{Title: "ouo", User: user}
	tpl, _ := template.ParseFiles("templates/if.html")
	tpl.Execute(w, &v)
}
func testRange(w http.ResponseWriter, r *http.Request) {
	u1 := User{Username: "wayne"}
	u2 := User{Username: "alibaba"}
	posts := []Post{
		Post{User: u1, Body: "I am the god of new world"},
		Post{User: u2, Body: "The world!!!"},
	}
	v := indexViewModelRange{Title: "owo", User: u1, Posts: posts}
	tpl, _ := template.ParseFiles("templates/range.html")
	tpl.Execute(w, &v)

}
