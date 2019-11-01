package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	Username string
}
type Post struct {
	User
	Body string
}
type indexViewModel struct {
	User
	Title string
	Posts []Post
}

func main() {
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/range/", rangeFunc)
	http.ListenAndServe(":8888", nil)

}
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)
	base := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Fail to open dir:" + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Fail to read content from dir:" + err.Error())
	}
	for _, fi := range fis {

		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		defer f.Close()
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Fail to read content from file '" + fi.Name() + "'")
		}
		tmpl := template.Must(base.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Fail to parse content of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl

	}
	return result
}
func indexFunc(w http.ResponseWriter, r *http.Request) {
	u1 := User{Username: "wayne"}
	u2 := User{Username: "alibaba"}
	posts := []Post{
		Post{User: u1, Body: "za world"},
		Post{User: u2, Body: "olaolaolaola"},
	}
	v := indexViewModel{Title: "Homepage", Posts: posts, User: u1}
	templates := PopulateTemplates()
	templates["index.html"].Execute(w, &v)
}
func rangeFunc(w http.ResponseWriter, r *http.Request) {
	u1 := User{Username: "wayne"}
	u2 := User{Username: "alibaba"}
	posts := []Post{
		Post{User: u1, Body: "za world"},
		Post{User: u2, Body: "olaolaolaola"},
	}
	v := indexViewModel{Title: "Homepage", Posts: posts, User: u1}
	templates := PopulateTemplates()
	templates["range.html"].Execute(w, &v)
}
