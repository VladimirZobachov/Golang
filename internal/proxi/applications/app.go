package applications

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type AppAPI struct{}

// тоже лучше конструктор, несмотря на то, что у структуры нет полей
func NewAppAPI() *AppAPI {
	return new(AppAPI)
}

// todo: тебе здесь совершенно не нужны ни реквест ни респонзврайтер, ты просто возвращаешь []byte и ошибку

func (a *AppAPI) GetData(accessToken string) ([]byte, error) {
	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   30 * time.Second, // не забывай устанавливать таймаут
	}

	req, err := http.NewRequest("GET", "https://sad.oshs.mvd.ru/api", nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken) // Set the appKey in the request header

	resp, err := client.Do(req)
	if err != nil {
		// todo: обернуть и вернуть ошибку
	}
	defer resp.Body.Close()

	// не используй ioutil.ReadAll - он депрекейтед
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// todo: обернуть и вернуть ошибку
	}

	// если ошибка нил, то лучше явно это писать в ретёрне
	return body, nil
}
