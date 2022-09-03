package kata

import "testing"

type testCase struct {
	n, expected int
}

var testCases = []testCase{
	{21, 12},
	{531, 513},
	{2071, 2017},
	{123456798, 123456789},
	{1234567908, 1234567890},
	{9, -1},
	{123456789, -1},
	{111, -1},
	{135, -1},
	{1027, -1},
}

func TestNextSmaller(t *testing.T) {
	for _, test := range testCases {
		if output := NextSmaller(test.n); output != test.expected {
			t.Errorf(`Expected
			<int>: %d
		to equal
			<int>: %d`, output, test.expected)
		}
	}
}
