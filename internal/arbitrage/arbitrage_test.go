package arbitrage

import "testing"

func TestCalculateArbitrage(t *testing.T) {

	t.Log("test calculate Arbitrage")
	arb := CalculateArbitrage(20, 10)
	expected_arb := 50.0
	if arb != expected_arb {
		t.Errorf("Arbitrage Percentage was incorrect, got: %f, want: %f.", arb, expected_arb)
	}

	arb2 := CalculateArbitrage(10, 15)
	expected_arb2 := -50.0
	if arb2 != expected_arb2 {
		t.Errorf("Arbitrage Percentage was incorrect, got: %f, want: %f.", arb2, expected_arb2)
	}
}
