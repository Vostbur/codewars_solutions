package kata

import (
	"reflect"
	"sort"
	"testing"
)

type testCase struct {
	s        string
	expected []string
}

var abcd = []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
	"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
	"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
	"dabc", "dacb", "dbac", "dbca", "dcab", "dcba"}

var testCases = []testCase{
	{"a", []string{"a"}},
	{"ab", []string{"ab", "ba"}},
	{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
	{"abcd", abcd},
	{"bcda", abcd},
	{"dcba", abcd},
	{"aa", []string{"aa"}},
	{"aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	{"aaaab", []string{"aaaab", "aaaba", "aabaa", "abaaa", "baaaa"}},
	{"aaaaab", []string{"aaaaab", "aaaaba", "aaabaa", "aabaaa", "abaaaa", "baaaaa"}},
}

func TestPermutations(t *testing.T) {
	for _, test := range testCases {
		output := Permutations(test.s)
		sort.Strings(output)
		sort.Strings(test.expected)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf(`Expected
			<[]string>: %v
		to equal
			<[]string>: %v`, output, test.expected)
		}
	}
}
