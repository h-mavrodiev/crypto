package stex

import "strconv"

type CurrencyPairFees struct {
	SellFee string `json:"sell_fee"`
	BuyFee  string `json:"buy_fee"`
}

func (c *StexClient) GetCurrencyPairFees(pair int) (*CurrencyPairFees, error) {
	resource := "/fees" + "/" + strconv.Itoa(pair)

	req, err := c.CreateGetRequest(c.Endpoints.Trading, resource, "", "")
	if err != nil {
		return nil, err
	}

	res := CurrencyPairFees{}
	c.Authenticate(req)
	if err = c.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
