package api

type OneCallService struct {
	client *Client
}

// get access to current weather, minute forecast for 1 hour, hourly forecast for 48 hours, daily forecast
// for 8 days and government weather alerts,
func GetCurrentWeatherForLoc(lat, long float32) {

}

// function to to get weather overview with a human-readable weather summary
// for today and tomorrow's forecast
func GetWeatherOverview(lat, long float32) {
}
