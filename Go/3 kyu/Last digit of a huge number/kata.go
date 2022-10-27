package kata

import "math"

func LastDigit(as []int) int {
	var a int = 1
	for i := len(as) - 1; i >= 0; i-- {
		exp := a%4 + 4
		if a < 4 {
			exp = a
		}
		base := as[i]%20 + 20
		if as[i] < 20 {
			base = as[i]
		}
		a = int(math.Pow(float64(base), float64(exp)))
	}
	return a % 10
}
