package stex

import (
	"strconv"
	"sync"
)

type SafePrices struct {
	mu     sync.Mutex
	Prices Prices
}

func (p *SafePrices) updatePrices(o *safeOrderBook) {
	p.mu.Lock()
	p.Prices.CalcPriceAndVolume(&o.orderBook)
	p.mu.Unlock()
}
func (s *Prices) CalcPriceAndVolume(o *orderBook) {
	s.CalAskPricePerFixedAmount(o)
	s.CalBidPricePerFixedAmount(o)
}

func (s *Prices) CalAskPricePerFixedAmount(o *orderBook) {

	var minTradeVolume, wPrice, wUSDPriceSum, wUSDPrice, wPriceSum, sumVolume float64

	for _, order := range o.Ask {
		price, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			s.Sells = 8888.88
		}
		volume, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			s.SellsVolume = 9999.99
		}
		sumVolume += volume

		// USD/ETH * USD
		usdPrice := 1 / price
		if wUSDPrice == 0 {
			minTradeVolume = minTrade * usdPrice
		} else {
			minTradeVolume = minTrade * wUSDPrice
		}

		// In this case the amount is the weight for the price
		wPriceSum += price * volume
		wUSDPriceSum += usdPrice * volume

		wPrice = wPriceSum / sumVolume
		wUSDPrice = wUSDPriceSum / sumVolume

		if sumVolume >= minTradeVolume {
			s.Sells = wPrice
			s.SellsVolume = sumVolume
			break
		}

	}

}

func (s *Prices) CalBidPricePerFixedAmount(o *orderBook) {

	var minTradeVolume, wPrice, wUSDPriceSum, wUSDPrice, wPriceSum, sumVolume float64

	for _, order := range o.Bid {
		price, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			s.Buys = 8888.88
		}
		volume, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			s.BuysVolume = 9999.9
		}

		sumVolume += volume

		// USD/ETH * USD
		usdPrice := 1 / price
		if wUSDPrice == 0 {
			minTradeVolume = minTrade * usdPrice
		} else {
			minTradeVolume = minTrade * wUSDPrice
		}

		// In this case the amount is the weight for the price
		wPriceSum += price * volume
		wUSDPriceSum += usdPrice * volume

		wPrice = wPriceSum / sumVolume
		wUSDPrice = wUSDPriceSum / sumVolume

		if sumVolume >= minTradeVolume {
			s.Buys = wPrice
			s.BuysVolume = sumVolume
			break
		}

	}

}
