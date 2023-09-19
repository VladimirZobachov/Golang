package controller

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

func (c *Controller) Auth(appKey string) (accessToken string, err error) {
	return c.sudis.Auth(appKey)
}

func (c *Controller) GetData(accessToken string) (data []byte, err error) {
	return c.applications.GetData(accessToken)
}
