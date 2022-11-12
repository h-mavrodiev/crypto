package stex

import (
	"errors"
	"strconv"
)

func (c *StexClient) GetCurrencyPairDetails(pair int, ch chan<- interface{}) error {
	resource := "/currency_pairs" + "/" + strconv.Itoa(pair)

	req, err := c.CreateGetRequest(c.Endpoints.Public, resource, "", "")
	if err != nil {
		return errors.New("faild create get request for stex pair details")
	}

	res := CurrencyPairDetails{}
	if err = c.SendRequest(req, &res); err != nil {
		return errors.New("failed get request for stex pair details")
	}
	ch <- res

	return nil
}

func (c *StexClient) GetOrderBookDetails(pair int, ch chan<- OrderBookDetails) error {
	// ETH-USDT code is 407

	resource := "/orderbook" + "/" + strconv.Itoa(pair)

	req, err := c.CreateGetRequest(c.Endpoints.Public, resource, "", "")
	if err != nil {
		return errors.New("faild create get request for stex order book")
	}

	res := OrderBookDetails{}
	if err = c.SendRequest(req, &res); err != nil {
		return errors.New("failed get request for stex order book")
	}

	ch <- res

	return nil
}
