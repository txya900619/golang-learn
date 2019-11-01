package main

import (
	"net/http"

	"github.com/txya900619/golangLearning/fileStruct/controller"
)

func main() {

	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
