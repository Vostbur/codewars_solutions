package kata

import "strings"

func High(s string) string {
	max, lenght, sum := 0, 0, func(v string) int {
		var s int
		for _, i := range v {
			s += int(i) - 0x60
		}
		return s
	}

	arr := strings.Split(s, " ")
	for idx, val := range arr {
		if s := sum(val); lenght < s {
			lenght, max = s, idx
		}
	}
	return arr[max]
}
