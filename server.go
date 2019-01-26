package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dbrudner/go-nba/nba"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")

	router.HandleFunc("/standings", GetStandings).Methods("GET")
	router.HandleFunc("/stats", GetPlayerStats).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func GetStandings(w http.ResponseWriter, r *http.Request) {
	endpoints := nba.FetchTodayInfo()
	standings := new(nba.ConfStandings)
	nba.FetchNBADataJSON(endpoints.Links.LeagueConfStandings, standings)
	json.NewEncoder(w).Encode(standings)
}

func GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	playerName := r.URL.Query().Get("name")
	endpoints := nba.FetchTodayInfo()
	playerID, err := nba.FetchAllPlayersAndFindPlayerID(endpoints.Links.LeagueRosterPlayers, playerName)
	fmt.Print(playerID)
	if err != nil {
		fmt.Print("Error")
	}

	player := nba.FetchPlayerProfile(endpoints.Links.PlayerProfile, playerID)

	json.NewEncoder(w).Encode(player)
}
