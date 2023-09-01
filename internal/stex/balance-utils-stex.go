package stex

import "sync"

type SafeBalance struct {
	mu      sync.Mutex
	Balance Balance
}

// updateBalance takes wallet update from HTTP request and updates Balance object
func (b *SafeBalance) updateBalanceFromHTTP(spotBalance *balanceData) {
	b.mu.Lock()
	for _, profB := range *spotBalance {
		switch {
		case profB.CurrencyCode == "BTC":
			b.Balance.BTC = profB.Balance
		case profB.CurrencyCode == "ETH":
			b.Balance.ETH = profB.Balance
		case profB.CurrencyCode == "USDT":
			b.Balance.USDT = profB.Balance
		}
	}
	b.mu.Unlock()
}
