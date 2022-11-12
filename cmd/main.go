package main

import (
	"crypto/configs"
	// "crypto/internal/arbitrage"
	"crypto/internal/gate"
	// "crypto/internal/stex"
	"fmt"
	"os"
)

// var (
// 	gatePriceInfo         gate.GateInfo
// 	stexPriceInfo         stex.StexInfo
// 	ArbitrageResponseList []arbitrage.ArbitrageInfo
// )

func main() {

	conf_path := os.Getenv("CRYPTO_CONFIG_PATH")

	err := configs.LoadConfig("config", "yml", conf_path)
	if err != nil {
		fmt.Println(err)
	}

	// go arbitrage.ExecuteArbitrage(&gatePriceInfo, &stexPriceInfo, &ArbitrageResponseList)
	// go client.StartPlaftormsClient(&gatePriceInfo, &stexPriceInfo)
	// r := server.Server(&gatePriceInfo, &stexPriceInfo, &ArbitrageResponseList)
	// r.Run(":8080")
	gate.GateWSClient()
}
