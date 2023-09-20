package applications

import (
	"io/ioutil"
	"net/http"
)

type AppAPI struct{}

func (a *AppAPI) GetData(w http.ResponseWriter, r *http.Request, accessToken string) (data []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://sad.oshs.mvd.ru/api", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken) // Set the appKey in the request header

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return []byte(body), err
}
