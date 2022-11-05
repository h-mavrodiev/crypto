package main

import (
	"crypto/internal/client"
	"crypto/internal/gate"
	"crypto/internal/server"
	"crypto/internal/stex"
	"log"
)

var (
	gatePriceInfo gate.GateInfo
	stexPriceInfo stex.StexInfo
)

func printArbitrage(gatePriceInfo *gate.GateInfo, stexPriceInfo *stex.StexInfo) {

	percentGate := (gatePriceInfo.TheySellForWeighted - gatePriceInfo.ICanSellFromStexForWeighted) * 100 / gatePriceInfo.ICanSellFromStexForWeighted
	percantStex := (stexPriceInfo.TheySellForWeighted - stexPriceInfo.ICanSellFromGateForWeighted) * 100 / stexPriceInfo.ICanSellFromGateForWeighted

	for {
		switch {
		case percentGate > 1:
			log.Println("$$$$$$$$$$$$$$$$$$$ STEX ---> GATE ARBITRAGE $$$$$$$$$$$$$$$$$$$")
		case percantStex > 1:
			log.Println("$$$$$$$$$$$$$$$$$$$ GATE ---> STEX ARBITRAGE $$$$$$$$$$$$$$$$$$$")
		}
	}
}

func main() {

	go printArbitrage(&gatePriceInfo, &stexPriceInfo)
	go client.StartPlaftormsClient(&gatePriceInfo, &stexPriceInfo)
	r := server.Server(&gatePriceInfo, &stexPriceInfo)
	r.Run(":8080")

}
