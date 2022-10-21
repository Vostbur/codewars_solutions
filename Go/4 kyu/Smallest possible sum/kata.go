package kata

func Solution(ar []int) int {
	list := make([]int, len(ar))
	copy(list, ar)

	for count(list[0], list) != len(list) {
		small := min(list)
		for i, v := range list {
			p := v - (small * (v / small))
			if v > small && p != 0 {
				list[i] = p
			} else {
				list[i] = small
			}
		}
	}
	return list[0] * len(ar)
}

func count(el int, ar []int) (n int) {
	for _, v := range ar {
		if v == el {
			n++
		}
	}
	return
}

func min(ar []int) (m int) {
	m = ar[0]
	for _, v := range ar {
		if m > v {
			m = v
		}
	}
	return
}
