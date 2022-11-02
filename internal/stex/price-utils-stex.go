package stex

import (
	"strconv"
)

type StexInfo struct {
	AskPrice             float64
	AskDollarPrice       float64
	AskFixedDollarDemand float64
	AskAmount            float64
	BidPrice             float64
	BidDollarPrice       float64
	BidFixedDollarDemand float64
	BidAmount            float64
}

func (s *StexInfo) CalcPriceAndVolume(o OrderBookDetails, askFixedDollarDemand float64, bidFixedDollarDemand float64) {
	s.CalAskPricePerFixedAmount(o, askFixedDollarDemand)
	s.CalBidPricePerFixedAmount(o, bidFixedDollarDemand)
}

func (s *StexInfo) CalAskPricePerFixedAmount(o OrderBookDetails, askFixedDollarDemand float64) {

	var nextSpentAmount, spentAmount, price, dollarPrice float64

	for _, order := range o.Ask {
		p, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			s.AskPrice = 8888.88
		}
		price = p

		amount, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			s.AskAmount = 9999.99
		}

		dollarPrice = 1 / p
		nextSpentAmount += amount * dollarPrice

		if nextSpentAmount > askFixedDollarDemand {
			break
		} else {
			spentAmount = nextSpentAmount
		}
	}
	s.AskPrice = price
	s.AskDollarPrice = dollarPrice
	s.AskFixedDollarDemand = askFixedDollarDemand
	spentAmount = nextSpentAmount
	s.AskAmount = spentAmount
}

func (s *StexInfo) CalBidPricePerFixedAmount(o OrderBookDetails, bidFixedDollarDemand float64) {

	var nextSpentAmount, spentAmount, price, dollarPrice float64

	for _, order := range o.Bid {
		p, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			s.BidPrice = 8888.88
		}
		price = p

		amount, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			s.BidAmount = 9999.9
		}

		dollarPrice = 1 / price
		nextSpentAmount += amount * dollarPrice

		if nextSpentAmount > bidFixedDollarDemand {
			break
		} else {
			spentAmount = nextSpentAmount
		}
	}
	s.BidPrice = price
	s.BidDollarPrice = dollarPrice
	s.BidFixedDollarDemand = bidFixedDollarDemand
	spentAmount = nextSpentAmount
	s.BidAmount = spentAmount
}
