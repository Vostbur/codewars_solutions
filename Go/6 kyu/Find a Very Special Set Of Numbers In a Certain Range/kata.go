package kata

import (
	"strconv"
	"strings"
)

func FindUs(n1, n2, k int, primeFactors, digits []int) []int {
	var r []int

	var isDivisible = func(v int) bool {
		for _, c := range primeFactors {
			if v%c != 0 {
				return false
			}
		}
		return true
	}

	var haveDigits = func(v int) bool {
		s := strconv.Itoa(v)
		for _, d := range digits {
			if !strings.Contains(s, strconv.Itoa(d)) {
				return false
			}
		}
		return true
	}

	for v := n1; v <= k*n2+n1; v++ {
		if isDivisible(v) && haveDigits(v) {
			r = append(r, v)
		}
	}
	return r
}
