package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"api/internal/proxy/sudis/thrift/gen-go/ccispauth"
)

func main() {
	cfg := &thrift.TConfiguration{
		ConnectTimeout: 15 * time.Second,
		SocketTimeout:  15 * time.Second,
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(cfg)
	transportFactory := thrift.NewTTransportFactory()

	transport, err := thrift.NewTServerSocket("localhost:9090")
	if err != nil {
		log.Fatalln(err)
	}

	processor := ccispauth.NewTCciSpAuthProcessor(NewHandler())
	srv := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	if err = srv.Serve(); err != nil {
		fmt.Println(err)
	}
}

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateToken(ctx context.Context, arguments *ccispauth.TCciSpAuthCreateTokenArgs_) (_r *ccispauth.TCciSpAuthTokenResult_, _err error) {
	spCode := "your_code_1"
	targetSpCode := "your_code_target_1"
	return &ccispauth.TCciSpAuthTokenResult_{
		Version:        0,
		ResponseMillis: 0,
		ResponseNonce:  nil,
		Result_:        "123",
		ResultMessage:  "321",
		TokenData: &ccispauth.TCciSpAuthTokenData{
			Version:      ccispauth.TCciSpAuthTokenData_Version_DEFAULT,
			TokenId:      []byte("1234"),
			ExpireMillis: 1000,
			SpCode:       &spCode,
			TargetSpCode: &targetSpCode,
		},
	}, nil
}

func (h *Handler) TokenData(ctx context.Context, arguments *ccispauth.TCciSpAuthTokenDataArgs_) (_r *ccispauth.TCciSpAuthTokenResult_, _err error) {
	spCode := "your_code_2"
	targetSpCode := "your_code_target_2"
	return &ccispauth.TCciSpAuthTokenResult_{
		Version:        ccispauth.TCciSpAuthVersion_V1,
		ResponseMillis: 1000,
		ResponseNonce:  nil,
		Result_:        "123",
		ResultMessage:  "321",
		TokenData: &ccispauth.TCciSpAuthTokenData{
			Version:      ccispauth.TCciSpAuthTokenData_Version_DEFAULT,
			TokenId:      []byte("1234"),
			ExpireMillis: 1000,
			SpCode:       &spCode,
			TargetSpCode: &targetSpCode,
		},
	}, nil
}
