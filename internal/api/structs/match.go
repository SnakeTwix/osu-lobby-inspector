package structs

import "time"

type MatchEvent struct {
	Id        int              `json:"id"`
	Detail    MatchEventDetail `json:"detail"`
	Timestamp time.Time        `json:"timestamp"`
	UserId    *int             `json:"user_id"` // Could be null
	Game      *MatchGame       `json:"game"`    // Optional
}

type MatchEventDetail struct {
	Type string  `json:"type"`
	Text *string `json:"text"`
}

type Match struct {
	Id        int       `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Name      string    `json:"name"`
}
