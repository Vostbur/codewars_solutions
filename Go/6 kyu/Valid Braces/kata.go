package kata

func ValidBraces(str string) bool {
	var tracer []rune
	for _, c := range str {
		if c == '(' || c == '{' || c == '[' {
			tracer = append(tracer, c)
		} else {
			if len(tracer) == 0 {
				return false
			}
			lastVal := tracer[len(tracer)-1]
			if (c == ')' && lastVal == '(') || (c == '}' && lastVal == '{') || (c == ']' && lastVal == '[') {
				tracer = tracer[:len(tracer)-1]
			} else {
				break
			}
		}
	}
	return len(tracer) == 0
}
