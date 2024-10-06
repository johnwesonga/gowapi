package api

type OneCallService struct {
	client *Client
}

type OneCallResponse struct {
	Current Current `json:"current"`
}

type Current struct {
	Temp      float32 `json:"temp"`
	Humidity  int     `json:"humidity"`
	WindSpeed float32 `json:"wind_speed"`
	Weather   Weather `json:"weather"`
	Daily     []Daily `json:"daily"`
}
type Weather []struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Daily []struct {
	Summmary  string  `json:"summary"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	DewPoint  float32 `json:"dew_point"`
	WindSpeed float32 `json:"wind_speed"`
	WindDeg   int     `json:"wind_deg"`
	Weather   []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}
}

// get access to current weather, minute forecast for 1 hour, hourly forecast for 48 hours, daily forecast
// for 8 days and government weather alerts,
func GetCurrentWeatherForLoc(lat, long float32) {

}

// function to to get weather overview with a human-readable weather summary
// for today and tomorrow's forecast
func GetWeatherOverview(lat, long float32) {
}
