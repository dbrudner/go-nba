package main

// Structs for endpoint to retrieve teams
// src: http://data.nba.net/prod/v2/2018/teams.json
// date accessed: 2018-08-30 20:00:04.422

type Teams struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		XSLT        string `json:"xslt"`
		EventName   string `json:"eventName"`
	}

	League struct {
		Standard   []Africa `json:"standard"`
		Africa     []Africa `json:"africa"`
		Sacramento []Africa `json:"sacramento"`
		Vegas      []Africa `json:"vegas"`
		Utah       []Africa `json:"utah"`
	}
}

type Africa struct {
	IsNBAFranchise bool         `json:"isNBAFranchise"`
	IsAllStar      bool         `json:"isAllStar"`
	City           string       `json:"city"`
	AltCityName    string       `json:"altCityName"`
	FullName       string       `json:"fullName"`
	Tricode        string       `json:"tricode"`
	TeamID         string       `json:"teamId"`
	Nickname       string       `json:"nickname"`
	URLName        string       `json:"urlName"`
	ConfName       ConfNameEnum `json:"confName"`
	DivName        DivNameEnum  `json:"divName"`
}

type ConfNameEnum string

const (
	ConfName   ConfNameEnum = ""
	East       ConfNameEnum = "East"
	Intl       ConfNameEnum = "Intl"
	Sacramento ConfNameEnum = "Sacramento"
	Summer     ConfNameEnum = "summer"
	Utah       ConfNameEnum = "Utah"
	West       ConfNameEnum = "West"
)

type DivNameEnum string

const (
	Atlantic  DivNameEnum = "Atlantic"
	Central   DivNameEnum = "Central"
	DivName   DivNameEnum = ""
	Northwest DivNameEnum = "Northwest"
	Pacific   DivNameEnum = "Pacific"
	Southeast DivNameEnum = "Southeast"
	Southwest DivNameEnum = "Southwest"
)

func getTeams(endpoint string) Teams {
	teams := new(Teams)
	getNBAJSON(endpoint, teams)
	return *teams
}
