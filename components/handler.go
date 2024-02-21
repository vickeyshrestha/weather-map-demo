package openweather

import (
	"encoding/json"
	"net/http"
	"os"
)

// Handler processes incoming HTTP requests and generate responses
func Handler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters for latitude and longitude
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" || lon == "" {
		http.Error(w, "missing latitude and longitude as query parameters", http.StatusBadRequest)
		return
	}

	// Call OpenWeather API
	client, err := NewClient(os.Getenv("api_key"), os.Getenv("base_url"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weather, err := client.GetWeather(lat, lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Determine temperature description
	tempDesc := "moderate"
	if weather.Main.Temp < 283.15 { // Less than 10°C
		tempDesc = "cold"
	} else if weather.Main.Temp > 303.15 { // More than 30°C
		tempDesc = "hot"
	}

	// Construct response
	response := map[string]string{
		"weather_condition": weather.Weather[0].Main,
		"temperature":       tempDesc,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
