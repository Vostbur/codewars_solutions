package kata

import "strconv"

var digits = map[byte][]int{
	'0': {0},
	'1': {1},
	'2': {6, 2, 4, 8},
	'3': {1, 3, 9, 7},
	'4': {6, 4},
	'5': {5},
	'6': {6},
	'7': {1, 7, 9, 3},
	'8': {6, 8, 4, 2},
	'9': {1, 9},
}

func LastDigit(n1, n2 string) int {
	if (n1 == "0" && n2 == "0") || n2 == "0" {
		return 1
	}

	n1Last, n2Last := n1[len(n1)-1], 0
	if len(n2) > 1 {
		n2Last, _ = strconv.Atoi(string(n2[len(n2)-2]) + string(n2[len(n2)-1]))
	} else {
		n2Last, _ = strconv.Atoi(string(n2[len(n2)-1]))
	}

	return digits[n1Last][n2Last%len(digits[n1Last])]
}
