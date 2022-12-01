package gate

import "sync"

type SafeBalance struct {
	mu      sync.Mutex
	Balance Balance
}

// updateBalance takes wallet update from HTTP request and updates Balance object
func (b *SafeBalance) updateBalanceFromHTTP(spotBalance *SpotBalance) {
	b.mu.Lock()
	for _, spotB := range *spotBalance {
		switch {
		case spotB.Currency == "BTC":
			b.Balance.BTC = spotB.Available
		case spotB.Currency == "ETH":
			b.Balance.ETH = spotB.Available
		case spotB.Currency == "USDT":
			b.Balance.USDT = spotB.Available
		}

	}
	b.mu.Unlock()
}

// updateBalanceFromWS takes wallet update notification and updates Balance object
func (b *SafeBalance) updateBalanceFromWS(wsBallance *WalletUpdateNotification) {
	b.mu.Lock()
	for _, spotB := range *wsBallance {
		switch {
		case spotB.Currency == "BTC":
			b.Balance.BTC = spotB.Available
		case spotB.Currency == "ETH":
			b.Balance.ETH = spotB.Available
		case spotB.Currency == "USDT":
			b.Balance.USDT = spotB.Available
		}

	}
	b.mu.Unlock()
}
