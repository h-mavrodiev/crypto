package stex

import "log"

const (
	numRetry int = 5
)

var (
	PriceInfo   SafePrices
	BalanceInfo SafeBalance = SafeBalance{Balance: Balance{}}
)

func (c *StexClient) RunStex(prices *SafePrices, balance *SafeBalance, errs chan error) {
	var (
		ob  safeOrderBook = safeOrderBook{orderBook: orderBook{}}
		err error
	)

	// Initial call to obtain balance
	err = c.GetProfileWalletBalanceDetails(balance)
	if err != nil {
		log.Println(err)
		errs <- err
	}

	go c.CallStexGetOrderBookDetails(&ob, prices, errs)
	go c.CallStexGetProfileWalletBalanceDetails(balance, errs)
}
