package main

import (
	"crypto/internal/arbitrage"
	"crypto/internal/client"
	"crypto/internal/gate"
	"crypto/internal/server"
	"crypto/internal/stex"
	"fmt"
)

var (
	GateToStex            float64
	StexToGate            float64
	gatePriceInfo         gate.GateInfo
	stexPriceInfo         stex.StexInfo
	ArbitrageResponseList []arbitrage.ArbitrageInfo
)

func main() {

	go arbitrage.ExecuteArbitrage(&gatePriceInfo, &stexPriceInfo, &ArbitrageResponseList, &GateToStex, &StexToGate)
	go client.StartPlaftormsClient(&gatePriceInfo, &stexPriceInfo)
	fmt.Println(ArbitrageResponseList)
	fmt.Println("HI")

	r := server.Server(&gatePriceInfo, &stexPriceInfo, &ArbitrageResponseList)
	r.Run(":8080")

}
