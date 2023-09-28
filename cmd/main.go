package main

import (
	"errors"
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"api/internal/controller"
	"api/internal/proxy/applications"
	"api/internal/proxy/sudis"
	"api/internal/proxy/sudis/thrift/gen-go/ccispauth"

	"net/http"
)

func main() {
	// уточнить какой используется транспорт, если не стандартный
	transport, err := thrift.NewTHttpClient("https://sudis.mvd.ru/api")
	if err != nil {
		log.Fatalln("new thrift client:", err)
	}
	defer func() {
		_ = transport.Close()
	}()

	// todo: уточнить какой протокол и клиент использовать, их здесь несколько штук. может быть конфигурация ещё понадобится
	// я взял для примера инициализации
	client := thrift.NewTStandardClient(thrift.NewTJSONProtocol(transport), thrift.NewTJSONProtocol(transport))

	sds := sudis.NewSudisAPI(ccispauth.NewTCciSpAuthClient(client))
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
