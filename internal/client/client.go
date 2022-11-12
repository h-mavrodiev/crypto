package client

import (
	"crypto/configs"
	"crypto/internal/gate"
	"crypto/internal/stex"
)

var (
	c Clients = loadClients()
)

// Loads configuration yml for clients and returns Clients struct
func loadClients() Clients {

	c := Clients{}
	c.InitClients(configs.Conf)

	return c
}

func StartPlaftormsClient(gatePriceInfo *gate.GateInfo, stexPriceInfo *stex.StexInfo) {
	// Those channels were never closed... find out if you should actually close them ...
	gateOrders := make(chan gate.OrderBookDetails)
	stexOrders := make(chan stex.OrderBookDetails)

	go c.CallStexGetOrderBookDetails(stexOrders, "1.5s", 407)
	go c.CallGateGetOrderBookDetails(gateOrders, "1s", "ETH_USDT")

	for {
		select {
		case gateOrder := <-gateOrders:
			gatePriceInfo.CalcPriceAndVolume(gateOrder, 1)
			// fmt.Println("GATE")
			// print, _ := json.Marshal(gatePriceInfo)
			// fmt.Println(string(print))
		case stexOrder := <-stexOrders:
			stexPriceInfo.CalcPriceAndVolume(stexOrder, 1)
			// fmt.Println("STEX")
			// print, _ := json.Marshal(stexPriceInfo)
			// fmt.Println(string(print))
		}
	}
}
