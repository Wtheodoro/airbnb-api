package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SearchResults struct {
	ID       int    `json:"id"`
	Image    string `json:"img"`
	Location string `json:"location"`
	Title string `json:"title"`
	Description string `json:"description"`
	Star float32 `json:"star"`
	Price string `json:"price"`
	Total string `json:"total"`
	Long float64 `json:"long"`
	Lat float64 `json:"lat"`
}

type LiveAnywhere struct {
	ID       int    `json:"id"`
	Image    string `json:"img"`
	Title string `json:"title"`
}

type ExplorePlaces struct {
	ID       int    `json:"id"`
	Image    string `json:"img"`
	Location string `json:"location"`
	Distance string `json:"distance"`
}

// TODO: Implement DB

func getSearchResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]SearchResults{
		{
			ID: 0,
			Image: "https://links.papareact.com/xqj",
			Location: "Private room in center of London",
			Title: "Stay at this spacious Edwardian House",
			Description: "1 guest · 1 bedroom · 1 bed · 1.5 shared bthrooms · Wifi · Kitchen · Free parking · Washing Machine",
			Star: 4.73,
			Price: "£30 / night",
			Total: "£117 total",
			Long: -0.0022275,
			Lat: 51.5421655,
		},
		{
			ID: 1,
			Image: "https://links.papareact.com/hz2",
			Location: "Private room in center of London",
			Title: "Independant luxury studio apartment",
			Description: "2 guest · 3 bedroom · 1 bed · 1.5 shared bthrooms · Wifi · Kitchen",
			Star: 4.3,
			Price: "£40 / night",
			Total: "£157 total",
			Long: -0.095091,
			Lat: 51.48695,
		},
		{
			ID: 2,
			Image: "https://links.papareact.com/uz7",
			Location: "Private room in center of London",
			Title: "London Studio Apartments",
			Description: "4 guest · 4 bedroom · 4 bed · 2 bathrooms · Free parking · Washing Machine",
			Star: 3.8,
			Price: "£35 / night",
			Total: "£207 total",
			Long: -0.135638,
			Lat: 51.521916,
		},
		{
			ID: 3,
			Image: "https://links.papareact.com/6as",
			Location: "Private room in center of London",
			Title: "30 mins to Oxford Street, Excel London",
			Description: "1 guest · 1 bedroom · 1 bed · 1.5 shared bthrooms · Wifi · Kitchen · Free parking · Washing Machine",
			Star: 4.1,
			Price: "£55 / night",
			Total: "£320 total",
			Long: -0.069961,
			Lat: 51.472618,
		},
		{
			ID: 4,
			Image: "https://links.papareact.com/xhc",
			Location: "Private room in center of London",
			Title: "Spacious Peaceful Modern Bedroom",
			Description: "3 guest · 1 bedroom · 1 bed · 1.5 shared bthrooms · Wifi · Free parking · Dry Cleaning",
			Star: 5.0,
			Price: "£60 / night",
			Total: "£450 total",
			Long: -0.08472,
			Lat: 51.510794,
		},
		{
			ID: 5,
			Image: "https://links.papareact.com/pro",
			Location: "Private room in center of London",
			Title: "The Blue Room In London",
			Description: "2 guest · 1 bedroom · 1 bed · 1.5 shared bthrooms · Wifi · Washing Machine",
			Star: 4.23,
			Price: "£65 / night",
			Total: "£480 total",
			Long: -0.094116,
			Lat: 51.51401,
		},
		{
			ID: 6,
			Image: "https://links.papareact.com/8w2",
			Location: "Private room in center of London",
			Title: "5 Star Luxury Apartment",
			Description: "3 guest · 1 bedroom · 1 bed · 1.5 shared bthrooms · Wifi · Kitchen · Free parking · Washing Machine",
			Star: 3.85,
			Price: "£90 / night",
			Total: "£650 total",
			Long: -0.109889,
			Lat: 51.521245,
		},
	})
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

	router.HandleFunc("/api/searchResults", getSearchResults).Methods("GET")
	router.HandleFunc("/api/explorePlaces", getExploreNearby).Methods("GET")
	router.HandleFunc("/api/liveAnywhere", getLiveAnywhere).Methods("GET")
	router.HandleFunc("/", getHomePage).Methods("GET")

	fmt.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
