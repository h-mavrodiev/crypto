package stex

import (
	"errors"
	"strconv"
	"sync"
)

type CurrencyPairDetails []struct {
	ID                int    `json:"id"`
	CurrencyID        int    `json:"currency_id"`
	CurrencyCode      string `json:"currency_code"`
	CurrencyName      string `json:"currency_name"`
	MarketCurrencyID  int    `json:"market_currency_id"`
	MarketCode        string `json:"market_code"`
	MarketName        string `json:"market_name"`
	MinOrderAmount    string `json:"min_order_amount"`
	MinBuyPrice       string `json:"min_buy_price"`
	MinSellPrice      string `json:"min_sell_price"`
	BuyFeePercent     string `json:"buy_fee_percent"`
	SellFeePercent    string `json:"sell_fee_percent"`
	Active            bool   `json:"active"`
	Delisted          bool   `json:"delisted"`
	Message           string `json:"message"`
	CurrencyPrecision int    `json:"currency_precision"`
	MarketPrecision   int    `json:"market_precision"`
	Symbol            string `json:"symbol"`
	GroupName         string `json:"group_name"`
	GroupID           int    `json:"group_id"`
	AmountMultiplier  int    `json:"amount_multiplier"`
	TradingPrecision  int    `json:"trading_precision"`
}

func (c *StexClient) GetCurrencyPairDetails(pair int) (*CurrencyPairDetails, error) {
	resource := "/currency_pairs" + "/" + strconv.Itoa(pair)

	req, err := c.CreateGetRequest(c.Endpoints.Public, resource, "", "")
	if err != nil {
		return nil, err
	}

	res := CurrencyPairDetails{}
	if err = c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type OrderBookDetails struct {
	Ask            Ask     `json:"ask"`
	Bid            Bid     `json:"bid"`
	AskTotalAmount float64 `json:"ask_total_amount"`
	BidTotalAmount float64 `json:"bid_total_amount"`
}

type Ask []struct {
	CurrencyPairID   int     `json:"currency_pair_id"`
	Amount           string  `json:"amount"`
	Price            string  `json:"price"`
	Amount2          string  `json:"amount2"`
	Count            int     `json:"count"`
	CumulativeAmount float64 `json:"cumulative_amount"`
}

type Bid []struct {
	CurrencyPairID   int     `json:"currency_pair_id"`
	Amount           string  `json:"amount"`
	Price            string  `json:"price"`
	Amount2          string  `json:"amount2"`
	Count            int     `json:"count"`
	CumulativeAmount float64 `json:"cumulative_amount"`
}

func (c *StexClient) GetOrderBookDetails(pair int, ch chan<- interface{}, wg *sync.WaitGroup) error {
	// ETH-USDT code is 407
	defer wg.Done()

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
