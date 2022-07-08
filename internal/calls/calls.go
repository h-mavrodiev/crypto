package calls

import (
	"fmt"
	"sync"

	"crypto/configs"

	GateClient "crypto/internal/gate"
	StexClient "crypto/internal/stex"
)

type Clients struct {
	GateClient *GateClient.GateClient
	StexClient *StexClient.StexClient
}

func (c *Clients) InitClients(conf configs.Config) {

	c.GateClient = GateClient.NewClient(conf.Gate.Host,
		conf.Gate.Prefix,
		conf.Gate.Endpoints,
		conf.Gate.CommonHeaders,
		conf.Gate.APIKey,
		conf.Gate.APISecret)

	c.StexClient = StexClient.NewClient(conf.Stex.Host,
		conf.Stex.APIKey,
		conf.Stex.Endpoints,
		conf.Stex.CommonHeaders)

}

// func (c *Clients) Calls() {

// 	r, err := c.GateClient.GetListChains("currency", "USDT")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		jsonSTR, err := json.MarshalIndent(r, "", "")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(string(jsonSTR))
// 	}

// 	res, err := c.GateClient.GetWithdrawalRecords("", "")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		jsonSTR, err := json.MarshalIndent(res, "", "")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(string(jsonSTR))
// 	}

// 	balance, err := c.GateClient.GetTotalBalance("", "")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		jsonSTR, err := json.MarshalIndent(balance, "", "")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(string(jsonSTR))
// 	}

// 	info, err := c.StexClient.GetProfileInfo()
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		jsonSTR, err := json.MarshalIndent(info, "", "")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(string(jsonSTR))
// 	}

// 	pairFees, err := c.StexClient.GetCurrencyPairFees(1)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		jsonSTR, err := json.MarshalIndent(pairFees, "", "")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(string(jsonSTR))
// 	}

// }

func (c *Clients) CallStexGetCurrencyPairDetails(ch chan<- interface{}) {
	// ETH-USDT code is 407
	stexPairDetails, err := c.StexClient.GetCurrencyPairDetails(407)
	if err != nil {
		fmt.Println(err)
	}
	ch <- stexPairDetails
}

func (c *Clients) GetOrderBooksConcurrently(
	gateChanel chan interface{},
	stexChanel chan interface{}) {

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		err := c.GateClient.GetOrderBookDetails("ETH_USDT", gateChanel, &wg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	go func() {
		err := c.StexClient.GetOrderBookDetails(407, stexChanel, &wg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	wg.Wait()
}
