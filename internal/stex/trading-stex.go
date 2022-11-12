package stex

import (
	"errors"
	"strconv"
)

func (c *StexClient) GetCurrencyPairFees(pair int, ch chan<- interface{}) error {
	resource := "/fees" + "/" + strconv.Itoa(pair)

	req, err := c.CreateGetRequest(c.Endpoints.Trading, resource, "", "")
	if err != nil {
		return errors.New("failed to create get request for stex currency pari fees")
	}

	res := CurrencyPairFees{}
	authenticate(c, req)
	if err = c.SendRequest(req, &res); err != nil {
		return errors.New("failed to send get request for stex currency pari fees")
	}

	ch <- res

	return nil
}
