package kata

import "strings"

func LongestConsec(strarr []string, k int) (r string) {
	n := len(strarr)
	if n == 0 || k > n || k <= 0 {
		return
	}

	var res []string
	for i := 0; i <= n-k; i++ {
		var tmp []string
		for j := i; j < i+k; j++ {
			tmp = append(tmp, strarr[j])
		}
		res = append(res, strings.Join(tmp, ""))
	}

	l, r := len(res[0]), res[0]
	for _, v := range res {
		if l < len(v) {
			l = len(v)
			r = v
		}
	}
	return
}

// === BEST SOLUTION ===
// package kata

// import "strings"

// func LongestConsec(strarr []string, k int) string {
//   var buffer string
//   var largest string

//   for i := 0; i <= len(strarr)-k; i++ {
//     buffer = strings.Join(strarr[i : i+k][:], "")
//     if len(buffer) > len(largest) {
//       largest = buffer
//     }
//   }
//   return largest
// }