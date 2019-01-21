package main

import (
	"errors"
	"strings"
)

// endpoint to retrieve teams
// src: http://data.nba.net/prod/v2/2018/teams.json
// date accessed: 2018-08-30 20:00:04.422

type Teams struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		XSLT        string `json:"xslt"`
		EventName   string `json:"eventName"`
	}

	League struct {
		Standard   []Teams_Team `json:"standard"`
		Africa     []Teams_Team `json:"africa"`
		Sacramento []Teams_Team `json:"sacramento"`
		Vegas      []Teams_Team `json:"vegas"`
		Utah       []Teams_Team `json:"utah"`
	}
}

type Teams_Team struct {
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

func getAllTeams(endpoint string) Teams {
	teams := new(Teams)
	getNBAJSON(endpoint, teams)
	return *teams
}

func getNBATeams(teams []Teams_Team) []Teams_Team {
	NBATeams := []Teams_Team{}

	for _, team := range teams {
		if team.IsNBAFranchise {
			NBATeams = append(NBATeams, team)
		}
	}

	return NBATeams
}

// getTeamID searches through teams and returns an ID for a single team as a string
// tricode is a string with length of 3 (NYK for knicks)
// teams is an array of team
// returns a single team ID as a string
func getTeamID(tricode string, teams []Teams_Team) (string, error) {
	for _, team := range teams {
		if strings.ToLower(tricode) == strings.ToLower(team.Tricode) {
			return team.TeamID, nil
		}
	}
	return "", errors.New("No team found")
}
