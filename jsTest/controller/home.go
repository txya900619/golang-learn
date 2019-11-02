package controller

import (
	"net/http"

	"github.com/txya900619/golangLearning/jsTest/vm"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/range/", rangeHandler)
	http.HandleFunc("/home/", homeHandler)
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
