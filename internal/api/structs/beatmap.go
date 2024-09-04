package structs

type Beatmap struct {
	Id               int         `json:"id"`
	BeatmapSetId     int         `json:"beatmapset_id"`
	BeatmapSet       *BeatmapSet `json:"beatmapset"`
	DifficultyRating float64     `json:"difficulty_rating"`
	Mode             Ruleset     `json:"mode"`
	Status           string      `json:"status"`
	TotalLength      int         `json:"total_length"`
	UserId           int         `json:"user_id"`
	Version          string      `json:"version"`
}

type BeatmapSet struct {
	Id            int    `json:"id"`
	Artist        string `json:"artist"`
	ArtistUnicode string `json:"artist_unicode"`

	// API FIX: Covers
	Covers map[string]any `json:"covers"`

	Creator       string `json:"creator"`
	FavoriteCount int    `json:"favorite_count"`
	Nsfw          bool   `json:"nsfw"`
	Offset        int    `json:"offset"`
	PlayCount     int    `json:"play_count"`
	PreviewUrl    string `json:"preview_url"`
	Source        string `json:"source"`
	Status        string `json:"status"`
	Spotlight     bool   `json:"spotlight"`
	Title         string `json:"title"`
	TitleUnicode  string `json:"title_unicode"`
	UserId        int    `json:"user_id"`
	Video         bool   `json:"video"`
}
