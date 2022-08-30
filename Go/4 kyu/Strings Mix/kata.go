package kata

import (
	"sort"
	"strings"
)

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func Mix(s1, s2 string) string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	r := make(map[int][]string)
	for _, v := range letters {
		v1 := strings.Count(s1, string(v))
		v2 := strings.Count(s2, string(v))
		idx := ""
		m := max(v1, v2)
		if m > 1 {
			switch {
			case v1 > v2:
				idx = "1"
			case v1 < v2:
				idx = "2"
			case v1 == v2:
				idx = "="
			}
			r[m] = append(r[m], idx+":"+strings.Repeat(string(v), m))
		}
	}

	r_map := make(map[int]string)
	for i, v := range r {
		sort.Strings(v)
		r_map[-i] = strings.Join(v, "/")
	}

	keys := make([]int, 0, len(r_map))
	for k := range r_map {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var result []string
	for _, k := range keys {
		result = append(result, r_map[k])
	}

	return strings.Join(result, "/")
}
