package stex

import (
	"crypto/configs"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

// CreateGetReqeust creates GET http request
func (c *StexClient) CreateGetRequest(endpoint string, resource string, queryParam string, queryString string) (*http.Request, error) {

	urlStr := (c.Host + endpoint + resource)

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", c.CommonHeaders.Accept)
	req.Header.Add("Content-Type", c.CommonHeaders.ContentType)
	q := req.URL.Query()
	q.Add(queryParam, queryString)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (c *StexClient) SendRequest(req *http.Request, v interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	log.Printf("|| %s || %s || -> STEX request ... \n", res.Request.Method, res.Status)

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

func NewClient(
	host string,
	apiKey string,
	endpoints configs.StexEndpoints,
	headers configs.StexCommonHeaders,
	pair int) *StexClient {
	return &StexClient{
		Host:          host,
		Endpoints:     endpoints,
		CommonHeaders: headers,
		Pair:          pair,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
