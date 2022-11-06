package kata

import "testing"

type testCase struct {
	strarr   []string
	k        int
	expected string
}

var testCases = []testCase{
	{[]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta"},
	{[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1,
		"oocccffuucccjjjkkkjyyyeehh"},
	{[]string{}, 3, ""},
	{[]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2,
		"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck"},
	{[]string{"wlwsasphmxx", "owiaxujylentrklctozmymu", "wpgozvxxiu"}, 2, "wlwsasphmxxowiaxujylentrklctozmymu"},
}

func TestPersistence(t *testing.T) {
	for _, test := range testCases {
		if output := LongestConsec(test.strarr, test.k); output != test.expected {
			t.Errorf(`Expected
			<int>: %s
		to equal
			<int>: %s`, output, test.expected)
		}
	}
}
