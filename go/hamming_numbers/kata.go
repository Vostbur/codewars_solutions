package kata

func Hammer(n int) uint {
	ham := make([]uint, n)
	ham[0] = 1
	var dig2, dig3, dig5, nxt2, nxt3, nxt5 uint = 2, 3, 5, 2, 3, 5
	i, j, k := 0, 0, 0
	for m := 1; m < n; m++ {
		ham[m] = min(nxt2, min(nxt3, nxt5))
		if ham[m] == nxt2 {
			i++
			nxt2 = dig2 * ham[i]
		}
		if ham[m] == nxt3 {
			j++
			nxt3 = dig3 * ham[j]
		}
		if ham[m] == nxt5 {
			k++
			nxt5 = dig5 * ham[k]
		}
	}
	return ham[n-1]
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
