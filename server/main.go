package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dbrudner/go-nba/nba"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/standings", GetStandings).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetStandings(w http.ResponseWriter, r *http.Request) {
	endpoints := nba.GetTodayInfo()
	standings := new(nba.ConfStandings)
	nba.GetNBAJSON(endpoints.Links.LeagueConfStandings, standings)
	json.NewEncoder(w).Encode(standings)
}
