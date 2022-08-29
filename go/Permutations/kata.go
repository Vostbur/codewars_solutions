package kata

func Permutations(s string) []string {
	var (
		res  []string
		perm func([]rune, int, int)
	)
	perm = func(r []rune, i1, i2 int) {
		if i1 == i2 {
			res = append(res, string(r))
		} else {
			for i := i1; i <= i2; i++ {
				r[i1], r[i] = r[i], r[i1]
				perm(r, i1+1, i2)
				r[i1], r[i] = r[i], r[i1]
			}
		}
	}
	perm([]rune(s), 0, len(s)-1)

	// remove duplicates
	keys := make(map[string]bool)
	l := []string{}
	for _, v := range res {
		if !keys[v] {
			keys[v] = true
			l = append(l, v)
		}
	}
	return l
}
