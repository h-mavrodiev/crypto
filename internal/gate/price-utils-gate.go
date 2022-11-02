package gate

import (
	"strconv"
)

type GateInfo struct {
	AskPrice             float64
	AskDollarPrice       float64
	AskFixedDollarDemand float64
	AskAmount            float64
	BidPrice             float64
	BidDollarPrice       float64
	BidFixedDollarDemand float64
	BidAmount            float64
}

func (g *GateInfo) CalcPriceAndVolume(o OrderBookDetails, askFixedDollarDemand float64, bidFixedDollarDemand float64) {
	g.CalAskPricePerFixedAmount(o, askFixedDollarDemand)
	g.CalBidPricePerFixedAmount(o, bidFixedDollarDemand)
}

func (g *GateInfo) CalAskPricePerFixedAmount(o OrderBookDetails, askFixedDollarDemand float64) {

	var nextSpentAmount, spentAmount, price, dollarPrice float64

	for _, order := range o.Asks {
		p, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			g.AskPrice = 8888.88
		}
		price = p

		amount, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			g.AskAmount = 9999.99
		}

		dollarPrice = 1 / p
		nextSpentAmount += amount * dollarPrice

		if nextSpentAmount > askFixedDollarDemand {
			break
		} else {
			spentAmount = nextSpentAmount
		}
	}
	g.AskPrice = price
	g.AskDollarPrice = dollarPrice
	g.AskFixedDollarDemand = askFixedDollarDemand
	spentAmount = nextSpentAmount
	g.AskAmount = spentAmount
}

func (g *GateInfo) CalBidPricePerFixedAmount(o OrderBookDetails, bidFixedDollarDemand float64) {

	var nextSpentAmount, spentAmount, price, dollarPrice float64

	for _, order := range o.Bids {
		p, err := strconv.ParseFloat(order[0], 64)
		if err != nil {
			g.BidPrice = 8888.88
		}
		price = p

		amount, err := strconv.ParseFloat(order[1], 64)
		if err != nil {
			g.BidAmount = 9999.9
		}

		dollarPrice = 1 / price
		nextSpentAmount += amount * dollarPrice

		if nextSpentAmount > bidFixedDollarDemand {
			break
		} else {
			spentAmount = nextSpentAmount
		}
	}
	g.BidPrice = price
	g.BidDollarPrice = dollarPrice
	g.BidFixedDollarDemand = bidFixedDollarDemand
	spentAmount = nextSpentAmount
	g.BidAmount = spentAmount
}
