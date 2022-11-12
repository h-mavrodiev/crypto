package gate

import (
	"errors"
	"sync"
)

func (c *GateClient) GetCurrencyPairDetails(pair string, ch chan<- interface{}, wg *sync.WaitGroup) error {
	defer wg.Done()

	resource := "/currency_pairs" + "/" + pair

	req, err := c.CreateGetRequest(c.Endpoints.Spot, resource, "", "")
	if err != nil {
		return errors.New("faild create get request for gate currency pair details")
	}

	res := CurrencyPairDetails{}
	if err = c.SendRequest(req, &res); err != nil {
		return errors.New("failed to send get request for gate currency pair details")
	}

	ch <- res

	return nil
}

func (c *GateClient) GetOrderBookDetails(pair string, ch chan<- OrderBookDetails) error {

	resource := "/order_book"

	req, err := c.CreateGetRequest(c.Endpoints.Spot, resource, "currency_pair", pair)
	if err != nil {
		return errors.New("faild create get request for gate order book")
	}

	res := OrderBookDetails{}
	if err = c.SendRequest(req, &res); err != nil {
		return errors.New("failed to send get request for gate order book")
	}

	ch <- res
	return nil
}
