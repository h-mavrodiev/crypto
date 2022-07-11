package calls

import (
	"fmt"
	"time"

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

func (c *Clients) CallGateListChains(ch chan<- interface{}) {
	// ETH-USDT code is 407
	go func() {
		for {
			err := c.GateClient.GetListChains("currency", "USDT", ch)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

func (c *Clients) CallGateWithdrawalRecords(ch chan<- interface{}) {
	// ETH-USDT code is 407
	go func() {
		for {
			err := c.GateClient.GetWithdrawalRecords("", "", ch)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

func (c *Clients) CallGateTotalBalance(ch chan<- interface{}) {
	// ETH-USDT code is 407
	go func() {
		for {
			err := c.GateClient.GetTotalBalance("", "", ch)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

func (c *Clients) CallStexProfileInfo(ch chan<- interface{}) {
	// ETH-USDT code is 407
	go func() {
		for {
			err := c.StexClient.GetProfileInfo(ch)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

func (c *Clients) CallStexCurrencyPairFees(pair int, ch chan<- interface{}) {
	// ETH-USDT code is 407
	go func() {
		for {
			err := c.StexClient.GetCurrencyPairFees(pair, ch)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

// }

func (c *Clients) CallStexGetCurrencyPairDetails(ch chan<- interface{}) {
	// ETH-USDT code is 407
	go func() {
		for {
			err := c.StexClient.GetCurrencyPairDetails(407, ch)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}

func (c *Clients) CallStexGetOrderBookDetails(ch chan<- interface{}) {
	// ETH-USDT code is 407
	go func() {
		for {
			// Sleep is needed as Stex API is much faster to response ...
			// 1/4 seconds seems to work best
			time.Sleep(time.Second / 4)
			err := c.StexClient.GetOrderBookDetails(407, ch)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
}

func (c *Clients) CallGateGetOrderBookDetails(ch chan<- interface{}) {
	go func() {
		for {
			// time.Sleep(time.Second / 5)
			err := c.GateClient.GetOrderBookDetails("ETH_USDT", ch)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
}
