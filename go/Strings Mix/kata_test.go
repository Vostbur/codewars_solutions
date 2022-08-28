package kata

import "testing"

type testCase struct {
	arg1, arg2, expected string
}

var testCases = []testCase{
	{"Are they here", "yes, they are here", "2:eeeee/2:yy/=:hh/=:rr"},
	{"uuuuuu", "uuuuuu", "=:uuuuuu"},
	{"looping is fun but dangerous", "less dangerous than coding",
		"1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg"},
}

func TestMix(t *testing.T) {
	for _, test := range testCases {
		if output := Mix(test.arg1, test.arg2); output != test.expected {
			t.Errorf(`Expected
				<string>: %s
			to equal
				<string>: %s`, output, test.expected)
		}
	}
}
