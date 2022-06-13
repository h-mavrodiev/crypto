package main

import (
	"crypto/configs"
	"fmt"
	"log"
	"os"

	GateClient "crypto/internal/gate"
	StexClient "crypto/internal/stex"
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

	// conf, err := ReadInput()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	conf, err := configs.LoadConfig("config", "yml", "/home/icetwo/configs")
	if err != nil {
		fmt.Println(err)
	}

	gateClinet, err := GateClient.NewClient(conf.Gate.Host, conf.Gate.Prefix, conf.Gate.GateEndpoints,
		conf.Gate.CommonHeaders, conf.Gate.APIKey, conf.Gate.APISecret)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gateClinet)
	}

	r, err := gateClinet.GetListChains("currency=GT")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("###")
		fmt.Print(r)
		fmt.Println("###")
	}

	// fmt.Println("###")
	// tss, errr := gateClinet.GenSign(gateClinet.Endpoints.Wallet, "GET", "", p)
	// if errr != nil {
	// 	fmt.Println(errr)
	// }
	// fmt.Println(tss)
	// fmt.Println("###")

	stexClinet := StexClient.NewClient(conf.Stex.Host, conf.Stex.APIKey, conf.Stex.StexEndpoints)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stexClinet)
	}
}
