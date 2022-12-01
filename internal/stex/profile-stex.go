package stex

import (
	"errors"
	"fmt"
)

func (c *StexClient) GetProfileWalletBalanceDetails(b *SafeBalance) error {

	resource := "wallets"

	req, err := c.CreateGetRequest(c.Endpoints.Profile, resource, "", "")
	if err != nil {
		return errors.New("failed create get request for stex profile wallet balance")
	}

	authenticate(c, req)

	res := ProfileBalance{}
	if err = c.SendRequest(req, &res); err != nil {
		return fmt.Errorf("\nfailed to send get request for stex profile wallet balance -> \n %s", err.Error())
	}
	b.updateBalanceFromHTTP(&res)

	return nil
}
