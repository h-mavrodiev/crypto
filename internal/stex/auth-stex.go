package stex

import "net/http"

func (c *StexClient) Authenticate(req *http.Request) {
	req.Header.Add("Authorization", "Bearer"+" "+c.ApiKey)
}
