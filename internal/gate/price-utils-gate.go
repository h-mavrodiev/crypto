package gate

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

func (g *Prices) CalcPriceAndVolume(o *orderBook, errs chan<- error) {
	g.CalAskPricePerFixedAmount(o, errs)
	g.CalBidPricePerFixedAmount(o, errs)
}

func (g *Prices) CalAskPricePerFixedAmount(o *orderBook, errs chan<- error) {
	var minTradeVolume, wPrice, wUSDPriceSum, wUSDPrice, wPriceSum, sumVolume float64

	for _, order := range o.Asks {
		price, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			errs <- fmt.Errorf("failed to parse Gate Sells price: %v", err)
			return
		}
		volume, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			errs <- fmt.Errorf("failed to parse Gate Sells volume: %v", err)
			return
		}
		sumVolume += volume

		// USD/ETH * USD
		usdPrice := 1 / price
		if wUSDPrice == 0 {
			minTradeVolume = configs.Conf.Gate.MinTrade * usdPrice
		} else {
			minTradeVolume = configs.Conf.Gate.MinTrade * wUSDPrice
		}

		// In this case the amount is the weight for the price
		wPriceSum += price * volume
		wUSDPriceSum += usdPrice * volume

		wPrice = wPriceSum / sumVolume
		wUSDPrice = wUSDPriceSum / sumVolume

		if sumVolume >= minTradeVolume {
			g.Sells = wPrice
			g.SellsVolume = sumVolume
			break
		}

	}

}

func (g *Prices) CalBidPricePerFixedAmount(o *orderBook, errs chan<- error) {
	var minTradeVolume, wPrice, wUSDPriceSum, wUSDPrice, wPriceSum, sumVolume float64

	for _, order := range o.Bids {
		price, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			errs <- fmt.Errorf("failed to parse Gate Buys price: %v", err)
			return
		}
		volume, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			errs <- fmt.Errorf("failed to parse Gate Buys volume: %v", err)
			return
		}

		sumVolume += volume

		// USD/ETH * USD
		usdPrice := 1 / price
		if wUSDPrice == 0 {
			minTradeVolume = configs.Conf.Gate.MinTrade * usdPrice
		} else {
			minTradeVolume = configs.Conf.Gate.MinTrade * wUSDPrice
		}

		// In this case the amount is the weight for the price
		wPriceSum += price * volume
		wUSDPriceSum += usdPrice * volume

		wPrice = wPriceSum / sumVolume
		wUSDPrice = wUSDPriceSum / sumVolume

		if sumVolume >= minTradeVolume {
			g.Buys = wPrice
			g.BuysVolume = sumVolume
			break
		}

	}

}
