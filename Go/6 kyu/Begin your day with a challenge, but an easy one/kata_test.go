package kata

import (
	"reflect"
	"testing"
)

type testCase struct {
	n        int
	expected [2]string
}

var testCases = []testCase{
	{0, [2]string{"0", "0"}},
	{1, [2]string{"1", "1"}},
	{2, [2]string{"2", "11"}},
	{3, [2]string{"3", "111"}},
	{4, [2]string{"4", "1111"}},
	{5, [2]string{"5", "11111"}},
	{6, [2]string{"6", "111111"}},
	{7, [2]string{"7", "1111111"}},
	{8, [2]string{"8", "11111111"}},
	{9, [2]string{"9", "111111111"}},
	{10, [2]string{"91", "1111111111"}},
	{11, [2]string{"92", "11111111111"}},
	{12, [2]string{"93", "111111111111"}},
	{13, [2]string{"94", "1111111111111"}},
	{14, [2]string{"95", "11111111111111"}},
	{15, [2]string{"96", "111111111111111"}},
	{16, [2]string{"97", "1111111111111111"}},
	{17, [2]string{"98", "11111111111111111"}},
	{18, [2]string{"99", "111111111111111111"}},
	{19, [2]string{"991", "1111111111111111111"}},
	{20, [2]string{"992", "11111111111111111111"}},
	{21, [2]string{"993", "111111111111111111111"}},
	{22, [2]string{"994", "1111111111111111111111"}},
	{23, [2]string{"995", "11111111111111111111111"}},
	{24, [2]string{"996", "111111111111111111111111"}},
	{25, [2]string{"997", "1111111111111111111111111"}},
	{26, [2]string{"998", "11111111111111111111111111"}},
	{27, [2]string{"999", "111111111111111111111111111"}},
	{28, [2]string{"9991", "1111111111111111111111111111"}},
}

func TestOneTwoThreeSol(t *testing.T) {
	for _, test := range testCases {
		if output := OneTwoThreeSol(test.n); !reflect.DeepEqual(output, test.expected) {
			t.Errorf(`Expected
			<[2]string>: %v
		to equal
			<[2]string>: %v`, output, test.expected)
		}
	}
}
