package controller

type SudisAPI interface {
	Auth(appKey string) (accessToken string, err error)
}

type AppAPI interface {
	Get(accessToken string) (data []byte, err error)
}

type Controller struct {
	sudis        SudisAPI
	applications AppAPI
}

func (c *Controller) Auth(appKey string) (accessToken string) {
	return "[{'app_key':'sadfafdafa'}]"
}
