package kata

import (
	"reflect"
	"testing"
)

type testCase struct {
	a, b, expected []int
}

var emptyArr []int

var testCases = []testCase{
	{[]int{1, 2}, []int{1}, []int{2}},
	{[]int{1, 2, 2}, []int{1}, []int{2, 2}},
	{[]int{1, 2, 2}, []int{2}, []int{1}},
	{[]int{1, 2, 2}, emptyArr, []int{1, 2, 2}},
	{emptyArr, []int{1, 2}, emptyArr},
	{[]int{1, 2, 3}, []int{1, 2}, []int{3}},
}

func TestArrayDiff(t *testing.T) {
	for _, test := range testCases {
		if output := ArrayDiff(test.a, test.b); !reflect.DeepEqual(output, test.expected) {
			t.Errorf(`Expected
			<[]int>: %v
		to equal
			<[]int>: %v`, output, test.expected)
		}
	}
}
