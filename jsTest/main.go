package main

import (
	"net/http"

	"./controller"
)

func main() {

	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
