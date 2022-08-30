package kata

func ProductFib(prod uint64) [3]uint64 {
	i, j := uint64(0), uint64(1)
	res := map[bool]uint64{
		true:  1,
		false: 0,
	}
	for i*j < prod {
		i, j = j, i+j
	}
	return [3]uint64{i, j, res[prod == i*j]}
}
