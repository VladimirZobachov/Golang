package applications

type AppAPI interface {
	GetEnd1(accessToken string) (data []byte, err error)
}
