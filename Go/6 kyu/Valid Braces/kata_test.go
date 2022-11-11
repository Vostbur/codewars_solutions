package kata

import "testing"

type testCase struct {
	str      string
	expected bool
}

var testCases = []testCase{
	{"(){}[]", true},
	{"([{}])", true},
	{"(}", false},
	{"[(])", false},
	{"[({)](]", false},
}

func TestValidBraces(t *testing.T) {
	for _, test := range testCases {
		if output := ValidBraces(test.str); output != test.expected {
			t.Errorf(`Expected
			<bool>: %v
		to equal
			<bool>: %v`, output, test.expected)
		}
	}
}
