package currency

import (
	"testing"
)

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}

type mockConverter struct {
	lastAmount float64
	lastFrom, lastTo string
	calls int
}

const mockResult float64 = 42.0

func (mc *mockConverter) Convert (amount float64, from, to string) float64 {
	mc.lastAmount = amount
	mc.lastFrom = from
	mc.lastTo = to
	mc.calls += 1

	return 42.0
}

func TestPriceIn(t *testing.T) {
	type test struct {
		amount float64
		from, to string
	}

	tests := map[string]test{
		"processes zero amount": {
			amount: 0,
			from: "dimension1",
			to: "dimension2",
		},
		"processes empty dimensions": {
			amount: 11.3,
			from: "",
			to: "",
		},
		"processes negative amount": {
			amount: -100,
			from: "dimension1",
			to: "dimension2",
		},
	}


	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mc := &mockConverter{}
			result := PriceIn(tt.amount, tt.from, tt.to, mc)

			assertEqual(t, result, mockResult)

			assertEqual(t, mc.lastAmount, tt.amount)
			assertEqual(t, mc.lastFrom, tt.from)
			assertEqual(t, mc.lastTo, tt.to)
			assertEqual(t, mc.calls, 1)
		})
	}
}