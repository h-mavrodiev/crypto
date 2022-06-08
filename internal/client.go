package client

import (
	"net/http"
	"time"
)

// Client struct
type Client struct {
	host       string
	prefix     string
	url        string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(host string, prefix string, apiKey string) *Client {
	return &Client{
		host:   host,
		prefix: prefix,
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
