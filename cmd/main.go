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
	gateOrders := make(chan interface{})
	// defer close(gateOrders)

	stexOrders := make(chan interface{})
	// defer close(stexOrders)
	gate := 0
	stex := 0

	// for {
	for i := 0; i < 300; i++ {
		// c.GetOrderBooksConcurrently(gateOrders, stexOrders)
		c.CallStexGetOrderBookDetails(stexOrders)
		c.CallGateGetOrderBookDetails(gateOrders)
		select {
		case gateOrder := <-gateOrders:
			_, err := json.MarshalIndent(gateOrder, "", "")
			if err != nil {
				fmt.Println("not nil")
			}
			gate++
			// fmt.Println(string(gateRes))
		case stexOrder := <-stexOrders:
			_, err := json.MarshalIndent(stexOrder, "", "")
			if err != nil {
				fmt.Println("not nil")
			}
			stex++
			// fmt.Println(string(stexRes))
		}
	}
	fmt.Printf("Gate requests: %d\n", gate)
	fmt.Printf("Stex requests: %d\n", stex)
}
