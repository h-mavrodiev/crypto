package gate

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	httpOBDelayTime     string = "1s"
	httpWalletDelaytime string = "1m"
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
		return fmt.Errorf("\nfailed to send get request for gate currency pair details ->\n %s", err.Error())
	}

	ch <- res

	return nil
}

func (c *GateClient) GetOrderBookDetails(o *safeOrderBook, p *SafePrices, errs chan error) error {

	resource := "/order_book"

	req, err := c.CreateGetRequest(c.Endpoints.Spot, resource, "currency_pair", c.Pair)
	if err != nil {
		return errors.New("failed to create get request for gate order book")
	}
	res := orderBook{}
	if err = c.SendRequest(req, &res); err != nil {
		// fmt.Errorf("failed to send get request for gate order book, %s", err.Error())
		return fmt.Errorf("\nfailed to send get request for gate order book ->\n %s", err.Error())
	}

	o.updateOrderBookFromHTTP(&res)
	p.updatePrices(o, errs)
	return nil
}

func (c *GateClient) CallGetOrderBookDetails(o *safeOrderBook, p *SafePrices, errs chan error) {
	counter := 0
	for {
		h, _ := time.ParseDuration(httpOBDelayTime)
		time.Sleep(h)
		err := c.GetOrderBookDetails(o, p, errs)
		if err != nil {
			counter++
			errs <- fmt.Errorf("HTTP Call Gate OB details fail: %v", err)
		}
		if counter > numRetry {
			errs <- errors.New("failed HTTP request for Gate OB details exceeded numRetry")
			return
		}
	}
}

func (c *GateClient) GetSpotWalletBalanceDetails(b *SafeBalance) error {
	resource := "/accounts"

	req, err := c.CreateGetRequest(c.Endpoints.Spot, resource, "", "")
	if err != nil {
		return errors.New("failed to create get request for gate spot balance")
	}

	err = c.signHTTPSReq(req, http.MethodGet, c.Endpoints.Spot, resource, req.URL.RawQuery, "")
	if err != nil {
		return errors.New("failed to sing the get request for gate spot balance")
	}

	res := SpotBalance{}
	if err = c.SendRequest(req, &res); err != nil {
		return fmt.Errorf("\nfailed to send get request for gate spot balance -> \n %s", err.Error())
	}
	b.updateBalanceFromHTTP(&res)

	return nil
}

func (c *GateClient) CallSpotWalletBalanceDetails(b *SafeBalance, errs chan error) {
	counter := 0
	for {
		h, _ := time.ParseDuration(httpWalletDelaytime)
		time.Sleep(h)
		err := c.GetSpotWalletBalanceDetails(b)
		if err != nil {
			counter++
			errs <- fmt.Errorf("HTTP Call Gate Spot Wallet fail: %v", err)
		}
		if counter > numRetry {
			errs <- errors.New("failed HTTP request for Gate Spot Wallet exceeded numRetry")
			return
		}
	}

}
