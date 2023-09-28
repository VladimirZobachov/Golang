package sudis

import (
	"context"
	"log"
	"time"

	"api/internal/proxy/sudis/thrift/gen-go/ccispauth"
)

type SudisAPI struct {
	client *ccispauth.TCciSpAuthClient
}

func NewSudisAPI(client *ccispauth.TCciSpAuthClient) *SudisAPI {
	return &SudisAPI{client: client}
}

func (s *SudisAPI) Auth(ctx context.Context, spCode, targetSpCode string) (string, error) {
	resp, err := s.client.CreateToken(ctx, &ccispauth.TCciSpAuthCreateTokenArgs_{
		Version:       ccispauth.TCciSpAuthVersion_V1,
		RequestMillis: time.Now().UnixMilli(),
		RequestNonce:  nil, // не знаю что сюда записать, видимо это один из кодов приложения, либо оставить nil для первого запроса
		SpCode:        &spCode,
		TargetSpCode:  &targetSpCode,
	})
	if err != nil {
		log.Println(err)
		return "", err
	}

	// тут не уверен, что надо взять именно это поле в качестве токена. надо уточнить механику дальнейшего взаимодействия
	return string(resp.TokenData.TokenId), nil
}
