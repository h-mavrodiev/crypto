package gate

import (
	"errors"
	"net/http"
)

// Send Get reuquests to the List Chains Gate enpoint
func (c *GateClient) GetListChains(queryParam string, queryString string, ch chan<- interface{}) error {
	resource := "/currency_chains"
	req, err := c.CreateGetRequest(c.Endpoints.Wallet, resource, queryParam, queryString)
	if err != nil {
		return errors.New("failed to create get request for gate list currency chains")
	}

	res := CurrencyChain{}
	if err := c.SendRequest(req, &res); err != nil {
		return errors.New("failed to send get request for gate list currency chains")
	}

	ch <- res

	return nil
}

func (c *GateClient) GetWithdrawalRecords(queryParam string, queryString string, ch chan<- interface{}) error {
	resource := "/withdrawals"

	req, err := c.CreateGetRequest(c.Endpoints.Wallet, resource, queryParam, queryString)
	if err != nil {
		return errors.New("failed to create get request for gate withdrawals")
	}

	err = c.signHTTPSReq(req, http.MethodGet, c.Endpoints.Wallet, resource, req.URL.RawQuery, "")
	if err != nil {
		return errors.New("failed to sing the get request for gate withdrawals")
	}

	res := WithdrawalRecords{}
	if err = c.SendRequest(req, &res); err != nil {
		return errors.New("failed to send get request for gate withdrawals")
	}

	ch <- res

	return nil
}

func (c *GateClient) GetTotalBalance(queryParam string, queryString string, ch chan<- *TotalBalance) error {

	resource := "/total_balance"

	req, err := c.CreateGetRequest(c.Endpoints.Wallet, resource, queryParam, queryString)
	if err != nil {
		return errors.New("failed to create get request for gate total balance")
	}

	err = c.signHTTPSReq(req, http.MethodGet, c.Endpoints.Wallet, resource, req.URL.RawQuery, "")
	if err != nil {
		return errors.New("failed to sign get request for gate total balance")
	}

	res := TotalBalance{}
	if err = c.SendRequest(req, &res); err != nil {
		return errors.New("failed to send get request for gate total balance")
	}

	ch <- &res

	return nil
}
