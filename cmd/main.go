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
	cfg := &thrift.TConfiguration{
		ConnectTimeout: 15 * time.Second,
		SocketTimeout:  15 * time.Second,
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(cfg)
	transportFactory := thrift.NewTTransportFactory()

	// уточнить какой используется транспорт, если не стандартный
	var transport thrift.TTransport
	transport = thrift.NewTSocketConf("localhost:9090", cfg)

	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = transport.Close() }()

	if err = transport.Open(); err != nil {
		log.Fatalln(err)
	}

	// todo: уточнить какой протокол и клиент использовать, их здесь несколько штук. может быть конфигурация ещё понадобится
	in := protocolFactory.GetProtocol(transport)
	out := protocolFactory.GetProtocol(transport)

	sds := sudis.NewSudisAPI(ccispauth.NewTCciSpAuthClient(thrift.NewTStandardClient(in, out)))
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
