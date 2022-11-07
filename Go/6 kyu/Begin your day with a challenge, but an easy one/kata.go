package kata

import (
	"fmt"
	"strings"
)

func OneTwoThreeSol(n int) (r [2]string) {
	if n == 0 {
		return [2]string{"0", "0"}
	}
	remain := n % 9
	if remain == 0 {
		r[0] = fmt.Sprintf("%s", strings.Repeat("9", n/9))
	} else {
		r[0] = fmt.Sprintf("%s%d", strings.Repeat("9", n/9), remain)
	}
	r[1] = strings.Repeat("1", n)
	return
}
