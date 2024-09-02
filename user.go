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

func (f *StatisticFetcher) GetUser(userId int) (User, error) {
	//fetchedUser, err := f.client

	user := User{UserId: 69}

	return user, nil
}

func (f *StatisticFetcher) GetUsers(userIds []int) ([]User, error) {
	var users []User

	for _, userId := range userIds {
		user, err := f.GetUser(userId)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
