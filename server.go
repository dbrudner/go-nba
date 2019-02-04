package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/dbrudner/go-nba/nba"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	tmpl := template.Must(template.ParseGlob("./templates/*.tmpl.html"))
	port := os.Getenv("PORT")

	// API routes
	router.HandleFunc("/standings", getStandings).Methods("GET")
	router.HandleFunc("/stats", getPlayerStats).Methods("GET")
	router.HandleFunc("/teams", getTeams).Methods("GET")
	router.HandleFunc("/today-schedule", getTodaySchedule).Methods("GET")
	router.HandleFunc("/today-simple-schedule", getTodaySimpleSchedule).Methods("GET")

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

	if err != nil {
		fmt.Print("Error")
	}

	player := nba.FetchPlayerProfile(endpoints.Links.PlayerProfile, playerID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

type getTeamsResponse struct {
	Teams []nba.Teams_Team
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	endpoints := nba.FetchTodayInfo()
	teams := nba.FetchAllNBATeams(endpoints.Links.Teams)
	response := getTeamsResponse{Teams: teams}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type getTodayScheduleResponse struct {
	TodaySchedule []nba.ScheduledGame
}

func getTodaySchedule(w http.ResponseWriter, r *http.Request) {
	todaySchedule := nba.GetTodaySchedule()
	fmt.Printf("%T", todaySchedule)
	response := getTodayScheduleResponse{TodaySchedule: todaySchedule}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type GetTodaySimpleScheduleResponse struct {
	TodaySchedule []nba.SimpleGameSchedule
}

func getTodaySimpleSchedule(w http.ResponseWriter, r *http.Request) {
	todaySchedule := nba.GetTodaySimpleSchedule()
	response := GetTodaySimpleScheduleResponse{TodaySchedule: todaySchedule}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	fmt.Print(response)
}
