package structs

import "time"

type Score struct {
	ID         *int            `json:"id"`
	Accuracy   float64         `json:"accuracy"`
	BestID     *int            `json:"best_id"`
	CreatedAt  time.Time       `json:"created_at"`
	Match      *ScoreMatchInfo `json:"match"`
	MaxCombo   int             `json:"max_combo"`
	Mode       string          `json:"mode"`
	ModeInt    int             `json:"mode_int"`
	Mods       []string        `json:"mods"`
	Passed     bool            `json:"passed"`
	Perfect    int             `json:"perfect"`
	Pp         *float64        `json:"pp"` // No idea if this is actually float
	Rank       string          `json:"rank"`
	Replay     bool            `json:"replay"`
	Score      int             `json:"score"`
	Statistics struct {
		Count100  int `json:"count_100"`
		Count300  int `json:"count_300"`
		Count50   int `json:"count_50"`
		CountGeki int `json:"count_geki"`
		CountKatu int `json:"count_katu"`
		CountMiss int `json:"count_miss"`
	} `json:"statistics"`
	Type   string `json:"type"`
	UserID int    `json:"user_id"`
}

type ScoreMatchInfo struct {
	Pass bool   `json:"pass"`
	Slot int    `json:"slot"`
	Team string `json:"team"`
}
