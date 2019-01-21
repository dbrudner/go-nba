package main

import (
	"flag"
	"fmt"
)

func main() {
	player := flag.String("player", "", "Search for a player by full name (case insensitive)")
	flag.Parse()
	fmt.Println(*player)

	todayInfo := getTodayInfo()

	// getAllPlayersEndpoint := todayInfo.Links.LeagueRosterPlayers

	// allPlayers := getAllPlayers(getAllPlayersEndpoint)

	// playerID, err := getPlayerID(*player, allPlayers.League.Standard)
	// if err != nil {
	// 	panic(err)
	// }

	// playerProfile := getPlayerProfile(playerID)

	// fmt.Println(playerProfile.League.Standard.Stats.Latest.Ppg)

	teams := getAllTeams(todayInfo.Links.Teams)

	teamID, err := getTeamID("nyk", teams.League.Standard)
	if err != nil {
		panic(err)
	}

	fmt.Print(teamID)
}
