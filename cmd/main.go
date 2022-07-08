package main

import (
	"crypto/configs"
	"encoding/json"
	"fmt"
	"log"
	"os"

	caller "crypto/internal/calls"
)

// ReadInput reads input from user
func ReadInput() (configs.Config, error) {
	var cn, ct, cp string
	var err error

	fmt.Println("path to config folder:  ")
	_, err = fmt.Scanln(&cp)
	if err != nil {
		fmt.Printf("No path was provided...\n")
		log.Fatal(err)
	} else {
		// Check if directory exists
		_, err := os.Stat(cp)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("config file name: ")
	_, err = fmt.Scanln(&cn)
	if err != nil {
		fmt.Printf("No name was provided...\n")
		log.Fatal(err)
	}

	fmt.Println("config file type: ")
	_, err = fmt.Scanln(&ct)
	if err != nil {
		fmt.Printf("No type was provided...\n")
		log.Fatal(err)
	}

	conf, err := configs.LoadConfig(cn, ct, cp)
	if err != nil {
		return configs.Config{}, err
	}

	return conf, nil
}

func main() {

	// uncomment to have the user point to a config
	// conf, err := ReadInput()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	conf, err := configs.LoadConfig("config", "yml", "/home/icetwo/configs")
	if err != nil {
		fmt.Println(err)
	}

	c := caller.Clients{}
	c.InitClients(conf)
	// c.MakeCalls()
	gateOrders := make(chan interface{}, 10)
	defer close(gateOrders)
	stexOrders := make(chan interface{}, 10)
	defer close(stexOrders)

	for i := 0; i < 100; i++ {
		c.GetOrderBooksConcurrently(gateOrders, stexOrders)

		select {
		case gateOrder := <-gateOrders:
			_, err := json.MarshalIndent(gateOrder, "", "")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string("GateOrder"))
			fmt.Println("=============================================")
		case stexOrder := <-stexOrders:
			_, err := json.MarshalIndent(stexOrder, "", "")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string("StexOrder"))
			fmt.Println("=============================================")
		}
	}
}
