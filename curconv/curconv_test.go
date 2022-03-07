package curconv

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestRoundToDecStr(t *testing.T) {
	tests := []struct {
		decimals int
		amount   decimal.Decimal
		expected string
	}{
		{2, decimal.NewFromFloat(1.0), "1"},
		{2, decimal.NewFromFloat(1.99), "1.99"},
		{2, decimal.NewFromFloat(1.004), "1"},
		{2, decimal.NewFromFloat(1.005), "1.01"},
		{3, decimal.NewFromFloat(1.0), "1"},
		{3, decimal.NewFromFloat(1.999), "1.999"},
		{3, decimal.NewFromFloat(1.0004), "1"},
		{3, decimal.NewFromFloat(1.0005), "1.001"},
		{4, decimal.NewFromFloat(1.0), "1"},
		{4, decimal.NewFromFloat(1.9999), "1.9999"},
		{4, decimal.NewFromFloat(1.00004), "1"},
		{4, decimal.NewFromFloat(1.00005), "1.0001"},
		{0, decimal.NewFromFloat(1.0), "1"},
		{0, decimal.NewFromFloat(1.4), "1"},
		{0, decimal.NewFromFloat(1.9), "2"},
		{0, decimal.NewFromFloat(1.5), "2"},
	}
	for _, test := range tests {
		actual := RoundToDecStr(test.amount, test.decimals)
		if actual != test.expected {
			t.Errorf("RoundToDecStr(%d, %s) = %s; expected %s", test.decimals, test.amount.String(), actual, test.expected)
		}
	}
}

func TestRoundToMax(t *testing.T) {
	tests := []struct {
		decimals int
		amount   decimal.Decimal
		expected string
	}{
		{2, decimal.NewFromFloat(1.0), "1"},
		{2, decimal.NewFromFloat(1.99), "1.99"},
		{2, decimal.NewFromFloat(1.004), "1"},
		{2, decimal.NewFromFloat(1.005), "1.01"},
		{3, decimal.NewFromFloat(1.0), "1"},
		{3, decimal.NewFromFloat(1.999), "1.999"},
		{3, decimal.NewFromFloat(1.0004), "1"},
		{3, decimal.NewFromFloat(1.0005), "1.001"},
		{4, decimal.NewFromFloat(1.0), "1"},
		{4, decimal.NewFromFloat(1.9999), "1.9999"},
		{4, decimal.NewFromFloat(1.00004), "1"},
		{4, decimal.NewFromFloat(1.00005), "1.0001"},
		{0, decimal.NewFromFloat(1.0), "1"},
		{0, decimal.NewFromFloat(1.4), "1"},
		{0, decimal.NewFromFloat(1.9), "2"},
		{0, decimal.NewFromFloat(1.5), "2"},
	}
	for _, test := range tests {
		expectedDec, _ := decimal.NewFromString(test.expected)
		curr := testRoundToMaxGetAny(t, test.decimals)
		actual := RoundToMaxStr(test.amount, curr)
		if actual != test.expected {
			t.Errorf("RoundToMaxStr(%s, %s) = %s; expected %s", curr, test.amount.String(), actual, test.expected)
		}

		actualDec := RoundToMax(test.amount, curr)
		if !actualDec.Equal(expectedDec) {
			t.Errorf("RoundToMax(%s, %s) = %s; expected %s", curr, test.amount.String(), actual, test.expected)
		}
	}
}

// testRoundToMaxGetAny gest the first found currency with n decimals.
func testRoundToMaxGetAny(t *testing.T, n int) string {
	if n == 2 {
		return "MXN"
	}
	for curr, dec := range decimals {
		if dec == n {
			return curr
		}
	}
	t.Fatalf("No currency found with %d decimals", n)
	return ""
}
