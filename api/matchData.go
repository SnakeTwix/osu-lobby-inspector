package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SnakeTwix/osu-lobby-inspector/api/structs"
	"log"
	"net/http"
	"strconv"
)

type MatchData struct {
	Match  structs.Match        `json:"match"`
	Events []structs.MatchEvent `json:"events"`

	// API FIX: User interface
	Users []map[string]any `json:"users"`

	FirstEventId  int `json:"first_event_id"`
	LatestEventId int `json:"latest_event_id"`
}

type GetMatchQuery struct {
	MatchId int
	Before  int
	After   int
	Limit   int
}

func (c *Client) GetMatch(query GetMatchQuery) (*MatchData, error) {
	if query.MatchId == 0 {
		return nil, errors.New("no id provided for match query")
	}

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/matches/%d", c.url, query.MatchId), nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	if query.Limit != 0 {
		q.Set("limit", strconv.Itoa(query.Limit))
	}

	if query.After != 0 {
		q.Set("after", strconv.Itoa(query.After))
	}

	if query.Before != 0 {
		q.Set("before", strconv.Itoa(query.Before))
	}

	request.URL.RawQuery = q.Encode()

	if c.authToken == nil {
		return nil, errors.New("no auth token specified")
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *c.authToken))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	log.Printf("Got match data")

	var matchData MatchData
	decoder := json.NewDecoder(response.Body)
	decoder.UseNumber()
	err = decoder.Decode(&matchData)
	if err != nil {
		return nil, err
	}

	return &matchData, nil
}

func (c *Client) GetFullMatch(id int) (*MatchData, error) {
	matchQuery := GetMatchQuery{MatchId: id}

	matchData, err := c.GetMatch(matchQuery)
	if err != nil {
		return nil, err
	}

	events := matchData.Events

	// Fetch until we get all the events for a match
	for events[0].Id != matchData.FirstEventId {
		matchQuery.Before = events[0].Id
		matchData, err = c.GetMatch(matchQuery)
		if err != nil {
			return nil, err
		}

		events = append(matchData.Events, events...)
	}

	matchData.Events = events

	return matchData, nil
}
