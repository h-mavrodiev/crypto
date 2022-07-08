package gate

import "net/http"

type CurrencyChain []struct {
	Chain              string `json:"chain"`
	NameCn             string `json:"name_cn"`
	NameEn             string `json:"name_en"`
	IsDisabled         int    `json:"is_disabled"`
	IsDepositDisabled  int    `json:"is_deposit_disabled"`
	IsWithdrawDisabled int    `json:"is_withdraw_disabled"`
}

// Send Get reuquests to the List Chains Gate enpoint
func (c *GateClient) GetListChains(queryParam string, queryString string) (*CurrencyChain, error) {
	resource := "/currency_chains"
	req, err := c.CreateGetRequest(c.Endpoints.Wallet, resource, queryParam, queryString)
	if err != nil {
		return nil, err
	}
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

func (c *GateClient) GetWithdrawalRecords(queryParam string, queryString string) (*WithdrawalRecords, error) {
	resource := "/withdrawals"

	req, err := c.CreateGetRequest(c.Endpoints.Wallet, resource, queryParam, queryString)
	if err != nil {
		return nil, err
	}

	err = c.SignReq(req, http.MethodGet, c.Endpoints.Wallet, resource, req.URL.RawQuery, "")
	if err != nil {
		return nil, err
	}

	res := WithdrawalRecords{}
	if err = c.SendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type TotalBalance struct {
	Details Details `json:"details"`
	Total   Total   `json:"total"`
}

type Total struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Details struct {
	CrossMargin CrossMargin `json:"cross_margin"`
	Spot        Spot        `json:"spot"`
	Finance     Finance     `json:"finance"`
	Margin      Margin      `json:"margin"`
	Quant       Quant       `json:"quant"`
	Futures     Futures     `json:"futures"`
	Delivery    Delivery    `json:"delivery"`
	Warrant     Warrant     `json:"warrant"`
	Cbbc        Cbbc        `json:"cbbc"`
}

type CrossMargin struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Spot struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Finance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Margin struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Quant struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Futures struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Delivery struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Warrant struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Cbbc struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func (c *GateClient) GetTotalBalance(queryParam string, queryString string) (*TotalBalance, error) {
	resource := "/total_balance"

	req, err := c.CreateGetRequest(c.Endpoints.Wallet, resource, queryParam, queryString)
	if err != nil {
		return nil, err
	}

	err = c.SignReq(req, http.MethodGet, c.Endpoints.Wallet, resource, req.URL.RawQuery, "")
	if err != nil {
		return nil, err
	}

	res := TotalBalance{}
	if err = c.SendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
