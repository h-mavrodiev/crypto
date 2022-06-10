package main

import (
	"crypto/configs"
	"fmt"

	GateClient "crypto/internal/gate"
	StexClient "crypto/internal/stex"
)

func main() {

	conf, err := configs.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conf.Gate.Host)
	}

	gateClinet := GateClient.NewClient(conf.Gate.Host, conf.Gate.Prefix, conf.Gate.APIKey, conf.Gate.GateEndpoints)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gateClinet)
	}

	stexClinet := StexClient.NewClient(conf.Stex.Host, conf.Stex.Prefix, conf.Stex.APIKey, conf.Stex.StexEndpoints)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stexClinet)
	}
}
