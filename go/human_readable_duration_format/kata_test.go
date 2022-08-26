package kata

import "testing"

type testCase struct {
	seconds  int64
	expected string
}

var testCases = []testCase{
	{1, "1 second"},
	{62, "1 minute and 2 seconds"},
	{120, "2 minutes"},
	{3600, "1 hour"},
	{3662, "1 hour, 1 minute and 2 seconds"},
	{132030240, "4 years, 68 days, 3 hours and 4 minutes"},
}

func TestFormatDuration(t *testing.T) {
	for _, test := range testCases {
		if output := FormatDuration(test.seconds); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}
