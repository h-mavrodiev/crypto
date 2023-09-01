package gate

import "sync"

type safeOrderBook struct {
	mu        sync.Mutex
	orderBook orderBook
}

// updateOrderBookFromHTTP takes order book update from http requests and updates OrderBookDetails object
func (o *safeOrderBook) updateOrderBookFromHTTP(ob *orderBook) {
	o.mu.Lock()
	o.orderBook = orderBook{
		Asks: ob.Asks,
		Bids: ob.Bids,
	}
	o.mu.Unlock()

}

// updateOrderBookFromHTTP takes order book update from ws notification and updates OrderBookDetails object
func (o *safeOrderBook) updateOrderBookFromWS(wsOb *orderBookUpdateNotification) {
	o.mu.Lock()
	o.orderBook = orderBook{
		Asks: wsOb.Asks,
		Bids: wsOb.Bids,
	}
	o.mu.Unlock()
}
