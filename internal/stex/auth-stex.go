package stex

import (
	"crypto/configs"
	"net/http"
)

func authenticate(c *StexClient, req *http.Request) {
	var key string = configs.Conf.Stex.APIKey

	req.Header.Add("Authorization", "Bearer"+" "+key)
}
