package stex

import "sync"

type safeOrderBook struct {
	mu        sync.Mutex
	orderBook orderBook
}

// updateOrderBookFromHTTP takes order book update from http requests and updates OrderBookDetails object
func (o *safeOrderBook) updateOrderBookFromHTTP(ob *orderBook) {
	o.mu.Lock()
	o.orderBook = orderBook{
		Ask: ob.Ask,
		Bid: ob.Bid,
	}
	o.mu.Unlock()

}
