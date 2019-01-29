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

	router.HandleFunc("/standings", getStandings).Methods("GET")
	router.HandleFunc("/stats", getPlayerStats).Methods("GET")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "indexPage", http.StatusInternalServerError)
	})
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Broken???
func getStandings(w http.ResponseWriter, r *http.Request) {
	endpoints := nba.FetchTodayInfo()
	standings := new(nba.ConfStandings)
	nba.FetchNBADataJSON(endpoints.Links.LeagueConfStandings, standings)
	json.NewEncoder(w).Encode(standings)
}

func getPlayerStats(w http.ResponseWriter, r *http.Request) {
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
