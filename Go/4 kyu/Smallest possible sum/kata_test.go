package kata

import "testing"

type testCase struct {
	ar       []int
	expected int
}

var testCases = []testCase{
	{[]int{9}, 9},
	{[]int{6, 9, 21}, 9},
	{[]int{1, 21, 55}, 3},
}

func TestSolution(t *testing.T) {
	for _, test := range testCases {
		if output := Solution(test.ar); output != test.expected {
			t.Errorf(`Expected
			<int>: %d
		to equal
			<int>: %d`, output, test.expected)
		}
	}
}
