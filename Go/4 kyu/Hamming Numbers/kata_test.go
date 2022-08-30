package kata

import "testing"

type testCase struct {
	n        int
	expected uint
}

var testCases = []testCase{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 4},
	{5, 5},
	{6, 6},
	{7, 8},
	{8, 9},
	{9, 10},
	{10, 12},
	{11, 15},
	{12, 16},
	{13, 18},
	{14, 20},
	{15, 24},
	{16, 25},
	{17, 27},
	{18, 30},
	{19, 32},
}

func TestHammer(t *testing.T) {
	for _, test := range testCases {
		if output := Hammer(test.n); output != test.expected {
			t.Errorf(`Expected
			<uint>: %d
		to equal
			<uint>: %d`, output, test.expected)
		}
	}
}
