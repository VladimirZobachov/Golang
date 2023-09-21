package sudis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SudisAPI struct{}

func (s *SudisAPI) Auth(appKey string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://sudis.mvd.ru/api", nil)
	if err != nil {
		return "", fmt.Errorf("create request: %v", err)
	}

	req.Header.Set("Authorization", appKey) // Set the appKey in the request header

	resp, err := client.Do(req)
	if err != nil {
		// todo: обернуть и вернуть ошибку
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// todo: обернуть и вернуть ошибку
	}

	var auth Auth
	if err = json.Unmarshal(body, &auth); err != nil {
		// todo: обернуть и вернуть ошибку
	}

	// todo: вероятнее всего, здесь тебе придётся анмаршалить структуру токена в нужную структуру с полями и
	// возвращать конкретное поле, пример ниже
	return auth.Token, err
}

// это просто наугад как может выглядеть ответ сервиса авторизации
type Auth struct {
	Token     string    `json:"token"`
	ExpiredOn time.Time `json:"expired_on"`
	SomeField string    `json:"some_field"`
}
