// endpoint to retrieve a team schedule
// src: http://data.nba.net/10s/prod/v1/20171005/0011700027_boxscore.json
// date accessed: 2018-04-12 09:47:01.137

package nba

type PlayerGamelog struct {
	Internal Internal `json:"_internal"`
	League   League   `json:"league"`
}

type League struct {
	Standard []Gamelog_Standard `json:"standard"`
}

type Gamelog_Standard struct {
	GameID      string `json:"gameId"`
	GameDateUTC string `json:"gameDateUTC"`
	GameURLCode string `json:"gameUrlCode"`
	IsHomeGame  bool   `json:"isHomeGame"`
	HTeam       Team   `json:"hTeam"`
	VTeam       Team   `json:"vTeam"`
	Stats       struct {
		Points  string `json:"points"`
		OffReb  string `json:"offReb"`
		DefReb  string `json:"defReb"`
		TotReb  string `json:"totReb"`
		Assists string `json:"assists"`
	}
}

type Team struct {
	TeamID   string `json:"teamId"`
	Score    string `json:"score"`
	IsWinner bool   `json:"isWinner"`
}
