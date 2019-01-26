// endpoint to retrieve a team schedule
// src: https://stats.nba.com/stats/playercareerstats?PlayerID=201142&perMode=PerGame

package nba

type PlayerCareerStats struct {
	Resource   string      `json:"resource"`
	Parameters Parameters  `json:"parameters"`
	ResultSets []ResultSet `json:"resultSets"`
}

type Parameters struct {
	PerMode  string      `json:"PerMode"`
	PlayerID int64       `json:"PlayerID"`
	LeagueID interface{} `json:"LeagueID"`
}

type ResultSet struct {
	Name    string     `json:"name"`
	Headers []string   `json:"headers"`
	RowSet  [][]RowSet `json:"rowSet"`
}

type RowSet struct {
	Double *float64
	String *string
}

func FetchPlayerCareerStats(endpoint string, id string) PlayerCareerStats {
	playerCareerStats := new(PlayerCareerStats)
	endpoint = "https://stats.nba.com/stats/playercareerstats?PlayerID=" + id + "&perMode=PerGame"
	FetchNBADataJSON(endpoint, playerCareerStats)

	return *playerCareerStats
}
