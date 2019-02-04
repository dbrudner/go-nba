// endpoint to retrieve a league schedule
// src: http://data.nba.net/10s/prod/v1/2018/schedule.json
// date accessed: 2019-02-03 17:42:06.962

package nba

import (
	"fmt"
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
	StartTimeEastern  StartTimeEastern     `json:"startTimeEastern"`
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

type StartTimeEastern string

const (
	The1000PmEt StartTimeEastern = "10:00 PM ET"
	The100PmEt  StartTimeEastern = "1:00 PM ET"
	The1030PmEt StartTimeEastern = "10:30 PM ET"
	The1100AmEt StartTimeEastern = "11:00 AM ET"
	The1100PmEt StartTimeEastern = "11:00 PM ET"
	The1130PmEt StartTimeEastern = "11:30 PM ET"
	The1200PmEt StartTimeEastern = "12:00 PM ET"
	The1230PmEt StartTimeEastern = "12:30 PM ET"
	The200PmEt  StartTimeEastern = "2:00 PM ET"
	The300PmEt  StartTimeEastern = "3:00 PM ET"
	The330PmEt  StartTimeEastern = "3:30 PM ET"
	The400PmEt  StartTimeEastern = "4:00 PM ET"
	The430PmEt  StartTimeEastern = "4:30 PM ET"
	The500PmEt  StartTimeEastern = "5:00 PM ET"
	The530PmEt  StartTimeEastern = "5:30 PM ET"
	The600PmEt  StartTimeEastern = "6:00 PM ET"
	The630PmEt  StartTimeEastern = "6:30 PM ET"
	The700PmEt  StartTimeEastern = "7:00 PM ET"
	The730AmEt  StartTimeEastern = "7:30 AM ET"
	The730PmEt  StartTimeEastern = "7:30 PM ET"
	The800AmEt  StartTimeEastern = "8:00 AM ET"
	The800PmEt  StartTimeEastern = "8:00 PM ET"
	The830PmEt  StartTimeEastern = "8:30 PM ET"
	The900PmEt  StartTimeEastern = "9:00 PM ET"
	The930PmEt  StartTimeEastern = "9:30 PM ET"
)

type Tag string

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
			fmt.Println(game.StartDateEastern)
			fmt.Println(game.StartTimeEastern)
			schedule = append(schedule, game)
		}
	}

	return schedule
}
