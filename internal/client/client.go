package client

import (
	"crypto/configs"
	"crypto/internal/gate"
	"crypto/internal/stex"
	"log"
	"time"
)

var (
	c    Clients
	errs chan error = make(chan error)
)

func (c *Clients) InitHTTPClients() {
	var err error

	c.GateClient, err = gate.NewClient(configs.Conf.Gate.Host,
		configs.Conf.Gate.Prefix,
		configs.Conf.Gate.Endpoints,
		configs.Conf.Gate.CommonHeaders,
		configs.Conf.Gate.Pair,
		configs.Conf.Gate.APIKey,
		configs.Conf.Gate.APISecret)
	if err != nil {
		log.Printf("Could not initiate Gate Client: %s", err)
	}

	c.StexClient = stex.NewClient(configs.Conf.Stex.Host,
		configs.Conf.Stex.APIKey,
		configs.Conf.Stex.Endpoints,
		configs.Conf.Stex.CommonHeaders,
		configs.Conf.Stex.Pair)

}

func StartPlaftormsClient(gatePriceInfo *gate.SafePrices, gateBalance *gate.SafeBalance, stexPriceInfo *stex.SafePrices, stexBalance *stex.SafeBalance) {
	c.InitHTTPClients()
	go c.GateClient.RunGate(gatePriceInfo, gateBalance, errs)
	go c.StexClient.RunStex(stexPriceInfo, stexBalance, errs)

	// just to give chance for the clients to spin up
	time.Sleep(2 * time.Second)

	for {
		err := <-errs
		log.Println(err)
	}
}

// t, _ := time.ParseDuration("4s")
// time.Sleep(t)
// fmt.Println("GATE")
// print, _ := json.Marshal(gatePriceInfo)
// fmt.Println(string(print))
