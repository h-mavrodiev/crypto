package stex

import (
	"crypto/configs"
	"fmt"
	"strconv"
	"sync"
)

type SafePrices struct {
	mu     sync.Mutex
	Prices Prices
}

func (p *SafePrices) updatePrices(o *safeOrderBook, errs chan<- error) {
	p.mu.Lock()
	p.Prices.CalcPriceAndVolume(&o.orderBook, errs)
	p.mu.Unlock()
}
func (s *Prices) CalcPriceAndVolume(o *orderBook, errs chan<- error) {
	s.CalAskPricePerFixedAmount(o, errs)
	s.CalBidPricePerFixedAmount(o, errs)
}

func (s *Prices) CalAskPricePerFixedAmount(o *orderBook, errs chan<- error) {
	var minTradeVolume, wPrice, wUSDPriceSum, wUSDPrice, wPriceSum, sumVolume float64

	for _, order := range o.Ask {
		price, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			errs <- fmt.Errorf("failed to parse Stex Sells price: %v", err)
			return
		}
		volume, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			errs <- fmt.Errorf("failed to parse Stex Sells volume: %v", err)
			return
		}
		sumVolume += volume

		// USD/ETH * USD
		usdPrice := 1 / price
		if wUSDPrice == 0 {
			minTradeVolume = configs.Conf.Stex.MinTrade * usdPrice
		} else {
			minTradeVolume = configs.Conf.Stex.MinTrade * wUSDPrice
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

func (s *Prices) CalBidPricePerFixedAmount(o *orderBook, errs chan<- error) {
	var minTradeVolume, wPrice, wUSDPriceSum, wUSDPrice, wPriceSum, sumVolume float64

	for _, order := range o.Bid {
		price, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			errs <- fmt.Errorf("failed to parse Stex Buys price: %v", err)
			return
		}
		volume, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			errs <- fmt.Errorf("failed to parse Stex Buys volume: %v", err)
			return
		}

		sumVolume += volume

		// USD/ETH * USD
		usdPrice := 1 / price
		if wUSDPrice == 0 {
			minTradeVolume = configs.Conf.Stex.MinTrade * usdPrice
		} else {
			minTradeVolume = configs.Conf.Stex.MinTrade * wUSDPrice
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
