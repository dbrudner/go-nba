package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/dbrudner/go-nba/nba"
	"github.com/gorilla/mux"
)

type PageData struct {
}

// our main function
func main() {
	router := mux.NewRouter()
	tmpl := template.Must(template.ParseGlob("./templates/*.tmpl.html"))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// API routes
	router.HandleFunc("/api/standings", GetStandings).Methods("GET")
	router.HandleFunc("/api/stats", GetPlayerStats).Methods("GET")

	// HTML routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "indexPage", http.StatusInternalServerError)
	})

	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Broken???
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
