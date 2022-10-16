package kata

import "testing"

type testCase struct {
	n, expected int
}

var testCases = []testCase{
	{39, 3},
	{4, 0},
	{25, 2},
	{999, 4},
}

func TestPersistence(t *testing.T) {
	for _, test := range testCases {
		if output := Persistence(test.n); output != test.expected {
			t.Errorf(`Expected
			<int>: %d
		to equal
			<int>: %d`, output, test.expected)
		}
	}
}
