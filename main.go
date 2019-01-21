package main

import "fmt"

func main() {
	todayInfo := getTodayInfo()
	allPlayers := getAllPlayers()

	fmt.Println(todayInfo.Links.AllstarRoster)
	fmt.Println(allPlayers.Internal.PubDateTime)

	kd := getPlayerProfile("kevin durant", allPlayers.League.Standard)

	fmt.Println(kd.HeightFeet)
}
