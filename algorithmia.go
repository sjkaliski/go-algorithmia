package algorithmia

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	errMissingFields = errors.New("missing field")
	algorithmiaURI   = "https://api.algorithmia.com/api/%s/%s"
)

// Response represents an Algorithmia API response.
type Response struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

// Client represents an Algorithmia API client.
type Client struct {
	token string
}

// NewClient returns a client.
func NewClient(token string) *Client {
	return &Client{token}
}

// Query sends a post request to the Algorithmia API to execute the specified
// algorithm with provided input.
func (c *Client) Query(user, algo string, input interface{}) (*Response, error) {
	if user == "" || algo == "" {
		return nil, errMissingFields
	}

	var data bytes.Buffer
	encoder := json.NewEncoder(&data)
	err := encoder.Encode(input)
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf(algorithmiaURI, user, algo)
	req, err := http.NewRequest("POST", addr, &data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.token)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response *Response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
