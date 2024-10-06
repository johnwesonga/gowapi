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
	OneCallSvc   *OneCallService
}

type ErrorResponse struct {
	Response   *http.Response `json:"-"`
	Cod        int            `json:"cod"`
	Message    string         `json:"message"`
	Parameters []string       `json:"parameters,omitempty"`
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
		defaultApiKey: apiKey,
	}
	c.GeoCodingSvc = &GeoCodingService{client: c}
	c.OneCallSvc = &OneCallService{client: c}
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

	err = CheckResponse(resp)
	if err != (nil) {
		defer resp.Body.Close()
		return nil, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("Error\nMessage: %v", r.Message)
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range or equal to 202 Accepted.
// API error responses are expected to have response
// body, and a JSON response body that maps to ErrorResponse.
func CheckResponse(r *http.Response) error {
	if r.StatusCode == http.StatusAccepted {
		return nil
	}
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}

	if err := json.NewDecoder(r.Body).Decode(errorResponse); err != nil {
		return err
	}

	return errorResponse

}
