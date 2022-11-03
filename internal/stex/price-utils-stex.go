package stex

import (
	"strconv"
)

type StexInfo struct {
	AskWeightedPrice    float64
	AskFixedUSDDemand   float64
	AskWeightedUSDPrice float64
	AskAmount           float64
	BidWeightedPrice    float64
	BidFixedUSDDemand   float64
	BidWeightedUSDPrice float64
	BidAmount           float64
}

func (s *StexInfo) CalcPriceAndVolume(o OrderBookDetails, askFixedUSDDemand float64, bidFixedUSDDemand float64) {
	s.CalAskPricePerFixedAmount(o, askFixedUSDDemand)
	s.CalBidPricePerFixedAmount(o, bidFixedUSDDemand)
}

func (s *StexInfo) CalAskPricePerFixedAmount(o OrderBookDetails, askFixedUSDDemand float64) {

	var nextSpentAmount, spentAmount, wPriceSum, usdPrice, wUSDPriceSum, sumAmount float64

	for _, order := range o.Ask {
		p, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			s.AskWeightedPrice = 8888.88
		}
		amount, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			s.AskAmount = 9999.99
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

	s.AskWeightedPrice = wPriceSum / sumAmount
	s.AskFixedUSDDemand = askFixedUSDDemand
	s.AskWeightedUSDPrice = wUSDPriceSum / sumAmount
	spentAmount = nextSpentAmount
	s.AskAmount = spentAmount
}

func (s *StexInfo) CalBidPricePerFixedAmount(o OrderBookDetails, bidFixedUSDDemand float64) {

	var nextSpentAmount, spentAmount, wPriceSum, usdPrice, wUSDPriceSum, sumAmount float64

	for _, order := range o.Bid {
		p, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			s.BidWeightedPrice = 8888.88
		}
		amount, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			s.BidAmount = 9999.9
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

	s.BidWeightedPrice = wPriceSum / sumAmount
	s.BidFixedUSDDemand = bidFixedUSDDemand
	s.BidWeightedUSDPrice = wUSDPriceSum / sumAmount
	spentAmount = nextSpentAmount
	s.BidAmount = spentAmount
}
