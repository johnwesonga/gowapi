package main

import (
	"fmt"
	"github/johnwesonga/gowapi/api"
	"log"
)

const (
	geoCodingURL = "https://api.openweathermap.org/geo/1.0/direct"
	oneCallURL   = "https://api.openweathermap.org/data/3.0/onecall"
	mockURL      = "http://localhost:4444/geo"
)

func geoCodingExample() (lat, lon float32) {
	client := api.NewClient(geoCodingURL)
	geoCodingResp, err := client.GeoCodingSvc.GetGeoCordsByLocName("Concord", "USA")
	if err != nil {
		log.Fatal(err)
	}
	// Use geocoding response
	if len(geoCodingResp) > 0 {
		fmt.Printf("Lat: %v Lon: %v\n", geoCodingResp[0].Lat, geoCodingResp[0].Lon)
	} else {
		fmt.Println("No location found")
	}
	return geoCodingResp[0].Lat, geoCodingResp[0].Lon
}

func oneCallExample(lat, lon float32) {
	// do something
	c := api.NewClient(oneCallURL)
	weatherResp, err := c.OneCallSvc.GetCurrentWeatherForLoc(lat, lon)
	if err != nil {
		log.Fatal(err)
	}
	// use weather response
	fmt.Printf("Current Temp: %v\nHumidity: %v\nWeather:%v",
		weatherResp.Current.Temp,
		weatherResp.Current.Humidity,
		weatherResp.Current.Weather[0].Description)

}

func main() {
	lat, lon := geoCodingExample()
	oneCallExample(lat, lon)
}
