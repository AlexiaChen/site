package site

const (
	MessageAPI = "/api/stationMsg/addMsg"
)

type Site struct {
	APIUriPrefix  string
	APISignSecret string
}
