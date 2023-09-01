package client

import (
	"fmt"
	"time"

	"crypto/internal/gate"
)

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

func (c *Clients) CallGateTotalBalance(ch chan<- *gate.TotalBalance, delayTime string) {
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

func (c *Clients) CallStexCurrencyPairFees(ch chan<- interface{}, delayTime string) {
	// ETH-USDT code is 407
	for {
		h, _ := time.ParseDuration(delayTime)
		time.Sleep(h)
		err := c.StexClient.GetCurrencyPairFees(c.StexClient.Pair, ch)
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
		err := c.StexClient.GetCurrencyPairDetails(c.StexClient.Pair, ch)
		if err != nil {
			fmt.Println(err)
		}
	}
}
