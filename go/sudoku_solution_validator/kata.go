package kata

import "sort"

func compare(matrix []int) bool {
	tmp := make([]int, 9)
	copy(tmp, matrix)
	sort.Ints(tmp)
	for i, v := range tmp {
		if v-1 != i {
			return false
		}
	}
	return true
}

func ValidateSolution(m [][]int) bool {
	for i := 0; i < 9; i++ {
		if !compare(m[i]) {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		col := make([]int, 9)
		for j := 0; j < 9; j++ {
			col[j] = m[j][i]
		}
		if !compare(col) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		var square []int
		for j := 0; j < 9; j += 3 {
			square = append(square, m[i][j:j+3]...)
			square = append(square, m[i+1][j:j+3]...)
			square = append(square, m[i+2][j:j+3]...)
			if !compare(square) {
				return false
			}
		}
	}
	return true
}
