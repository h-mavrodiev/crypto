package gate

import (
	"fmt"
	"time"
)

type SignHeader struct {
	Key       string `json:"KEY"`
	Timestamp string `json:"Timestamps"`
	Sign      string `json:"SIGN"`
}

func (c *GateClient) GenSign(ep string,
	method string,
	queryString string,
	payloadString string) (*SignHeader, error) {

	timestamp := time.Now().String()
	//Request Method + "\n" + Request URL + "\n" + Query String + "\n" + HexEncode(SHA512(Request Payload)) + "\n" + Timestamp
	s := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", method, (c.Host + c.Prefix + ep), queryString, payloadString, timestamp)

	return &SignHeader{Key: c.ApiKey,
		Timestamp: timestamp,
		Sign:      s}, nil

}
