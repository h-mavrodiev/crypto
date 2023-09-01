package stex

import (
	"errors"
	"fmt"
	"time"
)

const httpWalletDelayTime string = "1m"

func (c *StexClient) GetProfileWalletBalanceDetails(b *SafeBalance) error {

	resource := "/wallets"

	req, err := c.CreateGetRequest(c.Endpoints.Profile, resource, "", "")
	if err != nil {
		return errors.New("failed create get request for stex profile wallet balance")
	}

	authenticate(c, req)

	res := balanceData{}
	if err = c.SendRequest(req, &res); err != nil {
		return fmt.Errorf("\nfailed to send get request for stex profile wallet balance -> \n %s", err.Error())
	}
	b.updateBalanceFromHTTP(&res)

	return nil
}

func (c *StexClient) CallStexGetProfileWalletBalanceDetails(b *SafeBalance, errs chan error) {
	counter := 0
	for {
		// ETH-USDT code is 407
		h, _ := time.ParseDuration(httpWalletDelayTime)
		time.Sleep(h)
		err := c.GetProfileWalletBalanceDetails(b)
		if err != nil {
			counter++
			errs <- fmt.Errorf("HTTP Call STEX wallet balance details fail: %v", err)
		}
		if counter > numRetry {
			errs <- errors.New("failed HTTP request for STEX wallet balance details exceeded numRetry")
			return
		}
	}
}
