package client

import (
	"fmt"
	"time"

	"crypto/configs"
	gate "crypto/internal/gate"
	stex "crypto/internal/stex"
)

func (c *Clients) InitClients(conf configs.Config) {

	c.GateClient = gate.NewClient(conf.Gate.Host,
		conf.Gate.Prefix,
		conf.Gate.Endpoints,
		conf.Gate.CommonHeaders,
		conf.Gate.APIKey,
		conf.Gate.APISecret)

	c.StexClient = stex.NewClient(conf.Stex.Host,
		conf.Stex.APIKey,
		conf.Stex.Endpoints,
		conf.Stex.CommonHeaders)

}

func (c *Clients) CallGateListChains(ch chan<- interface{}, delayTime string) {
	// ETH-USDT code is 407
	for {
		h, _ := time.ParseDuration(delayTime)
		time.Sleep(h)
		err := c.GateClient.GetListChains("currency", "USDT", ch)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *Clients) CallGateWithdrawalRecords(ch chan<- interface{}, delayTime string) {
	// ETH-USDT code is 407
	for {
		h, _ := time.ParseDuration(delayTime)
		time.Sleep(h)
		err := c.GateClient.GetWithdrawalRecords("", "", ch)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *Clients) CallGateTotalBalance(ch chan<- interface{}, delayTime string) {
	// ETH-USDT code is 407
	go func() {
		for {
			h, _ := time.ParseDuration(delayTime)
			time.Sleep(h)
			err := c.GateClient.GetTotalBalance("", "", ch)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

func (c *Clients) CallStexProfileInfo(ch chan<- interface{}, delayTime string) {
	// ETH-USDT code is 407
	for {
		h, _ := time.ParseDuration(delayTime)
		time.Sleep(h)
		err := c.StexClient.GetProfileInfo(ch)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *Clients) CallStexCurrencyPairFees(pair int, ch chan<- interface{}, delayTime string) {
	// ETH-USDT code is 407
	for {
		h, _ := time.ParseDuration(delayTime)
		time.Sleep(h)
		err := c.StexClient.GetCurrencyPairFees(pair, ch)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// }

func (c *Clients) CallStexGetCurrencyPairDetails(ch chan<- interface{}, delayTime string) {
	// ETH-USDT code is 407
	for {
		h, _ := time.ParseDuration(delayTime)
		time.Sleep(h)
		err := c.StexClient.GetCurrencyPairDetails(407, ch)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *Clients) CallStexGetOrderBookDetails(ch chan<- stex.OrderBookDetails, delayTime string, currencyPairId int) {
	for {
		// ETH-USDT code is 407
		h, _ := time.ParseDuration(delayTime)
		time.Sleep(h)
		err := c.StexClient.GetOrderBookDetails(currencyPairId, ch)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (c *Clients) CallGateGetOrderBookDetails(ch chan<- gate.OrderBookDetails, delayTime string, currencyPair string) {
	for {
		h, _ := time.ParseDuration(delayTime)
		time.Sleep(h)
		err := c.GateClient.GetOrderBookDetails(currencyPair, ch)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
