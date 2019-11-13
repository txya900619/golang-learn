package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"../model"
	"../vm"
	_ "github.com/go-sql-driver/mysql"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/range", rangeHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/todoList", todolistHandler)
	http.HandleFunc("/todoList/lists", todolistGetlistsHandler)
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
		db, err := sql.Open("mysql", "root:Fuck06050@/todolist?charset=utf8")
		defer db.Close()
		checkErr(err)
		var newTodo model.Todolist
		err = json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(newTodo.Todothing, newTodo.Deadline)

		stmt, err := db.Prepare("INSERT todo SET todo_describe=?, todo_deadline=?, todo_status=0")
		checkErr(err)
		_, err = stmt.Exec(newTodo.Todothing, newTodo.Deadline)

	}
	if r.Method == http.MethodGet {

		vop := vm.LoginViewModelOp{}
		v := vop.GetVM()
		templates["todolist.html"].Execute(w, &v)
	}
}
func todolistGetlistsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sql.Open("mysql", "root:Fuck06050@/todolist?charset=utf8&&parseTime=true")
	defer db.Close()
	checkErr(err)
	var getTodo []*model.TodolistGET
	rows, err := db.Query("SELECT * FROM todo")
	checkErr(err)
	for rows.Next() {
		var onetodo model.TodolistGET
		err = rows.Scan(&onetodo.Todothing, &onetodo.Deadline, &onetodo.Status, &onetodo.Id)
		checkErr(err)
		fmt.Println(onetodo)
		getTodo = append(getTodo, &onetodo)
	}
	err = json.NewEncoder(w).Encode(getTodo)
	checkErr(err)
}
func check(username, password string) bool {
	if username == "txya900619" && password == "25553706" {
		return true
	}
	return false
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
