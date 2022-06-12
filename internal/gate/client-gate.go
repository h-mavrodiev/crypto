package gate

import (
	"crypto/configs"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (c *GateClient) SendRequest(req *http.Request, v interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshal into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)

	}

	// Unmarshall and populate v
	fullResponse := successResponse{
		Data: v,
	}

	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}

// Client struct
type GateClient struct {
	Host       string
	Prefix     string
	Endpoints  configs.GateEndpoints
	ApiKey     string
	HTTPClient *http.Client
}

func NewClient(host string, prefix string, apiKey string, endpoints configs.GateEndpoints) *GateClient {
	return &GateClient{
		Host:      host,
		Prefix:    prefix,
		Endpoints: endpoints,
		ApiKey:    apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
