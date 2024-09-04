package osu_lobby_inspector

import (
	"errors"
	"fmt"
	"github.com/SnakeTwix/osu-lobby-inspector/internal/api/structs"
	"strings"
	"time"
)

type User struct {
	Id               int           `json:"id"`
	Username         string        `json:"username"`
	JoinCount        int           `json:"join_count"`
	TotalTimeInLobby time.Duration `json:"total_time_in_lobby"`
	HostCount        int           `json:"host_count"`
	TotalHits        int           `json:"total_hits"`
	TotalMisses      int           `json:"total_misses"`
	MaxCombo         int           `json:"max_combo"`
	MaxScore         int           `json:"max_score"`

	lastJoinTime *time.Time

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

		for _, event := range l.rawMatchData.Events {
			switch event.Detail.Type {
			case "other":
				processedUser.ProcessGame(&event)
			case "match-created":
				// Special case for lobby creator
				if *event.UserId == user.Id {
					processedUser.ProcessJoin(&event)
				}
			case "player-joined":
				processedUser.ProcessJoin(&event)
			case "player-left":
				err := processedUser.ProcessLeft(&event)
				if err != nil {
					return fmt.Errorf("match: %d, %w", l.LobbyId, err)
				}
			case "host-changed":
				processedUser.ProcessHost(&event)
			}

		}

		users = append(users, processedUser)
	}

	l.Users = users
	return nil
}

func (u *User) ProcessGame(event *structs.MatchEvent) {
	var userScore structs.Score

	for _, score := range event.Game.Scores {
		if score.UserID == u.Id {
			userScore = score
			break
		}
	}

	u.TotalHits += userScore.Statistics.Count50
	u.TotalHits += userScore.Statistics.Count100
	u.TotalHits += userScore.Statistics.Count300
	u.TotalMisses += userScore.Statistics.CountMiss
	u.MaxCombo = max(u.MaxCombo, userScore.MaxCombo)
	u.MaxScore = max(u.MaxScore, userScore.Score)
}

func (u *User) ProcessJoin(event *structs.MatchEvent) {
	if *event.UserId != u.Id {
		return
	}

	u.JoinCount++
	u.lastJoinTime = &event.Timestamp
}

func (u *User) ProcessLeft(event *structs.MatchEvent) error {
	if *event.UserId != u.Id {
		return nil
	}

	if u.lastJoinTime == nil {
		return errors.New(fmt.Sprintf("encountered player-left before player-join for player %s", u.Username))
	}

	lobbySessionDuration := event.Timestamp.Sub(*u.lastJoinTime)
	u.TotalTimeInLobby += lobbySessionDuration

	u.lastJoinTime = nil
	return nil
}

func (u *User) ProcessHost(event *structs.MatchEvent) {
	if *event.UserId != u.Id {
		return
	}

	u.HostCount++
}
