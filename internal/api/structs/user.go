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
	Username      string     `json:"username"`

	// API FIX: Add the additional fields as they are used in the endpoints covered by the wrapper
	Country *Country `json:"country"`
	// cover
	// is_restricted (Authenticated only)
	// kudosu
}

type UserExtended struct {
	User
	HasSupported bool      `json:"has_supported"`
	JoinDate     time.Time `json:"join_date"`
	Discord      *string   `json:"discord"`
	Twitter      *string   `json:"twitter"`
	Website      *string   `json:"website"`
	Interests    *string   `json:"interests"`
	Location     *string   `json:"location"`
	Occupation   *string   `json:"occupation"`
	MaxBlocks    int       `json:"max_blocks"`
	MaxFriends   int       `json:"max_friends"`
	Ruleset      Ruleset   `json:"ruleset"`
	Playstyle    []string  `json:"playstyle"`
	PostCount    int       `json:"post_count"`
	ProfileHue   *int      `json:"profile_hue"`
	Title        *string   `json:"title"`
	TitleUrl     *string   `json:"title_url"`

	// API FIX: profile_order
}
