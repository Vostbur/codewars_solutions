package kata

import (
	"math/big"
	"testing"
)

type TestCase struct {
	n        int64
	expected *big.Int
}

var TestCases = []TestCase{
	{0, big.NewInt(0)},
	{1, big.NewInt(1)},
	{2, big.NewInt(1)},
	{3, big.NewInt(2)},
	{5, big.NewInt(5)},
	{50, big.NewInt(12586269025)},
	{-67, big.NewInt(44945570212853)},
	{-77, big.NewInt(5527939700884757)},
	{-86, big.NewInt(-420196140727489673)},
}

func TestFibDinamic(t *testing.T) {
	for _, test := range TestCases {
		output := FibDinamic(test.n)
		if result := test.expected.Cmp(output); result != 0 {
			t.Errorf(`Expected
			<*big.Int>: %v
		to equal
			<*big.Int>: %v`, output, test.expected)
		}
	}
}

func TestFibMatrix(t *testing.T) {
	for _, test := range TestCases {
		output := FibMatrix(test.n)
		if result := test.expected.Cmp(output); result != 0 {
			t.Errorf(`Expected
			<*big.Int>: %v
		to equal
			<*big.Int>: %v`, output, test.expected)
		}
	}
}
