package stex

import (
	"strconv"
)

type StexInfo struct {
	IWannaBuyFor                float64
	TheySellForWeightedUSD      float64
	ICanBuy                     float64
	TheySellForWeighted         float64
	ICanSellFromGate            float64
	ICanSellFromGateForWeighted float64
}

func (s *StexInfo) CalcPriceAndVolume(o OrderBookDetails, askFixedUSDDemand float64, bidFixedUSDDemand float64) {
	s.CalAskPricePerFixedAmount(o, askFixedUSDDemand)
	s.CalBidPricePerFixedAmount(o, bidFixedUSDDemand)
}

func (s *StexInfo) CalAskPricePerFixedAmount(o OrderBookDetails, iWannaBuyFor float64) {

	var iCanBuy, wUSDPrice, wPrice, wUSDPriceSum, wPriceSum, sumAmount float64

	for _, order := range o.Ask {
		price, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			s.TheySellForWeighted = 8888.88
		}
		amount, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			s.ICanBuy = 9999.99
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

	s.IWannaBuyFor = iWannaBuyFor
	s.TheySellForWeightedUSD = wUSDPrice
	s.ICanBuy = iCanBuy
	s.TheySellForWeighted = wPrice
}

func (s *StexInfo) CalBidPricePerFixedAmount(o OrderBookDetails, MyAmountGate float64) {

	var wPrice, wPriceSum, sumAmount float64

	for _, order := range o.Bid {
		price, err := strconv.ParseFloat(order.Price, 64)
		if err != nil {
			s.ICanSellFromGateForWeighted = 8888.88
		}
		amount, err := strconv.ParseFloat(order.Amount, 64)
		if err != nil {
			s.ICanSellFromGate = 9999.9
		}

		sumAmount += amount
		wPriceSum += price * amount
		wPrice = wPriceSum / sumAmount

		if sumAmount > MyAmountGate {
			break
		}
	}
	s.ICanSellFromGate = MyAmountGate
	s.ICanSellFromGateForWeighted = wPrice
}
