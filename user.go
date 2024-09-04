package osu_lobby_inspector

import "time"

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

func (l *LobbyStatistic) ProcessUsers() {
	var users []User

	for _, user := range l.rawMatchData.Users {
		processedUser := User{
			UserId:   user.Id,
			Username: user.Username,
		}

		users = append(users, processedUser)
	}

	l.Users = users
}
