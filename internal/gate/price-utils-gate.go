package gate

import (
	"strconv"
)

type GateInfo struct {
	AskWeightedPrice    float64
	AskFixedUSDDemand   float64
	AskWeightedUSDPrice float64
	AskAmount           float64
	BidWeightedPrice    float64
	BidFixedUSDDemand   float64
	BidWeightedUSDPrice float64
	BidAmount           float64
}

func (g *GateInfo) CalcPriceAndVolume(o OrderBookDetails, askFixedUSDDemand float64, bidFixedUSDDemand float64) {
	g.CalAskPricePerFixedAmount(o, askFixedUSDDemand)
	g.CalBidPricePerFixedAmount(o, bidFixedUSDDemand)
}

func (g *GateInfo) CalAskPricePerFixedAmount(o OrderBookDetails, askFixedUSDDemand float64) {

	var nextSpentAmount, spentAmount, wPriceSum, usdPrice, wUSDPriceSum, sumAmount float64

	for _, order := range o.Asks {
		p, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			g.AskWeightedPrice = 8888.88
		}
		amount, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			g.AskAmount = 9999.99
		}

		wPriceSum += p * amount

		// next 2 rows below are the same as amount / price but is confusing AF....
		usdPrice = 1 / p
		wUSDPriceSum += usdPrice * amount

		sumAmount += amount
		nextSpentAmount += amount * usdPrice

		if nextSpentAmount > askFixedUSDDemand {
			break
		} else {
			spentAmount = nextSpentAmount
		}
	}

	g.AskWeightedPrice = wPriceSum / sumAmount
	g.AskFixedUSDDemand = askFixedUSDDemand
	g.AskWeightedUSDPrice = wUSDPriceSum / sumAmount
	spentAmount = nextSpentAmount
	g.AskAmount = spentAmount
}

func (g *GateInfo) CalBidPricePerFixedAmount(o OrderBookDetails, bidFixedUSDDemand float64) {

	var nextSpentAmount, spentAmount, wPriceSum, usdPrice, wUSDPriceSum, sumAmount float64

	for _, order := range o.Bids {
		p, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			g.BidWeightedPrice = 8888.88
		}
		amount, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			g.BidAmount = 9999.9
		}

		wPriceSum += p * amount

		// next 2 rows below are the same as amount / price but is confusing AF....
		usdPrice = 1 / p
		wUSDPriceSum += usdPrice * amount

		sumAmount += amount
		nextSpentAmount += amount * usdPrice

		if nextSpentAmount > bidFixedUSDDemand {
			break
		} else {
			spentAmount = nextSpentAmount
		}
	}

	g.BidWeightedPrice = wPriceSum / sumAmount
	g.BidFixedUSDDemand = bidFixedUSDDemand
	g.BidWeightedUSDPrice = wUSDPriceSum / sumAmount
	spentAmount = nextSpentAmount
	g.BidAmount = spentAmount
}
