package osu_lobby_inspector

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type User struct {
	Id               int
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

func (u *User) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("User id: %d\n", u.Id))
	builder.WriteString(fmt.Sprintf("Username: %s\n", u.Username))
	builder.WriteString(fmt.Sprintf("Join count: %d\n", u.JoinCount))
	builder.WriteString(fmt.Sprintf("Total time in lobby: %s\n", u.TotalTimeInLobby))
	builder.WriteString(fmt.Sprintf("Host count: %d\n", u.HostCount))
	builder.WriteString(fmt.Sprintf("Total hits: %d\n", u.TotalHits))
	builder.WriteString(fmt.Sprintf("Total misses: %d\n", u.TotalMisses))
	builder.WriteString(fmt.Sprintf("Max combo: %d\n", u.MaxCombo))
	builder.WriteString(fmt.Sprintf("Max score: %d\n", u.MaxScore))

	return builder.String()
}

func (l *LobbyStatistics) ProcessUsers() error {
	var users []User

	for _, user := range l.rawMatchData.Users {
		processedUser := User{
			Id:       user.Id,
			Username: user.Username,
		}

		var userLastJoinTime *time.Time = nil
		for _, event := range l.rawMatchData.Events {
			// Special case for lobby creator
			if event.Detail.Type == "match-created" && *event.UserId == user.Id {
				userLastJoinTime = &event.Timestamp
				processedUser.JoinCount++
			}

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
