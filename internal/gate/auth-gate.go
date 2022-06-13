package gate

import (
	"fmt"
	"net/http"
	"time"
)

type SignHeader struct {
	Key       string `json:"KEY"`
	Timestamp string `json:"Timestamps"`
	Sign      string `json:"SIGN"`
}

func (c *GateClient) SignReq(
	req *http.Request,
	method string,
	apiRes string,
	queryString string,
	payloadString string) error {

	timestamp := time.Now().String()
	//Request Method + "\n" + Request URL + "\n" + Query String + "\n" + HexEncode(SHA512(Request Payload)) + "\n" + Timestamp
	s := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", method, (c.Prefix + c.Endpoints.Wallet + apiRes), queryString, payloadString, timestamp)
	heads := &SignHeader{
		Key:       c.ApiKey,
		Timestamp: timestamp,
		Sign:      s,
	}

	req.Header.Add("KEY", heads.Key)
	req.Header.Add("Timestamp", heads.Timestamp)
	req.Header.Add("SIGN", heads.Sign)
	return nil

}
