package kata

import "testing"

type testCase struct {
	n, p     int
	expected int
}

var testCases = []testCase{
	{89, 1, 1},
	{92, 1, -1},
	{695, 2, 2},
	{46288, 3, 51},
}

func TestDigPow(t *testing.T) {
	for _, test := range testCases {
		if output := DigPow(test.n, test.p); output != test.expected {
			t.Errorf(`Expected
			<int>: %d
		to equal
			<int>: %d`, output, test.expected)
		}
	}
}
