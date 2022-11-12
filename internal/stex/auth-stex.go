package stex

import (
	"crypto/configs"
	"net/http"
)

func authenticate(c *StexClient, req *http.Request) {
	req.Header.Add("Authorization", "Bearer"+" "+configs.Conf.Stex.APIKey)
}
