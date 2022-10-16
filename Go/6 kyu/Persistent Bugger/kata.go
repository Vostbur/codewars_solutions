package kata

import (
	"strconv"
	"strings"
)

func mul(list []string) int {
	m := 1
	for _, l := range list {
		k, _ := strconv.Atoi(l)
		m *= k
	}
	return m
}

func Persistence(n int) (r int) {
	for ; len(strconv.Itoa(n)) > 1; r++ {
		l := strings.Split(strconv.Itoa(n), "")
		n = mul(l)
	}
	return
}
