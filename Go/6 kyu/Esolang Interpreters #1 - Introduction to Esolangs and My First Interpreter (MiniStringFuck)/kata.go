package kata

func Interpreter(code string) string {
	var (
		mem    byte
		output []byte
	)

	for _, c := range code {
		switch c {
		case '+':
			mem++
		case '.':
			output = append(output, mem)
		}
	}

	return string(output)
}
