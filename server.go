package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dbrudner/go-nba/nba"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/standings", GetStandings).Methods("GET")
	router.HandleFunc("/stats/{name}", GetPlayerStats).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetStandings(w http.ResponseWriter, r *http.Request) {
	endpoints := nba.FetchTodayInfo()
	standings := new(nba.ConfStandings)
	nba.FetchNBASON(endpoints.Links.LeagueConfStandings, standings)
	json.NewEncoder(w).Encode(standings)
}

func GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// playerName = vars["name"]
	// fmt.Print(playerName)
	endpoints := nba.FetchTodayInfo()
	playerID, err := nba.FetchAllPlayersAndFindPlayerID(endpoints.Links.LeagueRosterPlayers, "Stephen Curry")
	if err != nil {
		fmt.Print("Error")
	}

	player := nba.FetchPlayerProfile(endpoints.Links.PlayerProfile, playerID)

	json.NewEncoder(w).Encode(player)
}
