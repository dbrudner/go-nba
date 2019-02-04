// endpoint to retrieve a league schedule
// src: http://data.nba.net/10s/prod/v1/2018/schedule.json
// date accessed: 2019-02-03 17:42:06.962

package nba

import (
	"time"
)

type LeagueSchedule struct {
	Internal Internal `json:"_internal"`
	League   League   `json:"league"`
}

type League struct {
	Standard   []ScheduledGame `json:"standard"`
	Africa     []ScheduledGame `json:"africa"`
	Sacramento []ScheduledGame `json:"sacramento"`
	Vegas      []ScheduledGame `json:"vegas"`
	Utah       []ScheduledGame `json:"utah"`
}

type ScheduledGame struct {
	GameID            string               `json:"gameId"`
	SeasonStageID     int64                `json:"seasonStageId"`
	GameURLCode       string               `json:"gameUrlCode"`
	StatusNum         int64                `json:"statusNum"`
	ExtendedStatusNum int64                `json:"extendedStatusNum"`
	IsStartTimeTBD    bool                 `json:"isStartTimeTBD"`
	StartTimeUTC      string               `json:"startTimeUTC"`
	StartDateEastern  string               `json:"startDateEastern"`
	StartTimeEastern  string               `json:"startTimeEastern"`
	IsBuzzerBeater    bool                 `json:"isBuzzerBeater"`
	Period            LeagueSchedulePeriod `json:"period"`
	Nugget            Nugget               `json:"nugget"`
	Tags              []Tag                `json:"tags"`
	HTeam             Team                 `json:"hTeam"`
	VTeam             Team                 `json:"vTeam"`
	Watch             LeagueScheduleWatch  `json:"watch"`
}

type Team struct {
	TeamID string `json:"teamId"`
	Score  string `json:"score"`
	Win    string `json:"win"`
	Loss   string `json:"loss"`
}

type LeagueSchedulePeriod struct {
	Current    int64 `json:"current"`
	Type       int64 `json:"type"`
	MaxRegular int64 `json:"maxRegular"`
}

type LeagueScheduleWatch struct {
	Broadcast LeagueScheduleBroadcast `json:"broadcast"`
}

type LeagueScheduleBroadcast struct {
	Video LeagueScheduleVideo `json:"video"`
}

type LeagueScheduleVideo struct {
	RegionalBlackoutCodes string        `json:"regionalBlackoutCodes"`
	IsLeaguePass          bool          `json:"isLeaguePass"`
	IsNationalBlackout    bool          `json:"isNationalBlackout"`
	IsTNTOT               bool          `json:"isTNTOT"`
	CanPurchase           bool          `json:"canPurchase"`
	IsVR                  bool          `json:"isVR"`
	IsNextVR              bool          `json:"isNextVR"`
	IsNBAOnTNTVR          bool          `json:"isNBAOnTNTVR"`
	IsMagicLeap           bool          `json:"isMagicLeap"`
	IsOculusVenues        bool          `json:"isOculusVenues"`
	National              National      `json:"national"`
	Canadian              []Canadian    `json:"canadian"`
	SpanishNational       []interface{} `json:"spanish_national"`
}

type Canadian struct {
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
}

type National struct {
	Broadcasters []Canadian `json:"broadcasters"`
}

type Tag string

type SimpleGameSchedule struct {
	Away             string
	Home             string
	StartTimeEastern string
	StartTimeUTC     string
	StartDateEastern string
}

const (
	Africa       Tag = "AFRICA"
	Preseason    Tag = "PRESEASON"
	Slsacramento Tag = "SLSACRAMENTO"
	Slutah       Tag = "SLUTAH"
	Slvegas      Tag = "SLVEGAS"
)

func getTodayDate() string {
	return time.Now().Format("20060102")
}

func GetTodaySchedule() []ScheduledGame {
	todaySchedule := new(LeagueSchedule)
	todayInfo := FetchTodayInfo()
	todayDate := getTodayDate()
	FetchNBADataJSON(todayInfo.Links.LeagueSchedule, todaySchedule)

	schedule := []ScheduledGame{}

	for _, game := range todaySchedule.League.Standard {
		if game.StartDateEastern == todayDate {
			schedule = append(schedule, game)
		}
	}

	return schedule
}

func GetTodaySimpleSchedule() []SimpleGameSchedule {
	todaySchedule := new(LeagueSchedule)
	todayInfo := FetchTodayInfo()
	todayDate := getTodayDate()
	FetchNBADataJSON(todayInfo.Links.LeagueSchedule, todaySchedule)

	schedule := []SimpleGameSchedule{}

	for _, game := range todaySchedule.League.Standard {
		if game.StartDateEastern == todayDate {
			urlCode := game.GameURLCode
			team1 := urlCode[len(urlCode)-3:]
			team2 := urlCode[len(urlCode)-6 : len(urlCode)-3]

			simpleGame := SimpleGameSchedule{Away: team1,
				Home:             team2,
				StartTimeEastern: game.StartTimeEastern,
				StartTimeUTC:     game.StartTimeUTC,
				StartDateEastern: game.StartDateEastern,
			}

			schedule = append(schedule, simpleGame)
		}
	}
	return schedule
}
