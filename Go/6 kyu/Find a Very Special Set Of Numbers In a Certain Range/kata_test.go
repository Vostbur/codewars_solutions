package kata

import (
	"reflect"
	"testing"
)

type testCase struct {
	n1, n2, k                      int
	primeFactors, digits, expected []int
}

var testCases = []testCase{
	{30, 90, 4, []int{2, 3}, []int{6, 2}, []int{126, 162, 216, 246, 264, 276}},
	{30, 400, 18, []int{2, 3, 7}, []int{6, 2, 5}, []int{2562, 2856, 5628, 6258, 6552}},
}

func TestFindUs(t *testing.T) {
	for _, test := range testCases {
		if output := FindUs(test.n1, test.n2, test.k, test.primeFactors, test.digits); !reflect.DeepEqual(test.expected, output) {
			t.Errorf(`Expected
			<[]int>: %v
		to equal
			<[]int>: %v`, output, test.expected)
		}
	}
}
