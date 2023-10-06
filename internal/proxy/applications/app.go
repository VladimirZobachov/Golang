package applications

type AppAPI struct{}

func NewAppAPI() *AppAPI {
	return new(AppAPI)
}

func (a *AppAPI) GetData(accessToken string) ([]byte, error) {
	//client := &http.Client{
	//	Transport: http.DefaultTransport,
	//	Timeout:   30 * time.Second, // не забывай устанавливать таймаут
	//}
	//
	//req, err := http.NewRequest("GET", "https://sad.oshs.mvd.ru/api", nil)
	//if err != nil {
	//	return nil, fmt.Errorf("create request: %w", err)
	//}
	//
	//req.Header.Set("Authorization", "Bearer "+accessToken) // Set the appKey in the request header
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	// todo: обернуть и вернуть ошибку
	//}
	//defer func() {
	//	_ = resp.Body.Close()
	//}()
	//
	//// не используй ioutil.ReadAll - он депрекейтед
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	// todo: обернуть и вернуть ошибку
	//}

	// если ошибка нил, то лучше явно это писать в ретёрне
	return []byte(accessToken), nil
}
