package kata

import (
	"reflect"
	"testing"
)

type testCase struct {
	n        int
	expected [][]int
}

var testCases = []testCase{
	{0, [][]int{}},
	{1, [][]int{{1}}},
	{2, [][]int{{1}, {1, 1}}},
	{3, [][]int{{1}, {1, 1}, {1, 1, 1}}},
}

func TestPyramid(t *testing.T) {
	for _, test := range testCases {
		if output := Pyramid(test.n); !reflect.DeepEqual(output, test.expected) {
			t.Errorf(`Expected
			<[][]int: %v
		to equal
			<[][]int: %v`, output, test.expected)
		}
	}
}
