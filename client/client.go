package client

const BaseUrl = "https://osu.ppy.sh/api/v2"

type Client struct {
	url          string
	clientSecret string
	clientId     string
}

func New(clientSecret string, clientId string) Client {
	return Client{
		url:          BaseUrl,
		clientSecret: clientSecret,
		clientId:     clientId,
	}
}
