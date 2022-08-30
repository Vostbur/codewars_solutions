package kata

import "testing"

type testCase struct {
	numbers  [10]uint
	expected string
}

var testCases = []testCase{
	{[10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, "(123) 456-7890"},
	{[10]uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, "(111) 111-1111"},
}

func TestBouncingBall(t *testing.T) {
	for _, test := range testCases {
		if output := CreatePhoneNumber(test.numbers); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}
