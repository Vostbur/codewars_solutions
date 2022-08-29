package kata

import "testing"

type testCase struct {
	arr      []float32
	expected float32
}

var testCases = []testCase{
	{[]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}, float32(2)},
	{[]float32{0, 0, 0.55, 0, 0}, float32(0.55)},
}

func TestFindUniq(t *testing.T) {
	for _, test := range testCases {
		if output := FindUniq(test.arr); output != test.expected {
			t.Errorf(`Expected
			<float32>: %.2f
		to equal
			<float32>: %.2f`, output, test.expected)
		}
	}
}
