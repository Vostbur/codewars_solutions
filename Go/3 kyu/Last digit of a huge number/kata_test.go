package kata

import "testing"

type testCase struct {
	as       []int
	expected int
}

var testCases = []testCase{
	{[]int{1, 2}, 1},
	{[]int{3, 4, 5}, 1},
	{[]int{4, 3, 6}, 4},
	{[]int{7, 6, 21}, 1},
	{[]int{12, 30, 21}, 6},
	{[]int{2, 0, 1}, 1},
	{[]int{2, 2, 2, 0}, 4},
	{[]int{937640, 767456, 981242}, 0},
	{[]int{123232, 694022, 140249}, 6},
	{[]int{499942, 898102, 846073}, 6},
}

func TestLastDigit(t *testing.T) {
	for _, test := range testCases {
		if output := LastDigit(test.as); output != test.expected {
			t.Errorf(`Expected
			<int>: %d
		to equal
			<int>: %d`, output, test.expected)
		}
	}
}
