package kata

func Spiralize(size int) [][]int {
	r := make([][]int, size)
	for row := 0; row < size; row++ {
		r[row] = make([]int, size)
	}

	var (
		x, y, incX, incY int
		dir              int  = 1
		line             bool = true
	)
	for t := 0; t < size; {
		r[x][y] = 1
		if x == 0+incX && !line && dir == -1 {
			line = true
			dir = -dir
			t++
		} else if x == size-incX-1 && !line && dir == 1 {
			line = true
			incX += 2
			dir = -dir
			t++
		} else if y == size-incY-1 && line && dir == 1 {
			line = false
			t++
		} else if x != 0 && y == 0+incY && line && dir == -1 {
			line = false
			incY += 2
			t++
		}
		if line {
			y += dir
		} else {
			x += dir
		}
	}
	return r
}
