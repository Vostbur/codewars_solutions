package kata

func FindUniq(arr []float32) float32 {
	u := arr[0]
	for i := 1; i < len(arr); i++ {
		if u != arr[i] {
			if i > 1 || (i == 1 && u == arr[2]) {
				u = arr[i]
			}
			break
		}
	}
	return u
}
