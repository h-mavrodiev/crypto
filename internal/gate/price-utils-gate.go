package gate

import (
	"strconv"
)

type GateInfo struct {
	IWannaBuyFor                float64
	TheySellForWeightedUSD      float64
	ICanBuy                     float64
	TheySellForWeighted         float64
	ICanSellFromStex            float64
	ICanSellFromStexForWeighted float64
}

func (g *GateInfo) CalcPriceAndVolume(o OrderBookDetails, myUSDAmount float64, crossPlatfromMyAmount float64) {
	g.CalAskPricePerFixedAmount(o, myUSDAmount)
	g.CalBidPricePerFixedAmount(o, crossPlatfromMyAmount)
}

func (g *GateInfo) CalAskPricePerFixedAmount(o OrderBookDetails, iWannaBuyFor float64) {

	var iCanBuy, wUSDPrice, wPrice, wUSDPriceSum, wPriceSum, sumAmount float64

	for _, order := range o.Asks {
		price, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			g.TheySellForWeighted = 8888.88
		}
		amount, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			g.ICanBuy = 9999.99
		}

		// x*E = 20$ => x = 20*usdPrice
		usdPrice := 1 / price
		if wUSDPrice == 0 {
			iCanBuy = iWannaBuyFor * usdPrice
		} else {
			iCanBuy = iWannaBuyFor * wUSDPrice
		}

		// In this case the amount is the weight for the price
		sumAmount += amount
		wPriceSum += price * amount
		wUSDPriceSum += usdPrice * amount
		wPrice = wPriceSum / sumAmount
		wUSDPrice = wUSDPriceSum / sumAmount

		if iCanBuy < sumAmount {
			break
		}
	}

	g.IWannaBuyFor = iWannaBuyFor
	g.TheySellForWeightedUSD = wUSDPrice
	g.ICanBuy = iCanBuy
	g.TheySellForWeighted = wPrice
}

func (g *GateInfo) CalBidPricePerFixedAmount(o OrderBookDetails, MyAmountStex float64) {

	var wPrice, wPriceSum, sumAmount float64

	for _, order := range o.Bids {
		price, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			g.ICanSellFromStexForWeighted = 8888.88
		}
		amount, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			g.ICanSellFromStex = 9999.9
		}

		sumAmount += amount
		wPriceSum += price * amount
		wPrice = wPriceSum / sumAmount

		if sumAmount > MyAmountStex {
			break
		}
	}
	g.ICanSellFromStex = MyAmountStex
	g.ICanSellFromStexForWeighted = wPrice

}
