package main

import "fmt"

func main() {
	todayInfo := getTodayInfo()
	allPlayers := getAllPlayers()

	fmt.Println(todayInfo.Links.AllstarRoster)
	fmt.Println(allPlayers.Internal.PubDateTime)

	kdID := getPlayerID("kevin durant", allPlayers.League.Standard)
	kdProfile := getPlayerProfile(kdID)

	fmt.Println(kdProfile.League.Standard.Stats.Latest.Ppg)
}
