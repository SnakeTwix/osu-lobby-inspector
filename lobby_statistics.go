package osu_lobby_inspector

import (
	"github.com/SnakeTwix/osu-lobby-inspector/internal/api"
	"time"
)

type StatisticsFetcher struct {
	clientId     int
	clientSecret string
	client       api.Client
}

type LobbyStatistics struct {
	rawMatchData  *api.MatchData
	LobbyId       int             `json:"lobby_id"`
	CreatedBy     int             `json:"created_by"`
	Users         map[int]User    `json:"users"`
	CreationDate  time.Time       `json:"creation_date"`
	DisbandedDate time.Time       `json:"disbanded_date"`
	Beatmaps      map[int]Beatmap `json:"beatmaps"`
	Mappers       map[string]int  `json:"mappers"`
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
		return LobbyStatistics{}, err
	}

	lobbyStats := LobbyStatistics{
		rawMatchData:  matchData,
		LobbyId:       lobbyId,
		CreationDate:  matchData.Match.StartTime,
		DisbandedDate: matchData.Match.EndTime,
		Mappers:       map[string]int{},
	}

	err = lobbyStats.processUsers()
	if err != nil {
		return LobbyStatistics{}, err
	}

	matchCreatorId := *matchData.Events[0].UserId
	for userIndex := range lobbyStats.Users {
		if lobbyStats.Users[userIndex].Id == matchCreatorId {
			lobbyStats.CreatedBy = lobbyStats.Users[userIndex].Id
		}
	}

	lobbyStats.processBeatmaps()

	return lobbyStats, nil
}
