package osu_lobby_inspector

import (
	"github.com/SnakeTwix/osu-lobby-inspector/internal/api"
	"log"
	"time"
)

type StatisticsFetcher struct {
	clientId     int
	clientSecret string
	client       api.Client
}

type LobbyStatistics struct {
	rawMatchData  *api.MatchData
	LobbyId       int       `json:"lobby_id"`
	CreatedBy     *User     `json:"created_by"`
	Users         []User    `json:"users"`
	CreationDate  time.Time `json:"creation_date"`
	DisbandedDate time.Time `json:"disbanded_date"`

	// TODO Picked maps
	// TODO Mapper pick count
}

func NewStatisticsFetcher(clientId int, clientSecret string) (*StatisticsFetcher, error) {
	client := api.New(clientId, clientSecret)
	err := client.GetToken()
	if err != nil {
		return nil, err
	}

	return &StatisticsFetcher{
		clientId:     clientId,
		clientSecret: clientSecret,
		client:       client,
	}, nil

}

func (f *StatisticsFetcher) FetchLobbyStatistics(lobbyId int) (LobbyStatistics, error) {
	matchData, err := f.client.GetFullMatch(lobbyId)
	if err != nil {
		log.Fatal(err)
	}

	lobbyStats := LobbyStatistics{
		rawMatchData:  matchData,
		LobbyId:       lobbyId,
		CreationDate:  matchData.Match.StartTime,
		DisbandedDate: matchData.Match.EndTime,
	}

	err = lobbyStats.ProcessUsers()
	if err != nil {
		return LobbyStatistics{}, err
	}

	return lobbyStats, nil
}
