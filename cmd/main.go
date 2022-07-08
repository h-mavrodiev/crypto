package main

import (
	"crypto/configs"
	"encoding/json"
	"fmt"
	"time"

	caller "crypto/internal/calls"
)

func printExecutionTime(t time.Time) {
	fmt.Println("Execution time: ", time.Since(t))
}

func main() {

	// uncomment to have the user point to a config
	// conf, err := configs.LoadConfigFromInput()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	conf, err := configs.LoadConfig("config", "yml", "/home/icetwo/configs")
	if err != nil {
		fmt.Println(err)
	}

	startTime := time.Now()
	defer printExecutionTime(startTime)

	c := caller.Clients{}
	c.InitClients(conf)
	gateOrders := make(chan interface{}, 10)
	defer close(gateOrders)
	stexOrders := make(chan interface{}, 10)
	defer close(stexOrders)

	for i := 0; i < 15; i++ {
		c.GetOrderBooksConcurrently(gateOrders, stexOrders)
		select {
		case gateOrder := <-gateOrders:
			gateRes, _ := json.MarshalIndent(gateOrder, "", "")
			fmt.Println(string(gateRes))
		case stexOrder := <-stexOrders:
			stexRes, _ := json.MarshalIndent(stexOrder, "", "")
			fmt.Println(string(stexRes))
		}
	}
}
