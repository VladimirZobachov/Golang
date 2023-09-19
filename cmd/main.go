package main

import (
	"api/internal/controller"
	"net/http"
)

func main() {

	controllerInstance := controller.Controller{}

	response := controllerInstance.Auth("appKey")

	test := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}

	http.HandleFunc("/test", test)

	http.ListenAndServe(":8080", nil)
}
