package gate

import (
	"strconv"
)

type GateInfo struct {
	Sells       float64
	SellsVolume float64
	Buys        float64
	BuysVolume  float64
}

func (g *GateInfo) CalcPriceAndVolume(o OrderBookDetails, minTrade float64) {
	g.CalAskPricePerFixedAmount(o, minTrade)
	g.CalBidPricePerFixedAmount(o, minTrade)
}

func (g *GateInfo) CalAskPricePerFixedAmount(o OrderBookDetails, minTrade float64) {

	var minTradeVolume, wPrice, wUSDPriceSum, wUSDPrice, wPriceSum, sumVolume float64

	for _, order := range o.Asks {
		price, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			g.Sells = 8888.88
		}
		volume, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			g.SellsVolume = 9999.99
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
			g.Sells = wPrice
			g.SellsVolume = sumVolume
			break
		}

	}

}

func (g *GateInfo) CalBidPricePerFixedAmount(o OrderBookDetails, minTrade float64) {

	var minTradeVolume, wPrice, wUSDPriceSum, wUSDPrice, wPriceSum, sumVolume float64

	for _, order := range o.Bids {
		price, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			g.Buys = 8888.88
		}
		volume, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			g.BuysVolume = 9999.9
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
			g.Buys = wPrice
			g.BuysVolume = sumVolume
			break
		}

	}

}
