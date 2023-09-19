package main

import (
	"api/internal/controller"
	"net/http"
)

func main() {

	controllerInstance := controller.Controller{}

	response := controllerInstance.GetData("appKey")

	getData := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}

	http.HandleFunc("/data", getData)

	http.ListenAndServe(":8080", nil)
}
