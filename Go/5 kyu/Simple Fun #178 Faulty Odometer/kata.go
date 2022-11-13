package kata

import (
	"fmt"
	"math"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func FaultyOdometer(n int) (r int) {
	digs := "012356789"
	for i, v := range Reverse(fmt.Sprintf("%d", n)) {
		r += strings.Index(digs, string(v)) * PowInt(9, i)
	}
	return
}
