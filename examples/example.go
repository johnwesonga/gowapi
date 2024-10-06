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

func geoCodingExample() {
	client := api.NewClient(geoCodingURL)
	geoCodingResp, err := client.GeoCodingSvc.GetGeoCordsByLocName("Nairobi", "KE")
	if err != nil {
		log.Fatal(err)
	}
	// Use geocoding response
	if len(geoCodingResp) > 0 {
		fmt.Printf("Lat: %v Lon: %v", geoCodingResp[0].Lat, geoCodingResp[0].Lon)
	} else {
		fmt.Println("No location found")
	}
}

func oneCallExample() {
	// do something
}

func main() {
	geoCodingExample()
	oneCallExample()
}
