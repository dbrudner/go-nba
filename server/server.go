package main

import (
	"flag"
	"fmt"

	"github.com/dbrudner/go-nba/nba"
)

func main() {
	player := flag.String("player", "", "Search for a player by full name (case insensitive)")
	flag.Parse()
	fmt.Println(*player)

	endpoints := nba.GetTodayInfo()

	fmt.Print(endpoints)

	// playerID, err := getPlayerID(*player, allPlayers.League.Standard)
	// if err != nil {
	// 	panic(err)
	// }

	// playerProfile := getPlayerProfile(playerID)

	// fmt.Println(playerProfile.League.Standard.Stats.Latest.Ppg)

	// teams := nba.getAllPlayers(nba.getAllTeams(todayInfo.Links.Teams))

	// teamID, err := nba.getTeamID("nyk", teams.League.Standard)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Print(teamID)
}