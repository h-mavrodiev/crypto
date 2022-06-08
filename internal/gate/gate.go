package gate

import (
	"net/http"
	"net/url"

	"github.com/h-mavrodiev/crypto/client"
)

const (
	host   string = "https://api.gateio.ws"
	prefix string = "/api/v4"
	apiKey string = "placeholder"
)

func (c *client.Client) GetHandler() {
	resp, err := http.Get(host)
	if err != nil {
	}
	defer resp.Body.Close()
}

func PostHandler() {
	resp, err := http.Post(host + prefix)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
}

func PostFormHandler() {
	resp, err := http.PostForm("http://example.com/form",
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
}
