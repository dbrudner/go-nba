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
	League struct {
		Standard   Africa `json:"standard"`
		Africa     Africa `json:"africa"`
		Sacramento Africa `json:"sacramento"`
		Vegas      Africa `json:"vegas"`
		Utah       Africa `json:"utah"`
	}
}

type Africa struct {
	TeamID  string   `json:"teamId"`
	Players []Player `json:"players"`
}

type Player struct {
	PersonID string `json:"personId"`
}
