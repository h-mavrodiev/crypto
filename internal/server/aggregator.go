package server

import (
	"crypto/internal/gate"
	"crypto/internal/stex"
)

type PricesAggregator struct {
	Platform    string
	Sells       *float64
	SellsVolume *float64
	Buys        *float64
	BuysVolume  *float64
}

var prcAgg []PricesAggregator

func init() {
	prcAgg = []PricesAggregator{
		{
			Platform:    "Gate",
			Sells:       &gate.PriceInfo.Prices.Sells,
			SellsVolume: &gate.PriceInfo.Prices.SellsVolume,
			Buys:        &gate.PriceInfo.Prices.Buys,
			BuysVolume:  &gate.PriceInfo.Prices.BuysVolume,
		},
		{
			Platform:    "Stex",
			Sells:       &stex.PriceInfo.Prices.Sells,
			SellsVolume: &stex.PriceInfo.Prices.SellsVolume,
			Buys:        &stex.PriceInfo.Prices.Buys,
			BuysVolume:  &stex.PriceInfo.Prices.BuysVolume,
		},
	}
}

type BalanceAggregator struct {
	Platform string
	BTC      *string
	ETH      *string
	USDT     *string
}

var blncAgg []BalanceAggregator

func init() {
	blncAgg = []BalanceAggregator{
		{
			Platform: "Gate",
			BTC:      &gate.BalanceInfo.Balance.BTC,
			ETH:      &gate.BalanceInfo.Balance.ETH,
			USDT:     &gate.BalanceInfo.Balance.USDT,
		},
		{
			Platform: "Stex",
			BTC:      &stex.BalanceInfo.Balance.BTC,
			ETH:      &stex.BalanceInfo.Balance.ETH,
			USDT:     &stex.BalanceInfo.Balance.USDT,
		},
	}
}
