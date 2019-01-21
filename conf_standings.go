package main

// Structs for endpoint to retrieve conference standings
// src: http://data.nba.net/prod/v1/current/standings_conference.json
// date accessed: 2019-01-21 10:03:41.77

type ConfStandings struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		XSLT        string `json:"xslt"`
		EventName   string `json:"eventName"`
	}
	League struct {
		Standard Standard `json:"standard"`
	}
}

type Standard struct {
	SeasonYear    int64      `json:"seasonYear"`
	SeasonStageID int64      `json:"seasonStageId"`
	Conference    Conference `json:"conference"`
}

type Conferences struct {
	East []Conference `json:"east"`
	West []Conference `json:"west"`
}

type Conference struct {
	TeamID                 string           `json:"teamId"`
	Win                    string           `json:"win"`
	Loss                   string           `json:"loss"`
	WinPct                 string           `json:"winPct"`
	WinPctV2               string           `json:"winPctV2"`
	LossPct                string           `json:"lossPct"`
	LossPctV2              string           `json:"lossPctV2"`
	GamesBehind            string           `json:"gamesBehind"`
	DivGamesBehind         string           `json:"divGamesBehind"`
	ClinchedPlayoffsCode   string           `json:"clinchedPlayoffsCode"`
	ClinchedPlayoffsCodeV2 string           `json:"clinchedPlayoffsCodeV2"`
	ConfRank               string           `json:"confRank"`
	ConfWin                string           `json:"confWin"`
	ConfLoss               string           `json:"confLoss"`
	DivWin                 string           `json:"divWin"`
	DivLoss                string           `json:"divLoss"`
	HomeWin                string           `json:"homeWin"`
	HomeLoss               string           `json:"homeLoss"`
	AwayWin                string           `json:"awayWin"`
	AwayLoss               string           `json:"awayLoss"`
	LastTenWin             string           `json:"lastTenWin"`
	LastTenLoss            string           `json:"lastTenLoss"`
	Streak                 string           `json:"streak"`
	DivRank                string           `json:"divRank"`
	IsWinStreak            bool             `json:"isWinStreak"`
	TieBreakerPts          string           `json:"tieBreakerPts"`
	SortKey                map[string]int64 `json:"sortKey"`
}
