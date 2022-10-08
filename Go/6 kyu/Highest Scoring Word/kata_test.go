package kata

import "testing"

type testCase struct {
	s, expected string
}

var testCases = [...]testCase{
	// {"man i need a taxi up to ubud", "taxi"},
	// {"what time are we climbing up the volcano", "volcano"},
	// {"take me to semynak", "semynak"},
	// {"aa b", "aa"},
	// {"b aa", "b"},
	// {"bb d", "bb"},
	{"d bb", "d"},
	// {"aaa b", "aaa"},
}

func TestHigh(t *testing.T) {
	for _, test := range testCases {
		if output := High(test.s); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}
