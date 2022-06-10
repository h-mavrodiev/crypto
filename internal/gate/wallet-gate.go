package gate

import (
	"net/http"
)

type CurrencyChain []struct {
	Chain              string `json:"chain"`
	NameCn             string `json:"name_cn"`
	NameEn             string `json:"name_en"`
	IsDisabled         int    `json:"is_disabled"`
	IsDepositDisabled  int    `json:"is_deposit_disabled"`
	IsWithdrawDisabled int    `json:"is_withdraw_disabled"`
}

// Send Get reuquest to the List Chains Gate enpoint
func (c *GateClient) GetListChains() (*CurrencyChain, error) {
	url := (c.Host + c.Prefix + c.Endpoints.Wallet + "/currency_chains")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := CurrencyChain{}
	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
