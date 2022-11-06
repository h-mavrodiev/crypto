package arbitrage

import (
	"crypto/internal/gate"
	"crypto/internal/stex"
	"log"
)

type ArbitrageInfo struct {
	Platforms string
	Arbitrage float64
}

func prepareArbitrageInfo() {
	for {

		// print, _ := json.Marshal(ArbitrageResponseList)
		// fmt.Println(string(print))
	}
}

func CalculateArbitrage(ICanSellFromForValue float64, TheySellForValue float64) float64 {

	// (a-b)*100 / a  to find percentage diff
	// This way when there is no arbitrage, the percentage value is always negative
	return (ICanSellFromForValue - TheySellForValue) * 100 /
		ICanSellFromForValue

}

func ExecuteArbitrage(gatePriceInfo *gate.GateInfo,
	stexPriceInfo *stex.StexInfo,
	ArbitrageResponseList *[]ArbitrageInfo,
	GateToStex *float64,
	StexToGate *float64) {

	for {

		*GateToStex = CalculateArbitrage(gatePriceInfo.ICanSellFromStexForWeighted, gatePriceInfo.TheySellForWeighted)
		*StexToGate = CalculateArbitrage(stexPriceInfo.ICanSellFromGateForWeighted, stexPriceInfo.TheySellForWeighted)
		*ArbitrageResponseList = []ArbitrageInfo{{Platforms: "Gate to Stex", Arbitrage: *GateToStex},
			{Platforms: "Stex to Gate", Arbitrage: *StexToGate}}

		// print, _ := json.Marshal(*ArbitrageResponseList)
		// fmt.Println(string(print))

		switch {
		case *GateToStex > 1:
			log.Printf("$$$$$$$$$$$$$$$$$$$ GATE ---> STEX ARBITRAGE %f $$$$$$$$$$$$$$$$$$$\n", *GateToStex)
		case *StexToGate > 1:
			log.Printf("$$$$$$$$$$$$$$$$$$$ STEX ---> GATE ARBITRAGE %f $$$$$$$$$$$$$$$$$$$\n", *StexToGate)
		}
	}
}
