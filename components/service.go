package openweather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetWeather calls the open weather api to get the info based on passed coordinates
func (o openWeather) GetWeather(lat, lon string) (*WeatherResponse, error) {
	url := fmt.Sprintf("%s?lat=%s&lon=%s&appid=%s&units=metric", o.baseUrl, lat, lon, o.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("requesting weather data failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get weather data: %s", resp.Status)
	}

	var weather WeatherResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body failed: %v", err)
	}

	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("parsing weather data failed: %v", err)
	}

	return &weather, nil
}
