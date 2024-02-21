package openweather

type openWeather struct {
	apiKey  string
	baseUrl string
}

// Client is the gateway for http handlers
type Client interface {
	GetWeather(lat, lon string) (*WeatherResponse, error)
}

// NewClient provides the new instance of Client interface
func NewClient(apiKey, baseUrl string) (Client, error) {
	return &openWeather{
		apiKey:  apiKey,
		baseUrl: baseUrl,
	}, nil
}
