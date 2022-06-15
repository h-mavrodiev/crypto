package gate

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// SignReq provides gateAPIv4 authentication headers to a request
// To check for how to add basic/AccessToken authentication visit the gate Go SDK and search for sign in the client.go file
func (c *GateClient) SignReq(
	req *http.Request,
	method string,
	apiRes string,
	queryString string,
	payloadString string) error {

	authURL := c.Host + c.Prefix + c.Endpoints.Wallet + apiRes
	requestUrl, err := url.Parse(authURL)
	if err != nil {
		return err
	}

	h := sha512.New()
	h.Write([]byte(payloadString))

	hashedPayload := hex.EncodeToString(h.Sum(nil))

	t := strconv.FormatInt(time.Now().Unix(), 10)

	m := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", method, requestUrl.Path, queryString, hashedPayload, t)

	mac := hmac.New(sha512.New, []byte(c.ApiSecret))

	mac.Write([]byte(m))
	sign := hex.EncodeToString(mac.Sum(nil))

	req.Header.Add("KEY", c.ApiKey)
	req.Header.Add("Timestamp", t)
	req.Header.Add("SIGN", sign)

	return nil
}
