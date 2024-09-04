package osu_lobby_inspector

type Beatmap struct {
	Id          int     `json:"id"`
	Host        int     `json:"host"`
	TimesPlayed int     `json:"times_played"`
	StarRating  float64 `json:"star_rating"`
}

func (l *LobbyStatistics) processBeatmaps() {
	beatmaps := make(map[int]Beatmap)

	var currentHost int
	for _, event := range l.rawMatchData.Events {
		if event.Detail.Type == "host-change" {
			currentHost = *event.UserId
			continue
		}

		if event.Detail.Type != "other" {
			continue
		}

		// if Beatmap deleted
		if event.Game.Beatmap.BeatmapSet == nil {
			beatmap := beatmaps[0]
			beatmap.TimesPlayed++
			beatmaps[0] = beatmap

			continue
		}

		var beatmap, ok = beatmaps[event.Game.BeatmapId]
		if !ok {
			beatmap.Id = event.Game.BeatmapId
			beatmap.Host = currentHost
			beatmap.StarRating = event.Game.Beatmap.DifficultyRating
		}

		beatmap.TimesPlayed++
		beatmaps[beatmap.Id] = beatmap

		l.Mappers[event.Game.Beatmap.BeatmapSet.Creator]++
	}

	l.Beatmaps = beatmaps
}
