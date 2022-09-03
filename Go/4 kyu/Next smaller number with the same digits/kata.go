package kata

import "strconv"

func NextSmaller(n int) (r int) {
	s := []rune(strconv.Itoa(n))

	i := len(s) - 1
	for i > 0 && s[i-1] <= s[i] {
		i--
	}
	if i == 0 {
		return -1
	}

	j := len(s) - 1
	for s[j] >= s[i-1] {
		j--
	}
	s[i-1], s[j] = s[j], s[i-1]

	if s[0] == rune('0') {
		return -1
	}

	for k, m := i, len(s)-1; k < m; k, m = k+1, m-1 {
		s[k], s[m] = s[m], s[k]
	}
	r, _ = strconv.Atoi(string(s))
	return
}
