package main

import (
	"log"
	"net/http"
	"os"
	openweather "weather-map-demo/components"
)

func main() {
	if os.Getenv("api_key") == "" {
		log.Fatal("ERROR: missing api_key")
		os.Exit(1)
	}

	if os.Getenv("base_url") == "" {
		log.Fatal("ERROR: missing base_url")
		os.Exit(1)
	}

	http.HandleFunc("/weather", openweather.Handler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
