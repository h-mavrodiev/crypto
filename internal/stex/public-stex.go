package stex

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	httpOBDelayTime     string = "1.5s"
	httpWalletDelaytime string = "1m"
)

func (c *StexClient) GetCurrencyPairDetails(pair int, ch chan<- interface{}) error {
	resource := "/currency_pairs" + "/" + strconv.Itoa(pair)

	req, err := c.CreateGetRequest(c.Endpoints.Public, resource, "", "")
	if err != nil {
		return errors.New("faild create get request for stex pair details")
	}

	res := CurrencyPairDetails{}
	if err = c.SendRequest(req, &res); err != nil {
		return fmt.Errorf("\nfailed get request for stex pair details ->\n %s", err.Error())
	}
	ch <- res

	return nil
}

func (c *StexClient) GetOrderBookDetails(o *safeOrderBook, p *SafePrices) error {
	// ETH-USDT code is 407

	resource := "/orderbook" + "/" + strconv.Itoa(c.Pair)

	req, err := c.CreateGetRequest(c.Endpoints.Public, resource, "", "")
	if err != nil {
		return errors.New("failed create get request for stex order book")
	}

	res := orderBook{}
	if err = c.SendRequest(req, &res); err != nil {
		return fmt.Errorf("\nfailed get request for stex order book ->\n %s", err.Error())
	}

	o.updateOrderBookFromHTTP(&res)
	p.updatePrices(o)
	return nil
}

func (c *StexClient) CallStexGetOrderBookDetails(o *safeOrderBook, p *SafePrices, errs chan error) {
	counter := 0
	for {
		// ETH-USDT code is 407
		h, _ := time.ParseDuration(httpOBDelayTime)
		time.Sleep(h)
		err := c.GetOrderBookDetails(o, p)
		if err != nil {
			counter++
			errs <- fmt.Errorf("HTTP Call STEX OB details fail: %v", err)
		}
		if counter > numRetry {
			errs <- errors.New("failed HTTP request for STEX OB details exceeded numRetry")
			return
		}
	}
}
