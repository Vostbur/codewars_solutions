package kata

import (
	"sort"
	"strconv"
	"strings"
)

func OrderWeight(strng string) string {
	list := strings.Split(strng, " ")

	sum := func(s string) (r int) {
		sl := strings.Split(s, "")
		for _, v := range sl {
			tmp, _ := strconv.Atoi(v)
			r += tmp
		}
		return
	}

	sort.Slice(list, func(i, j int) bool {
		si, sj := sum(list[i]), sum(list[j])
		if si == sj {
			return list[i] < list[j]
		} else {
			return si < sj
		}
	})

	return strings.Join(list, " ")
}
