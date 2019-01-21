package main

// endpoint to retrieve a team schedule
// src: http://data.nba.net/10s/prod/v1/20171005/0011700027_boxscore.json
// date accessed: 2018-04-12 09:47:01.137

type Boxscore struct {
	Internal        Internal        `json:"_internal"`
	BasicGameData   BasicGameData   `json:"basicGameData"`
	PreviousMatchup PreviousMatchup `json:"previousMatchup"`
	Stats           Stats           `json:"stats"`
}

type BasicGameData struct {
	SeasonStageID         int64              `json:"seasonStageId"`
	SeasonYear            string             `json:"seasonYear"`
	GameID                string             `json:"gameId"`
	Arena                 Arena              `json:"arena"`
	IsGameActivated       bool               `json:"isGameActivated"`
	StatusNum             int64              `json:"statusNum"`
	ExtendedStatusNum     int64              `json:"extendedStatusNum"`
	StartTimeEastern      string             `json:"startTimeEastern"`
	StartTimeUTC          string             `json:"startTimeUTC"`
	EndTimeUTC            string             `json:"endTimeUTC"`
	StartDateEastern      string             `json:"startDateEastern"`
	Clock                 string             `json:"clock"`
	IsBuzzerBeater        bool               `json:"isBuzzerBeater"`
	IsPreviewArticleAvail bool               `json:"isPreviewArticleAvail"`
	IsRecapArticleAvail   bool               `json:"isRecapArticleAvail"`
	Tickets               Tickets            `json:"tickets"`
	HasGameBookPDF        bool               `json:"hasGameBookPdf"`
	IsStartTimeTBD        bool               `json:"isStartTimeTBD"`
	Nugget                Nugget             `json:"nugget"`
	Attendance            string             `json:"attendance"`
	GameDuration          GameDuration       `json:"gameDuration"`
	Period                Period             `json:"period"`
	VTeam                 BasicGameDataHTeam `json:"vTeam"`
	HTeam                 BasicGameDataHTeam `json:"hTeam"`
	Watch                 Watch              `json:"watch"`
	Officials             Officials          `json:"officials"`
}

type Arena struct {
	Name       string `json:"name"`
	IsDomestic bool   `json:"isDomestic"`
	City       string `json:"city"`
	StateAbbr  string `json:"stateAbbr"`
	Country    string `json:"country"`
}

type GameDuration struct {
	Hours   string `json:"hours"`
	Minutes string `json:"minutes"`
}

type BasicGameDataHTeam struct {
	TeamID     string      `json:"teamId"`
	TriCode    string      `json:"triCode"`
	Win        string      `json:"win"`
	Loss       string      `json:"loss"`
	SeriesWin  string      `json:"seriesWin"`
	SeriesLoss string      `json:"seriesLoss"`
	Score      string      `json:"score"`
	Linescore  []Linescore `json:"linescore"`
}

type Linescore struct {
	Score string `json:"score"`
}

type Nugget struct {
	Text string `json:"text"`
}

type Officials struct {
	Formatted []Formatted `json:"formatted"`
}

type Formatted struct {
	FirstNameLastName string `json:"firstNameLastName"`
}

type Period struct {
	Current       int64 `json:"current"`
	Type          int64 `json:"type"`
	MaxRegular    int64 `json:"maxRegular"`
	IsHalftime    bool  `json:"isHalftime"`
	IsEndOfPeriod bool  `json:"isEndOfPeriod"`
}

type Tickets struct {
	MobileApp  string `json:"mobileApp"`
	DesktopWeb string `json:"desktopWeb"`
	MobileWeb  string `json:"mobileWeb"`
}

type Watch struct {
	Broadcast Broadcast `json:"broadcast"`
}

type Broadcast struct {
	Broadcasters Broadcasters `json:"broadcasters"`
	Video        Video        `json:"video"`
	Audio        Audio        `json:"audio"`
}

type VTeamClass struct {
	Streams      []HTeamStream  `json:"streams"`
	Broadcasters []HTeamElement `json:"broadcasters"`
}

type HTeamElement struct {
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
}

type HTeamStream struct {
	Language string `json:"language"`
	IsOnAir  bool   `json:"isOnAir"`
	StreamID string `json:"streamId"`
}

