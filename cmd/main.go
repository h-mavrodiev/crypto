package main

import (
	"crypto/internal/client"
	"crypto/internal/gate"
	"crypto/internal/server"
	"crypto/internal/stex"
)

var (
	gatePriceInfo gate.GateInfo
	stexPriceInfo stex.StexInfo
)

func main() {

	go client.StartPlaftormsClient(&gatePriceInfo, &stexPriceInfo)
	r := server.Server(&gatePriceInfo, &stexPriceInfo)
	r.Run(":8080")

}
