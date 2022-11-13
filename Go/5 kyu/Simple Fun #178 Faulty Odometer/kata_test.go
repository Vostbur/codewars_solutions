package kata

import "testing"

type testCase struct {
	n, expected int
}

var testCases = []testCase{
	{13, 12},
	{15, 13},
	{55, 40},
	{2005, 1462},
	{1500, 1053},
	{999999, 531440},
	{165826622, 69517865},
}

func TestFaultyOdometer(t *testing.T) {
	for _, test := range testCases {
		if output := FaultyOdometer(test.n); output != test.expected {
			t.Errorf(`Expected
			<int>: %d
		to equal
			<int>: %d`, output, test.expected)
		}
	}
}
