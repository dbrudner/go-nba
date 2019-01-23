// endpoint to retrieve info on a single player
// src: http://data.nba.net/10s/prod/v1/2018/players/203500_profile.json
// date accessed: 2019-01-21 12:50:44.515

package nba

import "strings"

type PlayerProfile struct {
	Internal Internal `json:"_internal"`
	League   struct {
		Standard struct {
			TeamID string `json:"teamId"`
			Stats  struct {
				Latest        CareerSummary `json:"latest"`
				CareerSummary CareerSummary `json:"careerSummary"`
				RegularSeason struct {
					Season []Season `json:"season"`
				}
			}
		}
	}
}

type CareerSummary struct {
	Tpp           string  `json:"tpp"`
	FTP           string  `json:"ftp"`
	Fgp           string  `json:"fgp"`
	Ppg           string  `json:"ppg"`
	RPG           string  `json:"rpg"`
	Apg           string  `json:"apg"`
	Bpg           string  `json:"bpg"`
	Mpg           string  `json:"mpg"`
	Spg           string  `json:"spg"`
	Assists       string  `json:"assists"`
	Blocks        string  `json:"blocks"`
	Steals        string  `json:"steals"`
	Turnovers     string  `json:"turnovers"`
	OffReb        string  `json:"offReb"`
	DefReb        string  `json:"defReb"`
	TotReb        string  `json:"totReb"`
	Fgm           string  `json:"fgm"`
	Fga           string  `json:"fga"`
	TPM           string  `json:"tpm"`
	Tpa           string  `json:"tpa"`
	Ftm           string  `json:"ftm"`
	Fta           string  `json:"fta"`
	PFouls        string  `json:"pFouls"`
	Points        string  `json:"points"`
	GamesPlayed   string  `json:"gamesPlayed"`
	GamesStarted  string  `json:"gamesStarted"`
	PlusMinus     string  `json:"plusMinus"`
	Min           string  `json:"min"`
	Dd2           string  `json:"dd2"`
	Td3           string  `json:"td3"`
	SeasonYear    *int64  `json:"seasonYear,omitempty"`
	SeasonStageID *int64  `json:"seasonStageId,omitempty"`
	Topg          *string `json:"topg,omitempty"`
	TeamID        *string `json:"teamId,omitempty"`
}

type RegularSeason struct {
	Season []Season `json:"season"`
}

type Season struct {
	SeasonYear int64           `json:"seasonYear"`
	Teams      []CareerSummary `json:"teams"`
	Total      CareerSummary   `json:"total"`
}

func FetchPlayerProfile(endpoint string, id string) PlayerProfile {
	playerProfile := new(PlayerProfile)
	endpoint = strings.Replace(endpoint, "{{personId}}", id, 1)
	FetchNBAJSON(endpoint, playerProfile)

	return *playerProfile
}
