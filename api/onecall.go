package api

import "fmt"

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
func (svc *OneCallService) GetCurrentWeatherForLoc(lat, lon float32) (OneCallResponse, error) {
	// https://api.openweathermap.org/data/3.0/onecall?lat={lat}&lon={lon}&exclude={part}&appid={API key}
	url := fmt.Sprintf("?lat=%f&lon=%f&exclude=minutely,alerts,daily", lat, lon)
	req, err := svc.client.NewRequest("GET", url)
	if err != nil {
		return OneCallResponse{}, err
	}
	var oneCallResponse OneCallResponse
	_, err = svc.client.Do(req, &oneCallResponse)
	if err != nil {
		return OneCallResponse{}, err
	}

	return oneCallResponse, nil
}

// function to to get weather overview with a human-readable weather summary
// for today and tomorrow's forecast
func (svc *OneCallService) GetWeatherOverview(lat, lon float32) (OneCallResponse, error) {
	//https://api.openweathermap.org/data/3.0/onecall/overview?lat={lat}&lon={lon}&appid={API key}
	url := fmt.Sprintf("/overview?lat=%f&lon=%f", lat, lon)

	req, err := svc.client.NewRequest("GET", url)
	if err != nil {
		return OneCallResponse{}, err
	}
	var oneCallResponse OneCallResponse
	_, err = svc.client.Do(req, &oneCallResponse)
	if err != nil {
		return OneCallResponse{}, err
	}

	return oneCallResponse, nil
}
