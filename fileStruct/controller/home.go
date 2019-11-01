package controller

import (
	"net/http"

	"github.com/txya900619/golangLearning/fileStruct/vm"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}
