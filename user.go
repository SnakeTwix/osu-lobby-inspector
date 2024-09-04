package osu_lobby_inspector

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	UserId           int
	Username         string
	JoinCount        int
	TotalTimeInLobby time.Duration
	HostCount        int
	TotalHits        int
	TotalMisses      int
	MaxCombo         int
	MaxScore         int

	// TODO Picked maps
	// TODO Map Places
	// TODO Played Maps
	// TODO Mods played statistic
}

func (l *LobbyStatistics) ProcessUsers() error {
	var users []User

	for _, user := range l.rawMatchData.Users {
		processedUser := User{
			UserId:   user.Id,
			Username: user.Username,
		}

		var userLastJoinTime *time.Time = nil
		for _, event := range l.rawMatchData.Events {
			if event.UserId == nil || *event.UserId != user.Id {
				continue
			}

			if event.Detail.Type == "player-joined" {
				processedUser.JoinCount++
				userLastJoinTime = &event.Timestamp
			}

			if event.Detail.Type == "player-left" {
				if userLastJoinTime == nil {
					return errors.New(fmt.Sprintf("encountered player-left before player-join for player %s, match %d", user.Username, l.LobbyId))
				}

				lobbySessionDuration := event.Timestamp.Sub(*userLastJoinTime)
				processedUser.TotalTimeInLobby += lobbySessionDuration

				userLastJoinTime = nil
			}

		}

		users = append(users, processedUser)
	}

	l.Users = users
	return nil
}
