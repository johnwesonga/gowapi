package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type CountryLatLong struct {
	Country string  `json:"country"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
	Name    string  `json:"name"`
	State   string  `json:"state"`
}

const (
	defaultBaseURL = "https://api.openweathermap.org/data/3.0/onecall"
	geoCodingURL   = "http://api.openweathermap.org/geo/1.0/direct"
)

func GetWeatherForCity(lat, long float32) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	url := fmt.Sprintf("%s?lat=%f&lon=%f&exclude=minutely,alerts,daily&appid=%s", defaultBaseURL, lat, long, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var oneCallResponse OneCallResponse

	if err := json.Unmarshal([]byte(body), &oneCallResponse); err != nil {
		panic(err)
	}
	fmt.Printf("Temp %f\nWEATHER\nDescription: %v", oneCallResponse.Current.Temp, oneCallResponse.Current.Weather[0].Description)
}

func GetGeoCords(city, countryCode string) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	url := fmt.Sprintf("%s?q=%s&limit=1&appid=%s", geoCodingURL, city, apiKey)
	if countryCode != "" {
		url = fmt.Sprintf("%s?q=%s,%s&limit=1&appid=%s", geoCodingURL, city, countryCode, apiKey)

	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var geoCordCallResponse GeoCordResponse

	if err := json.Unmarshal([]byte(body), &geoCordCallResponse); err != nil {
		panic(err)
	}

	fmt.Printf("Lat: %v\nLon: %v\n", geoCordCallResponse[0].Lat, geoCordCallResponse[0].Lon)

}
