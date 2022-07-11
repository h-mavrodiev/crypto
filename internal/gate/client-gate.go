package gate

import (
	"crypto/configs"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// CreateGetReqeust creates GET http request
func (c *GateClient) CreateGetRequest(endpoint string, resource string, queryParam string, queryString string) (*http.Request, error) {
	urlStr := (c.Host + c.Prefix + endpoint + resource)

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

type errorResponse struct {
	Label   string `json:"label"`
	Message string `json:"message"`
}

func (c *GateClient) SendRequest(
	req *http.Request,
	target interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshal into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message + "\nError Label: " + errRes.Label)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)

	}

	if err = json.NewDecoder(res.Body).Decode(&target); err != nil {
		return err
	}

	return nil
}

// Client struct
type GateClient struct {
	Host          string
	Prefix        string
	Endpoints     configs.GateEndpoints
	CommonHeaders configs.GateCommonHeaders
	ApiKey        string
	ApiSecret     string
	HTTPClient    *http.Client
}

func NewClient(
	host string,
	prefix string,
	endpoints configs.GateEndpoints,
	headers configs.GateCommonHeaders,
	apiKey string,
	apiSecret string) *GateClient {

	client := &GateClient{
		Host:          host,
		Prefix:        prefix,
		Endpoints:     endpoints,
		CommonHeaders: headers,
		ApiKey:        apiKey,
		ApiSecret:     apiSecret,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
	return client
}
