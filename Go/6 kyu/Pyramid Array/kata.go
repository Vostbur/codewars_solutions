package kata

func Pyramid(n int) [][]int {
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			tmp[j] = 1
		}
		r[i] = tmp
	}
	return r
}
