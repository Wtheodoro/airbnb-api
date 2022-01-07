package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Live anywhere model
type LiveAnywhere struct {
	ID       int    `json:"id"`
	Image    string `json:"img"`
	Title string `json:"title"`
}

// Explore places model
type ExplorePlaces struct {
	ID       int    `json:"id"`
	Image    string `json:"img"`
	Location string `json:"location"`
	Distance string `json:"distance"`
}

func getLiveAnywhere(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]LiveAnywhere{
		{
			ID: 1,
			Image: "https://links.papareact.com/2io",
			Title: "Outdoor getaways",
		},
		{
			ID: 2,
			Image: "https://links.papareact.com/q7j",
			Title: "Unique stays",
		},
		{
			ID: 3,
			Image: "https://links.papareact.com/s03",
			Title: "Entire homes",
		},
		{
			ID: 4,
			Image: "https://links.papareact.com/8ix",
			Title: "Pet allowed",
		},
	})
}

func getExploreNearby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]ExplorePlaces{
		{
			ID:       1,
			Image:    "https://links.papareact.com/5j2",
			Location: "London",
			Distance: "45-minute drive",
		},
		{
			ID:       2,
			Image:    "https://links.papareact.com/1to",
			Location: "Manchester",
			Distance: "4.5-hour drive",
		},
		{
			ID:       3,
			Image:    "https://links.papareact.com/40m",
			Location: "Liverpool",
			Distance: "4.5-hour drive",
		},
		{
			ID:       4,
			Image:    "https://links.papareact.com/msp",
			Location: "York",
			Distance: "4-hour drive",
		},
		{
			ID:       5,
			Image:    "https://links.papareact.com/2k3",
			Location: "Cardiff",
			Distance: "45-minute drive",
		},
		{
			ID:       6,
			Image:    "https://links.papareact.com/ynx",
			Location: "Birkenhead",
			Distance: "4.5-hour drive",
		},
		{
			ID:       7,
			Image:    "https://links.papareact.com/kji",
			Location: "Newquay",
			Distance: "6-hour drive",
		},
		{
			ID:       8,
			Image:    "https://links.papareact.com/41m",
			Location: "Hove",
			Distance: "2-hour drive",
		},
	})
}

func getHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Try those endpoints: /api/explorePlaces")
}

func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/api/explorePlaces", getExploreNearby).Methods("GET")
	router.HandleFunc("/api/liveAnywhere", getLiveAnywhere).Methods("GET")
	router.HandleFunc("/", getHomePage).Methods("GET")

	fmt.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
