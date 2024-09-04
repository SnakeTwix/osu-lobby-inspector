package api

import "github.com/SnakeTwix/osu-lobby-inspector/internal/api/structs"

// GetUserQuery if UserId is provided, then it is used, if not then username. (Currently only userId)
type GetUserQuery struct {
	UserId   int
	Username string
	Ruleset  structs.Ruleset
}

//func (c *Client) GetUser(query GetUserQuery) (*structs.UserExtended, error) {
//	if query.UserId == 0 {
//		return nil, errors.New("no user id provided")
//	}
//
//	var requestUrlBuilder strings.Builder
//	requestUrlBuilder.WriteString(fmt.Sprintf("/users/%d", query.UserId))
//
//	if query.Ruleset != "" {
//
//	}
//
//}
