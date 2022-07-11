package gate

import (
	"errors"
	"sync"
)

type CurrencyPairDetails struct {
	ID              string `json:"id"`
	Base            string `json:"base"`
	Quote           string `json:"quote"`
	Fee             string `json:"fee"`
	MinBaseAmount   string `json:"min_base_amount"`
	MinQuoteAmount  string `json:"min_quote_amount"`
	AmountPrecision int    `json:"amount_precision"`
	Precision       int    `json:"precision"`
	TradeStatus     string `json:"trade_status"`
	SellStart       int    `json:"sell_start"`
	BuyStart        int    `json:"buy_start"`
}

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

type OrderBookDetails struct {
	ID      int        `json:"id"`
	Current int64      `json:"current"`
	Update  int64      `json:"update"`
	Asks    [][]string `json:"asks"`
	Bids    [][]string `json:"bids"`
}

// func (c *GateClient) GetOrderBookDetails(pair string, ch chan<- interface{}, wg *sync.WaitGroup) error {
func (c *GateClient) GetOrderBookDetails(pair string, ch chan<- interface{}) error {

	// defer wg.Done()

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
