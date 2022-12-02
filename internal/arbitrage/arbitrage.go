package arbitrage

import (
	"crypto/internal/gate"
	"crypto/internal/stex"
)

var (
	gateToStex            float64
	stexToGate            float64
	ArbitrageResponseList []ArbitrageInfo = []ArbitrageInfo{
		{
			Platforms: "Gate to Stex",
			Arbitrage: gateToStex,
		},
		{
			Platforms: "Stex to Gate",
			Arbitrage: stexToGate,
		}}
)

func CalculateArbitrage(sellsPrice float64, buysPrice float64) float64 {
	// how much the buy price is bigger compared to the sell price?
	// (a-b)*100 / a  to find percentage diff
	// This way when there is no arbitrage, the percentage value is always negative
	return (buysPrice - sellsPrice) * 100 /
		buysPrice

}

func ExecuteArbitrage(gatePriceInfo *gate.SafePrices,
	stexPriceInfo *stex.SafePrices) {

	for {
		gateToStex = CalculateArbitrage(gatePriceInfo.Prices.Sells, stexPriceInfo.Prices.Buys)
		stexToGate = CalculateArbitrage(stexPriceInfo.Prices.Sells, gatePriceInfo.Prices.Buys)

		ArbitrageResponseList = []ArbitrageInfo{
			{
				Platforms: "Gate to Stex",
				Arbitrage: gateToStex,
			},
			{
				Platforms: "Stex to Gate",
				Arbitrage: stexToGate,
			}}

		// print, _ := json.Marshal(*ArbitrageResponseList)
		// fmt.Println(string(print))

		switch {
		case gateToStex > 1:
			//log.Printf("$$$$$$$$$$$$$$$$$$$ GATE ---> STEX ARBITRAGE %f $$$$$$$$$$$$$$$$$$$\n", gateToStex)
		case stexToGate > 1:
			//log.Printf("$$$$$$$$$$$$$$$$$$$ STEX ---> GATE ARBITRAGE %f $$$$$$$$$$$$$$$$$$$\n", stexToGate)
		}
	}
}
