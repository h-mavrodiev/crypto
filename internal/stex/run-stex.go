package stex

const (
	numRetry int     = 5
	minTrade float64 = 1
)

var (
	PriceInfo   SafePrices
	BalanceInfo SafeBalance = SafeBalance{Balance: Balance{}}
)

func (c *StexClient) RunStex(prices *SafePrices, balance *SafeBalance, errs chan error) {
	var ob safeOrderBook = safeOrderBook{orderBook: orderBook{}}

	go c.CallStexGetOrderBookDetails(&ob, prices, errs)
}
