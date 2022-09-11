package kata

import "testing"

type testCase struct {
	morseCode, expected string
}

var testCases = []testCase{
	{".... . -.--   .--- ..- -.. .", "HEY JUDE"},
}

func TestDecodeMorse(t *testing.T) {
	for _, test := range testCases {
		if output := DecodeMorse(test.morseCode); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}
