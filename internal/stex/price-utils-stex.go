package stex

import (
	"fmt"
	"strconv"
	"sync"
)

type StexInfo struct {
	AskPrice  string
	AskAmount string
	BidPrice  string
	BidAmount string
}

func (s *StexInfo) UpdateStexVisualizeInfo(askPrice string, askAmount string, bidPrice string, bidAmount string) {
	s.AskPrice = askPrice
	s.AskAmount = askAmount
	s.BidPrice = bidPrice
	s.BidAmount = bidAmount
}

func (o *OrderBookDetails) FindPriceAndVolume(volumeByPrice float64) *StexInfo {

	var (
		priceOrderAsk       string
		sumOrdersVolumesAsk float64
		priceOrderBid       string
		sumOrderVolumeBid   float64
	)
	wg := sync.WaitGroup{}

	for sumOrdersVolumesAsk < volumeByPrice {

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for _, order := range o.Ask {
				s, err := strconv.ParseFloat(order.Amount, 64)
				if err != nil {
					priceOrderAsk = "Error Parsing Price"
					break
				}
				if sumOrdersVolumesAsk < volumeByPrice {
					sumOrdersVolumesAsk = sumOrdersVolumesAsk + s
				}
				priceOrderAsk = order.Price
			}
		}(&wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for _, order := range o.Bid {
				s, err := strconv.ParseFloat(order.Amount, 64)
				if err != nil {
					priceOrderBid = "Error Parsing Price"
					break
				}
				if sumOrdersVolumesAsk < volumeByPrice {
					sumOrderVolumeBid = sumOrderVolumeBid + s
				}
				priceOrderBid = order.Price
			}
		}(&wg)
	}
	wg.Wait()

	return &StexInfo{
		AskPrice:  priceOrderAsk,
		AskAmount: fmt.Sprintf("%f", sumOrdersVolumesAsk),
		BidPrice:  priceOrderBid,
		BidAmount: fmt.Sprintf("%f", sumOrderVolumeBid),
	}
}
