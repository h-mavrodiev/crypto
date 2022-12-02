package gate

import (
	"crypto/configs"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// SignReq provides gateAPIv4 authentication headers to a request
// To check for how to add basic/AccessToken authentication visit the gate Go SDK and search for sign in the client.go file
func (c *GateClient) signHTTPSReq(
	req *http.Request,
	method string,
	endpoint string,
	apiRes string,
	queryString string,
	payloadString string) error {

	var (
		secret  string = configs.Conf.Gate.APISecret
		key     string = configs.Conf.Gate.APIKey
		authURL string = c.Host + c.Prefix + endpoint + apiRes
	)

	requestUrl, err := url.Parse(authURL)
	if err != nil {
		return err
	}

	h := sha512.New()
	h.Write([]byte(payloadString))

	hashedPayload := hex.EncodeToString(h.Sum(nil))

	t := strconv.FormatInt(time.Now().Unix(), 10)

	m := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", method, requestUrl.Path, queryString, hashedPayload, t)

	mac := hmac.New(sha512.New, []byte(secret))

	mac.Write([]byte(m))
	sign := hex.EncodeToString(mac.Sum(nil))

	req.Header.Add("KEY", key)
	req.Header.Add("Timestamp", t)
	req.Header.Add("SIGN", sign)

	return nil
}

func sign(channel, event string, t int64) string {
	message := fmt.Sprintf("channel=%s&event=%s&time=%d", channel, event, t)
	h2 := hmac.New(sha512.New, []byte(configs.Conf.Gate.APISecret))
	io.WriteString(h2, message)
	return hex.EncodeToString(h2.Sum(nil))
}

func (msg *WSMsg) signWSMsg() {
	signStr := sign(msg.Channel, msg.Event, msg.Time)
	msg.Auth = &Auth{
		Method: "api_key",
		KEY:    configs.Conf.Gate.APISecret,
		SIGN:   signStr,
	}
}
