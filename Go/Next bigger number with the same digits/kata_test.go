package kata

import "testing"

type testCase struct {
	n, expected int
}

var testCases = []testCase{
	{12, 21},
	{513, 531},
	{2017, 2071},
	{414, 441},
	{144, 414},
	{9, -1},
	{111, -1},
	{531, -1},
	{1234567890, 1234567908},
}

func TestNextBigger(t *testing.T) {
	for _, test := range testCases {
		if output := NextBigger(test.n); output != test.expected {
			t.Errorf(`Expected
			<int>: %d
		to equal
			<int>: %d`, output, test.expected)
		}
	}
}
