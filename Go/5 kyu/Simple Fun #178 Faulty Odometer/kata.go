package kata

import (
	"fmt"
	"math"
	"strings"
)

func FaultyOdometer(n int) (r int) {
	digs := "012356789"
	s := fmt.Sprintf("%d", n)
	l := len(s) - 1
	for i, v := range s {
		r += strings.Index(digs, string(v)) * int(math.Pow(9, float64(l-i)))
	}
	return
}
