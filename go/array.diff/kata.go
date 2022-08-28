package kata

func ArrayDiff(a, b []int) (res []int) {
	m := map[int]bool{}
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if !m[v] {
			res = append(res, v)
		}
	}
	return
}
