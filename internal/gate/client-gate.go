package gate

import (
	"crypto/configs"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
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

func (c *GateClient) SendRequest(
	req *http.Request,
	target interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	log.Printf("|| %s || %s || -> GATE request ... \n", res.Request.Method, res.Status)

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

func NewClient(
	host string,
	prefix string,
	endpoints configs.GateEndpoints,
	headers configs.GateCommonHeaders,
	pair string,
	apiKey string,
	apiSecret string) (*GateClient, error) {

	stop := false
	retry := 0
	var conn *websocket.Conn
	var maxRetry int = 10

	for !stop {
		dialer := websocket.DefaultDialer
		u := url.URL{Scheme: configs.Conf.Gate.WS.Schema,
			Host: configs.Conf.Gate.WS.WSHost,
			Path: configs.Conf.Gate.WS.Path}
		c, _, err := dialer.Dial(u.String(), nil)
		if err != nil {
			if retry >= maxRetry {
				log.Printf("max reconnect time %d reached, give it up", maxRetry)
				return nil, err
			}
			retry++
			log.Printf("failed to connect to server for the %d time, try again later", retry)
			time.Sleep(time.Millisecond * (time.Duration(retry) * 500))
			continue
		} else {
			stop = true
			conn = c
		}
	}

	if retry > 0 {
		log.Printf("reconnect succeeded after retrying %d times", retry)
	}

	client := &GateClient{
		Host:          host,
		Prefix:        prefix,
		Endpoints:     endpoints,
		CommonHeaders: headers,
		Pair:          pair,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},

		WSConn: conn,
	}
	return client, nil
}
