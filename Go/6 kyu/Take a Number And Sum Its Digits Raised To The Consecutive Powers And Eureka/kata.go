package kata

func SumDigPow(a, b uint64) (r []uint64) {
	digits := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if a < 9 && b < 9 {
		r = append(r, digits[a-1:b]...)
		return
	}

	if a+b < 18 {
		r = append(r, digits[a-1:b-1]...)
		return
	}

	if a < 10 {
		r = append(r, digits[a-1:]...)
	}

	if a <= 89 && b >= 89 {
		r = append(r, 89)
	}

	if a <= 135 && b >= 135 {
		r = append(r, 135)
	}

	if a < 136 {
		a = 136
	}

	for a <= b {
		s := sliceFromDigs(a)
		if checkSum(s, a) {
			r = append(r, a)
		}
		a++
	}
	return
}

func sliceFromDigs(i uint64) (r []uint64) {
	for i > 0 {
		r = append(r, i%10)
		i /= 10
	}
	return
}

func checkSum(s []uint64, i uint64) bool {
	var t uint64
	var p uint64 = 1
	for i := len(s) - 1; i >= 0; i-- {
		t += uintPow(s[i], p)
		p++
	}
	return t == i
}

func uintPow(n, m uint64) (r uint64) {
	if m == 0 {
		return 1
	}
	r = n
	for i := uint64(2); i <= m; i++ {
		r *= n
	}
	return
}
