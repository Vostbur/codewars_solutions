package kata

import (
	"strconv"
	"strings"
)

func Solution(list []int) string {
	var seq []string
	for i := 0; ; {
		ii := i + 1
		for ii < len(list) && list[ii] == list[ii-1]+1 {
			ii++
		}
		s := strconv.Itoa(list[i])
		if ii == i+2 {
			s += "," + strconv.Itoa(list[ii-1])
		} else if ii > i+2 {
			s += "-" + strconv.Itoa(list[ii-1])
		}
		seq = append(seq, s)
		if ii == len(list) {
			break
		}
		i = ii
	}
	return strings.Join(seq, ",")
}
