package structs

import "time"

type User struct {
	AvatarUrl     string     `json:"avatar_url"`
	CountryCode   string     `json:"country_code"`
	DefaultGroup  *string    `json:"default_group"`
	Id            int        `json:"id"`
	IsActive      bool       `json:"is_active"`
	IsBot         bool       `json:"is_bot"`
	IsDeleted     bool       `json:"is_deleted"`
	IsOnline      bool       `json:"is_online"`
	IsSupporter   bool       `json:"is_supporter"`
	LastVisit     *time.Time `json:"last_visit"`
	PmFriendsOnly bool       `json:"pm_friends_only"`
	ProfileColour *string    `json:"profile_colour"`
	Username      *string    `json:"username"`

	// API FIX: Add the additional fields as they are used in the endpoints covered by the wrapper
	Country *Country `json:"country"`
}
