package osu_lobby_inspector

import (
	"github.com/SnakeTwix/osu-lobby-inspector/internal/api"
	"log"
	"time"
)

type StatisticFetcher struct {
	clientId     int
	clientSecret string
	client       api.Client
}

type LobbyStatistic struct {
	rawMatchData  *api.MatchData
	LobbyId       int       `json:"lobby_id"`
	CreatedBy     *User     `json:"created_by"`
	Users         []User    `json:"users"`
	CreationDate  time.Time `json:"creation_date"`
	DisbandedDate time.Time `json:"disbanded_date"`

	// TODO Picked maps
	// TODO Mapper pick count
}

func NewStatisticFetcher(clientId int, clientSecret string) (*StatisticFetcher, error) {
	client := api.New(clientId, clientSecret)
	err := client.GetToken()
	if err != nil {
		return nil, err
	}

	return &StatisticFetcher{
		clientId:     clientId,
		clientSecret: clientSecret,
		client:       client,
	}, nil

}

func (f *StatisticFetcher) GetLobbyStatistic(lobbyId int) (LobbyStatistic, error) {

	matchData, err := f.client.GetFullMatch(lobbyId)
	if err != nil {
		log.Fatal(err)
	}

	lobbyStats := LobbyStatistic{
		rawMatchData:  matchData,
		LobbyId:       lobbyId,
		CreationDate:  matchData.Match.StartTime,
		DisbandedDate: matchData.Match.EndTime,
	}

	lobbyStats.ProcessUsers()

	return lobbyStats, nil
}
