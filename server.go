package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/dbrudner/go-nba/nba"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type PageData struct {
}

// our main function
func main() {
	router := mux.NewRouter()
	tmpl := template.Must(template.ParseGlob("./templates/*.tmpl.html"))
	port := os.Getenv("PORT")

	// API routes
	router.HandleFunc("/standings", getStandings).Methods("GET")
	router.HandleFunc("/stats", getPlayerStats).Methods("GET")
	router.HandleFunc("/teams", getTeams).Methods("GET")

	// HTML routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "indexPage", http.StatusInternalServerError)
	})
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsObj)(router)))
}

func getStandings(w http.ResponseWriter, r *http.Request) {
	endpoints := nba.FetchTodayInfo()
	standings := new(nba.ConfStandings)
	nba.FetchNBADataJSON(endpoints.Links.LeagueConfStandings, standings)
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

type GetTeamsResponse struct {
	Teams []nba.Teams_Team
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	endpoints := nba.FetchTodayInfo()
	teams := nba.FetchAllNBATeams(endpoints.Links.Teams)
	response := GetTeamsResponse{Teams: teams}
	fmt.Printf("%T", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
