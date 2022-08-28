package kata

import (
	"regexp"
	"strconv"
	"testing"
)

func TestFunc(t *testing.T) {
	type test struct {
		value  string
		result bool
	}
	simpleTests := []test{{" 0", false}, {"abc", false}, {"000", true}, {"110", true}, {"111", false}, {strconv.FormatInt(12345678, 2), true}}
	for _, test := range simpleTests {
		if matched, _ := regexp.Match(MultipleOf3Regex, []byte(test.value)); matched != test.result {
			t.Errorf("should test that the solution returns the correct value for " + test.value)
		}
	}
}
