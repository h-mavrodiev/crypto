package client

import (
	"crypto/configs"
	"crypto/internal/gate"
	"crypto/internal/stex"
	"fmt"
)

var (
	c Clients = loadClients()
	// gatePriceInfo gate.GateInfo
	// stexPriceInfo stex.StexInfo
)

// Loads configuration yml for clients and returns Clients struct
func loadClients() Clients {
	conf, err := configs.LoadConfig("config", "yml", "/Users/I576893/nonwork/configs")
	if err != nil {
		fmt.Println(err)
	}
	c := Clients{}
	c.InitClients(conf)

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
			gatePriceInfo.CalcPriceAndVolume(gateOrder, 20, 20)
			// fmt.Println("GATE")
			// print, _ := json.Marshal(gatePriceInfo)
			// fmt.Println(string(print))
		case stexOrder := <-stexOrders:
			stexPriceInfo.CalcPriceAndVolume(stexOrder, 20, 20)
			// fmt.Println("STEX")
			// print, _ := json.Marshal(stexPriceInfo)
			// fmt.Println(string(print))
		}
	}
}
