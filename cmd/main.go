package main

import (
	"crypto/internal/arbitrage"
	"crypto/internal/client"
	"crypto/internal/gate"
	"crypto/internal/server"
	"crypto/internal/stex"
)

var (
	gatePriceInfo         gate.GateInfo
	stexPriceInfo         stex.StexInfo
	ArbitrageResponseList []arbitrage.ArbitrageInfo
)

func main() {

	go arbitrage.ExecuteArbitrage(&gatePriceInfo, &stexPriceInfo, &ArbitrageResponseList)
	go client.StartPlaftormsClient(&gatePriceInfo, &stexPriceInfo)
	r := server.Server(&gatePriceInfo, &stexPriceInfo, &ArbitrageResponseList)
	r.Run(":8080")

}
