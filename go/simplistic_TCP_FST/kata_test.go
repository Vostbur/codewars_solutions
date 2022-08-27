package kata

import "testing"

type testCase struct {
	events   []string
	expected string
}

var testCases = []testCase{
	{[]string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN"}, "CLOSE_WAIT"},
	{[]string{"APP_PASSIVE_OPEN", "RCV_SYN", "RCV_ACK"}, "ESTABLISHED"},
	{[]string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN", "APP_CLOSE"}, "LAST_ACK"},
	{[]string{"APP_ACTIVE_OPEN"}, "SYN_SENT"},
	{[]string{"APP_PASSIVE_OPN", "RCV_SYN", "RCV_ACK", "APP_CLOSE", "APP_SEND"}, "ERROR"},
}

func TestTraverseTCPStates(t *testing.T) {
	for _, test := range testCases {
		if output := TraverseTCPStates(test.events); output != test.expected {
			t.Errorf(`Expected
			<string>: %s
		to equal
			<string>: %s`, output, test.expected)
		}
	}
}
