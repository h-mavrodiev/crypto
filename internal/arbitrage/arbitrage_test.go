package arbitrage

import (
	"math"
	"testing"
)

func TestCalculateArbitrage(t *testing.T) {

	t.Log("test calculate Arbitrage")
	arb := CalculateArbitrage(20, 10)
	expected_arb := -100.0
	if arb != expected_arb {
		t.Errorf("Arbitrage Percentage was incorrect, got: %f, want: %f.", arb, expected_arb)
	}

	arb2 := CalculateArbitrage(1500, 1560)
	expected_arb2 := math.Round(3.846154)
	if math.Round(arb2) != expected_arb2 {
		t.Errorf("Arbitrage Percentage was incorrect, got: %f, want: %f.", arb2, expected_arb2)
	}
}
