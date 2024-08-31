package client

import (
	"encoding/json"
	"errors"
	"github.com/SnakeTwix/osu-lobby-inspector/util"
	"net/http"
)

const BaseUrl = "https://osu.ppy.sh/api/v2"

type Client struct {
	url          string
	clientSecret string
	clientId     int
	authToken    *string
}

func New(clientId int, clientSecret string) Client {
	return Client{
		url:          BaseUrl,
		clientSecret: clientSecret,
		clientId:     clientId,
	}
}

func (c *Client) GetToken() error {
	content := map[string]any{
		"client_id":     c.clientId,
		"client_secret": c.clientSecret,
		"grant_type":    "client_credentials",
		"scope":         "public",
	}

	body, err := util.MapToReader(content)
	if err != nil {
		return err
	}

	response, err := http.Post("https://osu.ppy.sh/oauth/token", "application/json", body)
	if err != nil {
		return err
	}

	var tokenMap map[string]any
	err = json.NewDecoder(response.Body).Decode(&tokenMap)
	if err != nil {
		return err
	}

	accessToken, ok := tokenMap["access_token"]
	if !ok {
		return err
	}

	stringToken, ok := accessToken.(string)
	if !ok {
		return errors.New("accessToken isn't string")
	}

	c.authToken = &stringToken
	return nil
}
