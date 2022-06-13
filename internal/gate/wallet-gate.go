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
func (c *GateClient) GetListChains(query string) (*CurrencyChain, error) {
	url := (c.Host + c.Prefix + c.Endpoints.Wallet + "/currency_chains" + "?" + query)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res := CurrencyChain{}
	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type WithdrawalRecords []struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Currency  string `json:"currency"`
	Address   string `json:"address"`
	Txid      string `json:"txid"`
	Amount    string `json:"amount"`
	Memo      string `json:"memo"`
	Status    string `json:"status"`
	Chain     string `json:"chain"`
}

func (c *GateClient) GetRetrieveWithdrawalRecords(query string) (*WithdrawalRecords, error) {
	resourse := "/withdrawals"
	url := (c.Host + c.Prefix + c.Endpoints.Wallet + resourse)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", c.CommonHeaders.Accept)
	req.Header.Add("Content-Type", c.CommonHeaders.ContetnType)

	if err := c.SignReq(req, http.MethodGet, resourse, "", ""); err != nil {
		return nil, err
	}

	res := WithdrawalRecords{}
	if err := c.SendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
