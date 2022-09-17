package kata

import (
	"reflect"
	"testing"
)

type testCase struct {
	size     int
	expected [][]int
}

var testCases = []testCase{
	{5, [][]int{
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1},
		{1, 1, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	},
	},
	{6, [][]int{
		{1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 0, 1},
		{1, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1},
	},
	},
	{7, [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	},
	},
	{8, [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	},
	},
}

func TestSpiralize(t *testing.T) {
	for _, test := range testCases {
		if output := Spiralize(test.size); !reflect.DeepEqual(test.expected, output) {
			t.Errorf(`Expected
			<[][]int>: %v
		to equal
			<[][]int: %v`, output, test.expected)
		}
	}
}
