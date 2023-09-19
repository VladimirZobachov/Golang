package sudis

type SudisAPI interface {
	Auth(appKey string) (accessToken string, err error)
}
