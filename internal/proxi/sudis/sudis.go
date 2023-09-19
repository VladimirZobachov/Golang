package sudis

import (
	"io/ioutil"
	"net/http"
)

type SudisAPI struct{}

func (s *SudisAPI) Auth(w http.ResponseWriter, r *http.Request, appKey string) (accessToken string, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://sudis.mvd.ru/api", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", appKey) // Set the appKey in the request header

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

	return string(body), err
}
