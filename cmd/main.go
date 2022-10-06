package main

import (
	"crypto/configs"
	"fmt"

	caller "crypto/internal/calls"
	"crypto/internal/gate"
	"crypto/internal/stex"
	"crypto/internal/tui"
)

func main() {

	// uncomment to have the user point to a config
	// conf, err := configs.LoadConfigFromInput()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	var (
		gatePriceInfo gate.GateInfo
		stexPriceInfo stex.StexInfo
	)

	conf, err := configs.LoadConfig("config", "yml", "/home/icetwo/configs")
	if err != nil {
		fmt.Println(err)
	}

	c := caller.Clients{}
	c.InitClients(conf)

	// Those channels were never closed... find out if you should actually close them ...
	gateOrders := make(chan gate.OrderBookDetails)
	stexOrders := make(chan stex.OrderBookDetails)

	gateAggOrders := make(chan *gate.GateInfo)
	stexAggOrders := make(chan *stex.StexInfo)

	go c.CallStexGetOrderBookDetails(stexOrders, "1.5s", 407)
	go c.CallGateGetOrderBookDetails(gateOrders, "1s", "ETH_USDT")
	go tui.RunTUI(gateAggOrders, stexAggOrders, gatePriceInfo, stexPriceInfo)

	for {
		select {
		case gateOrder := <-gateOrders:
			gateInfo := gateOrder.FindPriceAndVolume(1)
			gateAggOrders <- gateInfo
		case stexOrder := <-stexOrders:
			stexInfo := stexOrder.FindPriceAndVolume(2.2)
			stexAggOrders <- stexInfo
		}
	}

}
