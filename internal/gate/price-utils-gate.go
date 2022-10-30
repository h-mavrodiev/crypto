package gate

import (
	"fmt"
	"strconv"
	"sync"
)

type GateInfo struct {
	AskPrice  string
	AskAmount string
	BidPrice  string
	BidAmount string
}

func (o *OrderBookDetails) FindPriceAndVolume(volumeByPrice float64) *GateInfo {

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
			for _, order := range o.Asks {
				s, err := strconv.ParseFloat(order[1], 64)
				if err != nil {
					priceOrderAsk = "Error Parsing Price"
					break
				}
				if sumOrdersVolumesAsk < volumeByPrice {
					sumOrdersVolumesAsk = sumOrdersVolumesAsk + s
				}
				priceOrderAsk = order[0]
			}

		}(&wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for _, order := range o.Bids {
				s, err := strconv.ParseFloat(order[1], 64)
				if err != nil {
					priceOrderBid = "Error Parsing Price"
					break
				}
				if sumOrdersVolumesAsk < volumeByPrice {
					sumOrderVolumeBid = sumOrderVolumeBid + s
				}
				priceOrderBid = order[0]
			}
		}(&wg)
	}
	wg.Wait()

	return &GateInfo{
		AskPrice:  priceOrderAsk,
		AskAmount: fmt.Sprintf("%f", sumOrdersVolumesAsk),
		BidPrice:  priceOrderBid,
		BidAmount: fmt.Sprintf("%f", sumOrderVolumeBid),
	}

}
