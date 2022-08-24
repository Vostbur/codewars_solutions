package kata

import (
	"reflect"
	"testing"
)

type testCase struct {
	a, b     uint64
	expected []uint64
}

var testCases = []testCase{
	{1, 10, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{1, 100, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 89}},
	{10, 89, []uint64{89}},
	{10, 100, []uint64{89}},
	{90, 100, nil},
	{89, 135, []uint64{89, 135}},
	{2, 8, []uint64{2, 3, 4, 5, 6, 7, 8}},
	{517, 519, []uint64{518}},
	{196, 811, []uint64{518, 598}},
	{12157692622039623185, 12157692622039625655, []uint64{12157692622039623539}},
	{12157692622039623539, 12157692622039623539, []uint64{12157692622039623539}},
}

func TestSumDigPow(t *testing.T) {
	for _, test := range testCases {
		if output := SumDigPow(test.a, test.b); !reflect.DeepEqual(output, test.expected) {
			t.Errorf(`Expected
			<[]uint64>: %v
		to equal
			<[]uint64>: %v`, output, test.expected)
		}
	}
}
