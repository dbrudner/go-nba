// endpoint to retrieve today's endpoints for NBA stats API
// all endpoints should be used from here _NOT_ hardcoded
// src: http://data.nba.net/10s/prod/v1/today.json
// date accessed: 2019-01-21 12:00:00.47

package nba

type NBAInfoJSON struct {
	Internal Internal `json:"_internal"`
	Links    struct {
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

func FetchTodayInfo() NBAInfoJSON {
	NBAInfo := new(NBAInfoJSON)
	FetchNBAJSON("/prod/v1/today.json", NBAInfo)
	return *NBAInfo
}
