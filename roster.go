package main

// Structs for endpoint to retrieve team roster
// Only returns a list of players with playerIds (personId)
// src: http://data.nba.net/prod/v1/2018/teams/1610612738/roster.json
// date accessed: 2019-01-21 06:48:32.916

type Roster struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		XSLT        string `json:"xslt"`
		EventName   string `json:"eventName"`
	}
	Leagues struct {
		Standard   Roster_League `json:"standard"`
		Africa     Roster_League `json:"africa"`
		Sacramento Roster_League `json:"sacramento"`
		Vegas      Roster_League `json:"vegas"`
		Utah       Roster_League `json:"utah"`
	}
}

type Roster_League struct {
	TeamID  string `json:"teamId"`
	Players []struct {
		PersonID string `json:"personId"`
	} `json:"players"`
}
