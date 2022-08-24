package kata

import "testing"

type testCase struct {
	h, bounce, window float64
	expected          int
}

var testCases = []testCase{
	{3, 0.66, 1.5, 3},
	{40, 0.4, 10, 3},
	{10, 0.6, 10, -1},
	{40, 1, 10, -1},
	{5, -1, 1.5, -1},
	{2, 0.5, 1, 1},
	{4, 0.25, 1, 1},
	{30, 0.66, 1.5, 15},
}

func TestBouncingBall(t *testing.T) {
	for _, test := range testCases {
		if output := BouncingBall(test.h, test.bounce, test.window); output != test.expected {
			t.Errorf(`Expected
			<int>: %d
		to equal
			<int>: %d`, output, test.expected)
		}
	}
}
