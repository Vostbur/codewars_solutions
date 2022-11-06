package kata

import "testing"

type testCase struct {
	n, expected uint64
}

var testCases = []testCase{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{5, 7},
	{10, 42},
}

func TestExpSum(t *testing.T) {
	for _, test := range testCases {
		if output := ExpSum(test.n); output != test.expected {
			t.Errorf(`Expected
			<uint64>: %d
		to equal
			<uint64>: %d`, output, test.expected)
		}
	}
}
