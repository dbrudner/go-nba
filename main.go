package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type NBAInfoJSON struct {
	_internal struct {
		EventName   string `json:"eventName"`
		PubDateTime string `json:"pubDateTime"`
		Xslt        string `json:"xslt"`
	} `json:"_internal"`
	Links struct {
		AllstarRoster               string `json:"allstarRoster"`
		AnchorDate                  string `json:"anchorDate"`
		Boxscore                    string `json:"boxscore"`
		Calendar                    string `json:"calendar"`
		CurrentDate                 string `json:"currentDate"`
		CurrentScoreboard           string `json:"currentScoreboard"`
		GameBookPdf                 string `json:"gameBookPdf"`
		LeadTracker                 string `json:"leadTracker"`
		LeagueConfStandings         string `json:"leagueConfStandings"`
		LeagueDivStandings          string `json:"leagueDivStandings"`
		LeagueLastFiveGameTeamStats string `json:"leagueLastFiveGameTeamStats"`
		LeagueMiniStandings         string `json:"leagueMiniStandings"`
		LeagueRosterCoaches         string `json:"leagueRosterCoaches"`
		LeagueRosterPlayers         string `json:"leagueRosterPlayers"`
		LeagueSchedule              string `json:"leagueSchedule"`
		LeagueTeamStatsLeaders      string `json:"leagueTeamStatsLeaders"`
		LeagueUngroupedStandings    string `json:"leagueUngroupedStandings"`
		MiniBoxscore                string `json:"miniBoxscore"`
		Pbp                         string `json:"pbp"`
		PlayerGameLog               string `json:"playerGameLog"`
		PlayerProfile               string `json:"playerProfile"`
		PlayerUberStats             string `json:"playerUberStats"`
		PlayoffSeriesLeaders        string `json:"playoffSeriesLeaders"`
		PlayoffsBracket             string `json:"playoffsBracket"`
		PreviewArticle              string `json:"previewArticle"`
		RecapArticle                string `json:"recapArticle"`
		Scoreboard                  string `json:"scoreboard"`
		TeamLeaders                 string `json:"teamLeaders"`
		TeamLeaders2                string `json:"teamLeaders2"`
		TeamRoster                  string `json:"teamRoster"`
		TeamSchedule                string `json:"teamSchedule"`
		TeamScheduleYear            string `json:"teamScheduleYear"`
		TeamScheduleYear2           string `json:"teamScheduleYear2"`
		Teams                       string `json:"teams"`
		TeamsConfig                 string `json:"teamsConfig"`
		TeamsConfigYear             string `json:"teamsConfigYear"`
		TodayScoreboard             string `json:"todayScoreboard"`
	} `json:"links"`
	SeasonScheduleYear int  `json:"seasonScheduleYear"`
	ShowPlayoffsClinch bool `json:"showPlayoffsClinch"`
}

type LeagueRosterPlayers struct {
	Internal Internal `json:"_internal"`
	League   League   `json:"league"`   
}

type Internal struct {
	PubDateTime string `json:"pubDateTime"`
	XSLT        string `json:"xslt"`       
	EventName   string `json:"eventName"`  
}

type League struct {
	Standard   []Player `json:"standard"`  
	Africa     []Player `json:"africa"`    
	Sacramento []Player `json:"sacramento"`
	Vegas      []Player `json:"vegas"`     
	Utah       []Player `json:"utah"`      
}

type Player struct {
	FirstName       string           `json:"firstName"`      
	LastName        string           `json:"lastName"`       
	PersonID        string           `json:"personId"`       
	TeamID          string           `json:"teamId"`         
	Jersey          string           `json:"jersey"`         
	IsActive        bool             `json:"isActive"`       
	Pos             PosEnum          `json:"pos"`            
	HeightFeet      string           `json:"heightFeet"`     
	HeightInches    string           `json:"heightInches"`   
	HeightMeters    HeightMetersEnum `json:"heightMeters"`   
	WeightPounds    string           `json:"weightPounds"`   
	WeightKilograms string           `json:"weightKilograms"`
	DateOfBirthUTC  string           `json:"dateOfBirthUTC"` 
	Teams           []Team           `json:"teams"`          
	Draft           Draft            `json:"draft"`          
	NbaDebutYear    string           `json:"nbaDebutYear"`   
	YearsPro        string           `json:"yearsPro"`       
	CollegeName     string           `json:"collegeName"`    
	LastAffiliation string           `json:"lastAffiliation"`
	Country         string           `json:"country"`        
}

type Draft struct {
	TeamID     string `json:"teamId"`    
	PickNum    string `json:"pickNum"`   
	RoundNum   string `json:"roundNum"`  
	SeasonYear string `json:"seasonYear"`
}

type Team struct {
	TeamID      string `json:"teamId"`     
	SeasonStart string `json:"seasonStart"`
	SeasonEnd   string `json:"seasonEnd"`  
}

type HeightMetersEnum string
const (
	HeightMeters HeightMetersEnum = ""
	The175 HeightMetersEnum = "1.75"
	The178 HeightMetersEnum = "1.78"
	The18 HeightMetersEnum = "1.8"
	The183 HeightMetersEnum = "1.83"
	The185 HeightMetersEnum = "1.85"
	The188 HeightMetersEnum = "1.88"
	The19 HeightMetersEnum = "1.9"
	The193 HeightMetersEnum = "1.93"
	The196 HeightMetersEnum = "1.96"
	The198 HeightMetersEnum = "1.98"
	The201 HeightMetersEnum = "2.01"
	The203 HeightMetersEnum = "2.03"
	The206 HeightMetersEnum = "2.06"
	The208 HeightMetersEnum = "2.08"
	The211 HeightMetersEnum = "2.11"
	The213 HeightMetersEnum = "2.13"
	The216 HeightMetersEnum = "2.16"
	The218 HeightMetersEnum = "2.18"
	The221 HeightMetersEnum = "2.21"
)

type PosEnum string
const (
	C PosEnum = "C"
	CF PosEnum = "C-F"
	F PosEnum = "F"
	FC PosEnum = "F-C"
	FG PosEnum = "F-G"
	G PosEnum = "G"
	GF PosEnum = "G-F"
	Pos PosEnum = ""
)

var rootURL string = "http://data.nba.net/10s/"

func main() {
	target := new(NBAInfoJSON)
	getNBAJSON(target)

	fmt.Println(target.Links.AllstarRoster)

	x := getAllPlayers();
	fmt.Println(x)
}

var client = &http.Client{Timeout: 10 * time.Second}

func getNBAJSON(target interface{}) {
	url := rootURL + "prod/v1/today.json"

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(target)
}

func getAllPlayers() LeagueRosterPlayers {
	url := rootURL + "prod/v1/2018/players.json"
	var result LeagueRosterPlayers

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(result)
	return result
}