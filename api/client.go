package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	BaseURL       *url.URL
	client        *http.Client
	defaultApiKey string
	//Services
	GeoCodingSvc *GeoCodingService
}

func NewClient(defaultBaseURL string) *Client {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENWEATHER_API_KEY must be set")
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:        http.DefaultClient,
		BaseURL:       baseURL,
		defaultApiKey: apiKey}
	c.GeoCodingSvc = &GeoCodingService{client: c}
	return c
}

func (c *Client) NewRequest(method, path string) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)

	url := fmt.Sprintf("%s&appid=%s", u, c.defaultApiKey)
	log.Println(url)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	return req, nil
}
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	log.Printf("Received response with status: %s\n", resp.Status)

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}
