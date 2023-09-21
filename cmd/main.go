package main

import (
	"errors"
	"log"
	"time"

	"api/internal/controller"
	"api/internal/proxi/applications"
	"api/internal/proxi/sudis"

	"net/http"
)

func main() {
	sds := &sudis.SudisAPI{} // тоже переделать на конструктор
	app := applications.NewAppAPI()

	ctl := controller.NewController(sds, app)

	r := http.NewServeMux()
	r.HandleFunc("/data", ctl.GetData) // конкретный эндпоинт

	server := &http.Server{
		Addr:         "0.0.0.0:8080", // адрес и порт, где твой сервер ожидает запрос
		Handler:      r,
		ReadTimeout:  10 * time.Second, // todo: выставить таймауты
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		// если при закрытии сервера вернётся ошибка отличная от корректного завершения сервера, то её надо обработать
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}
}
