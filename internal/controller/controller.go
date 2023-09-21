package controller

import "net/http"

type endpoint string

const (
	ep1 endpoint = "get_data_1"
	// назови как удобно, добавь остальные эндпоинты
)

type SudisAPI interface {
	Auth(appKey string) (accessToken string, err error)
}

type AppAPI interface {
	GetData(accessToken string) (data []byte, err error)
}

type Controller struct {
	sudis        SudisAPI
	applications AppAPI
}

// todo: тут и в осталдьных местах лучше делать через конструктор
func NewController(s SudisAPI, a AppAPI) *Controller {
	return &Controller{
		sudis:        s,
		applications: a,
	}
}

func (c *Controller) GetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("incorrect method"))
		return
	}

	ep := r.Header.Get("endpoint")
	appKey := r.Header.Get("app_key")
	// todo: проверить, что параметры не пустые, если пустые вернуть ошибку, как выше сделано с проверкой метода

	accessToken, err := c.sudis.Auth(appKey)

	var data []byte
	// здесь мы приводим ep к нашему кастомному типу, чтобы сравнивать с заданными константами. в целом, можно
	// оставить и просто строку, но принято делать условный енум
	switch endpoint(ep) {
	case ep1:
		data, err = c.applications.GetData(accessToken)
	//case ep2:
	//	data, err = c.applications.GetEnd2(accessToken)
	//case ep3:
	//	data, err = c.applications.GetEnd1(accessToken)
	//case ep4:
	//	data, err = c.applications.GetEnd1(accessToken)
	default:
		// todo: вернуть ошибку, если метод не найден в определённых в константах
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
