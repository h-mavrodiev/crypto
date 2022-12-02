package main

import (
	"crypto/configs"
	"crypto/internal/arbitrage"
	"crypto/internal/client"
	"crypto/internal/gate"
	"crypto/internal/server"
	"crypto/internal/stex"
	"fmt"
	"os"
)

func main() {

	conf_path := os.Getenv("CRYPTO_CONFIG_PATH")

	err := configs.LoadConfig("config", "yml", conf_path)
	if err != nil {
		fmt.Println(err)
	}

	go client.StartPlaftormsClient(&gate.PriceInfo, &gate.BalanceInfo, &stex.PriceInfo, &stex.BalanceInfo)
	go arbitrage.ExecuteArbitrage(&gate.PriceInfo, &stex.PriceInfo)
	r := server.Server()
	r.Run(":8080")
}
