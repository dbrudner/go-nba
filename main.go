package main

import "fmt"

func main() {
	NBAInfo := new(NBAInfoJSON)
	allPlayers := new(LeagueRosterPlayers)

	getNBAJSON("prod/v1/today.json", NBAInfo)
	getNBAJSON("prod/v1/2018/players.json", allPlayers)

	fmt.Println(NBAInfo.Links.AllstarRoster)
	fmt.Println(allPlayers.Internal.PubDateTime)

	kd := getPlayer("kevin durant", allPlayers.League.Standard)

	fmt.Println(kd.HeightFeet)
}
