package kata

func BouncingBall(h, bounce, window float64) (r int) {
	r = -1
	if h <= 0 || !(bounce > 0 && bounce < 1) || window >= h {
		return
	}

	for ; h > window; h *= bounce {
		r += 2
	}
	return
}
