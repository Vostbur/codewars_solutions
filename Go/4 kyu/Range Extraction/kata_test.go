package kata

import "testing"

type testCase struct {
	list     []int
	expected string
}

var testCases = []testCase{
	{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, "-6,-3-1,3-5,7-11,14,15,17-20"},
}

func TestSolution(t *testing.T) {
	for _, test := range testCases {
		if output := Solution(test.list); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}
