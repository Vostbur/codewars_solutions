package kata

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	sum := 0
	nStr := strconv.Itoa(n)
	for _, v := range nStr {
		digit, _ := strconv.Atoi(string(v))
		sum += int(math.Pow(float64(digit), float64(p)))
		p++
	}
	if sum%n == 0 {
		return sum / n
	}
	return -1
}
