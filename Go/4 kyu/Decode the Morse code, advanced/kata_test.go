package kata

import "testing"

type testCase struct {
	morseCode, expected string
}

var testCases = []testCase{
	{"1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011",
		"HEY JUDE"},
	{"01110", "E"},
}

func TestDecodeMorse(t *testing.T) {
	for _, test := range testCases {
		if output := DecodeMorse(DecodeBits(test.morseCode)); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}
