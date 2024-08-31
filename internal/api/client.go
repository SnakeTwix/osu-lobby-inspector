package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SnakeTwix/osu-lobby-inspector/util"
	"io"
	"net/http"
)

const BaseUrl = "https://osu.ppy.sh/api/v2"

type Client struct {
	url          string
	clientSecret string
	clientId     int
	authToken    *string
	httpClient   *http.Client
}

func New(clientId int, clientSecret string) Client {
	return Client{
		url:          BaseUrl,
		clientSecret: clientSecret,
		clientId:     clientId,
		httpClient:   http.DefaultClient,
	}
}

func (c *Client) GetToken() error {
	// TODO: get rid of body Map mappings
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

	response, err := c.httpClient.Post("https://osu.ppy.sh/oauth/token", "application/json", body)
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

// Spits out a request object prepended with the v2 api url and sets the Bearer token
func (c *Client) getRequestV2(method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.url, url), body)
	if err != nil {
		return nil, err
	}

	if c.authToken == nil {
		return nil, errors.New("no auth token specified")
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *c.authToken))

	return request, nil

}