type Broadcasters struct {
	National        []HTeamElement `json:"national"`
	Canadian        []interface{}  `json:"canadian"`
	VTeam           []interface{}  `json:"vTeam"`
	HTeam           []HTeamElement `json:"hTeam"`
	SpanishHTeam    []interface{}  `json:"spanish_hTeam"`
	SpanishVTeam    []interface{}  `json:"spanish_vTeam"`
	SpanishNational []interface{}  `json:"spanish_national"`
}

type Video struct {
	RegionalBlackoutCodes string        `json:"regionalBlackoutCodes"`
	CanPurchase           bool          `json:"canPurchase"`
	IsLeaguePass          bool          `json:"isLeaguePass"`
	IsNationalBlackout    bool          `json:"isNationalBlackout"`
	IsTNTOT               bool          `json:"isTNTOT"`
	IsVR                  bool          `json:"isVR"`
	TntotIsOnAir          bool          `json:"tntotIsOnAir"`
	Streams               []VideoStream `json:"streams"`
	DeepLink              []interface{} `json:"deepLink"`
}

type VideoStream struct {
	StreamType            string `json:"streamType"`
	IsOnAir               bool   `json:"isOnAir"`
	DoesArchiveExist      bool   `json:"doesArchiveExist"`
	IsArchiveAvailToWatch bool   `json:"isArchiveAvailToWatch"`
	StreamID              string `json:"streamId"`
	Duration              int64  `json:"duration"`
}

type Internal struct {
	PubDateTime string `json:"pubDateTime"`
	XSLT        string `json:"xslt"`
	EventName   string `json:"eventName"`
}

type PreviousMatchup struct {
	GameID   string `json:"gameId"`
	GameDate string `json:"gameDate"`
}

type Stats struct {
	TimesTied     string         `json:"timesTied"`
	LeadChanges   string         `json:"leadChanges"`
	VTeam         StatsHTeam     `json:"vTeam"`
	HTeam         StatsHTeam     `json:"hTeam"`
	ActivePlayers []ActivePlayer `json:"activePlayers"`
}

type ActivePlayer struct {
	PersonID  *string          `json:"personId,omitempty"`
	TeamID    *string          `json:"teamId,omitempty"`
	IsOnCourt *bool            `json:"isOnCourt,omitempty"`
	Points    string           `json:"points"`
	Pos       *Pos             `json:"pos,omitempty"`
	Min       string           `json:"min"`
	Fgm       string           `json:"fgm"`
	Fga       string           `json:"fga"`
	Fgp       string           `json:"fgp"`
	Ftm       string           `json:"ftm"`
	Fta       string           `json:"fta"`
	FTP       string           `json:"ftp"`
	TPM       string           `json:"tpm"`
	Tpa       string           `json:"tpa"`
	Tpp       string           `json:"tpp"`
	OffReb    string           `json:"offReb"`
	DefReb    string           `json:"defReb"`
	TotReb    string           `json:"totReb"`
	Assists   string           `json:"assists"`
	PFouls    string           `json:"pFouls"`
	Steals    string           `json:"steals"`
	Turnovers string           `json:"turnovers"`
	Blocks    string           `json:"blocks"`
	PlusMinus string           `json:"plusMinus"`
	Dnp       *string          `json:"dnp,omitempty"`
	SortKey   map[string]int64 `json:"sortKey,omitempty"`
}

type StatsHTeam struct {
	FastBreakPoints    string       `json:"fastBreakPoints"`
	PointsInPaint      string       `json:"pointsInPaint"`
	BiggestLead        string       `json:"biggestLead"`
	SecondChancePoints string       `json:"secondChancePoints"`
	PointsOffTurnovers string       `json:"pointsOffTurnovers"`
	LongestRun         string       `json:"longestRun"`
	Totals             ActivePlayer `json:"totals"`
	Leaders            Leaders      `json:"leaders"`
}

type Leaders struct {
	Points   Assists `json:"points"`
	Rebounds Assists `json:"rebounds"`
	Assists  Assists `json:"assists"`
}

type Assists struct {
	Value   string `json:"value"`
	Players []struct {
		PersonID string `json:"personId"`
	} `json:"players"`
}

type Pos string

const (
	C     Pos = "C"
	Empty Pos = ""
	PG    Pos = "PG"
	Pf    Pos = "PF"
	Sf    Pos = "SF"
	Sg    Pos = "SG"
)
