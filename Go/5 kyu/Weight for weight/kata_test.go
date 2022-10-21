package kata

import "testing"

type testCase struct {
	strng, expected string
}

var testCases = []testCase{
	{"103 123 4444 99 2000", "2000 103 123 4444 99"},
	{"2000 10003 1234000 44444444 9999 11 11 22 123", "11 11 2000 10003 22 123 1234000 44444444 9999"},
	{"", ""},
}

func TestOrderWeight(t *testing.T) {
	for _, test := range testCases {
		if output := OrderWeight(test.strng); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}
