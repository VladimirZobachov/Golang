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

func (c *Controller) GetData(appKey string) (data []byte) {
	accessToken, err := c.sudis.Auth("appKey")
	if err != nil {
		return
	}

	result, err := c.applications.GetData(accessToken)
	if err != nil {
		return
	}

	return []byte(result)
}
