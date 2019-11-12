package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../model"
	"../vm"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/range", rangeHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/todoList", todolistHandler)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["_base.html"].Execute(w, &v)
}
func rangeHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["range.html"].Execute(w, &v)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.LoginViewModelOp{}
	v := vop.GetVM()
	if r.Method == http.MethodGet {
		templates["login.html"].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if len(username) < 3 {
			v.AddError("username must longer than 3")
		}
		if len(password) < 6 {
			v.AddError("password must longer than 6")
		}
		if !check(username, password) {
			v.AddError("username password not correct,please input again")
		}
		if len(v.Errs) > 0 {
			templates["login.html"].Execute(w, &v)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

}
func todolistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newTodo model.Todolist
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(newTodo.Todothing, newTodo.Deadline)

	}
	if r.Method == http.MethodGet {
		vop := vm.LoginViewModelOp{}
		v := vop.GetVM()
		templates["todolist.html"].Execute(w, &v)
	}
}
func check(username, password string) bool {
	if username == "txya900619" && password == "25553706" {
		return true
	}
	return false
}
