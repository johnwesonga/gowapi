package api

import "fmt"

type GeoCodingService struct {
	client *Client
}

type GeoCordResponse []struct {
	Name    string  `json:"name"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

func (s *GeoCodingService) GetGeoCordsByLocName(city, countryCode string) (GeoCordResponse, error) {
	url := fmt.Sprintf("?q=%s,%s&limit=1", city, countryCode)
	req, err := s.client.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}
	var geoCordCallResponse GeoCordResponse
	_, err = s.client.Do(req, &geoCordCallResponse)
	if err != nil {
		return nil, err
	}

	return geoCordCallResponse, nil
}

func (s *GeoCodingService) GetGeoCordsByZip(zip, countryCode string) (GeoCordResponse, error) {
	url := fmt.Sprintf("?q=%s,%s&limit=1", zip, countryCode)
	req, err := s.client.NewRequest("GET", url)
	if err != nil {
		return nil, err
	}
	var geoCordCallResponse GeoCordResponse
	_, err = s.client.Do(req, &geoCordCallResponse)
	if err != nil {
		return nil, err
	}

	return geoCordCallResponse, nil
}
