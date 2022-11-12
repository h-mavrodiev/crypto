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

func CalculateArbitrage(sellsPrice float64, buysPrice float64) float64 {
	// how much the buy price is bigger compared to the sell price?
	// (a-b)*100 / a  to find percentage diff
	// This way when there is no arbitrage, the percentage value is always negative
	return (buysPrice - sellsPrice) * 100 /
		buysPrice

}

func ExecuteArbitrage(gatePriceInfo *gate.GateInfo,
	stexPriceInfo *stex.StexInfo,
	ArbitrageResponseList *[]ArbitrageInfo) {

	for {
		GateToStex := CalculateArbitrage(gatePriceInfo.Sells, stexPriceInfo.Buys)
		StexToGate := CalculateArbitrage(stexPriceInfo.Sells, gatePriceInfo.Buys)
		*ArbitrageResponseList = []ArbitrageInfo{{Platforms: "Gate to Stex", Arbitrage: GateToStex},
			{Platforms: "Stex to Gate", Arbitrage: StexToGate}}

		// print, _ := json.Marshal(*ArbitrageResponseList)
		// fmt.Println(string(print))

		switch {
		case GateToStex > 1:
			log.Printf("$$$$$$$$$$$$$$$$$$$ GATE ---> STEX ARBITRAGE %f $$$$$$$$$$$$$$$$$$$\n", GateToStex)
		case StexToGate > 1:
			log.Printf("$$$$$$$$$$$$$$$$$$$ STEX ---> GATE ARBITRAGE %f $$$$$$$$$$$$$$$$$$$\n", StexToGate)
		}
	}
}
