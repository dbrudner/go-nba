package main

// endpoint to retrieve a team schedule
// src: http://data.nba.net/prod/v1/2018/teams/1610612738/schedule.json
// date accessed: 2019-01-21 12:45:46.317

type Schedule struct {
	Internal Internal `json:"_internal"`
	League   struct {
		LastStandardGamePlayedIndex int64         `json:"lastStandardGamePlayedIndex"`
		Standard                    []Standard    `json:"standard"`
		Africa                      []interface{} `json:"africa"`
		Sacramento                  []interface{} `json:"sacramento"`
		Vegas                       []Standard    `json:"vegas"`
		Utah                        []interface{} `json:"utah"`
	}
}

type Standard struct {
	SeasonStageID     int64         `json:"seasonStageId"`
	GameID            string        `json:"gameId"`
	StatusNum         int64         `json:"statusNum"`
	ExtendedStatusNum int64         `json:"extendedStatusNum"`
	StartTimeUTC      string        `json:"startTimeUTC"`
	StartDateEastern  string        `json:"startDateEastern"`
	IsHomeTeam        bool          `json:"isHomeTeam"`
	Watch             Watch         `json:"watch"`
	Nugget            Nugget        `json:"nugget"`
	VTeam             Schedule_Team `json:"vTeam"`
	HTeam             Schedule_Team `json:"hTeam"`
}

type Schedule_Team struct {
	TeamID string `json:"teamId"`
	Score  string `json:"score"`
}

type Audio struct {
	National HTeamClass `json:"national"`
	VTeam    HTeamClass `json:"vTeam"`
	HTeam    HTeamClass `json:"hTeam"`
}

type HTeamClass struct {
	Streams      []Schedule_HTeamStream `json:"streams"`
	Broadcasters []struct {
		ShortName string `json:"shortName"`
		LongName  string `json:"longName"`
	} `json:"broadcasters"`
}

type Schedule_HTeamStream struct {
	Language Language `json:"language"`
	IsOnAir  bool     `json:"isOnAir"`
	StreamID string   `json:"streamId"`
}

type DeepLink struct {
	Broadcaster         Broadcaster `json:"broadcaster"`
	RegionalMarketCodes string      `json:"regionalMarketCodes"`
	IosApp              string      `json:"iosApp"`
	AndroidApp          string      `json:"androidApp"`
	DesktopWeb          string      `json:"desktopWeb"`
	MobileWeb           string      `json:"mobileWeb"`
}

type Language string

const (
	English Language = "English"
	Spanish Language = "Spanish"
)

type Broadcaster string

const (
	Tnt   Broadcaster = "TNT"
	Tntot Broadcaster = "TNTOT"
)

type StreamType string

const (
	Condensed StreamType = "condensed"
	HTeam     StreamType = "hTeam"
	Mobile    StreamType = "mobile"
	VTeam     StreamType = "vTeam"
)
