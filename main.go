package main

import (
	"fmt"
	"net/http"

	"github.com/leon0707/go_webservice/controller"
)

func main() {
	uc := controller.RegisterController()
	fmt.Println("listen on 3001")
	error := http.ListenAndServe(":3001", uc)
	fmt.Println(error)
}
