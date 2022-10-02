package kata

func Interpreter(code, tape string) string {
	var mem = []byte(tape)
	for codeIdx, tapeIdx := 0, 0; codeIdx < len(code); codeIdx++ {
		if tapeIdx < 0 || tapeIdx >= len(tape) {
			break
		}
		switch code[codeIdx] {
		case '>':
			tapeIdx++
		case '<':
			tapeIdx--
		case '*':
			if mem[tapeIdx] == '0' {
				mem[tapeIdx] = '1'
			} else {
				mem[tapeIdx] = '0'
			}
		case '[':
			if mem[tapeIdx] == '0' {
				for bracks := 1; bracks > 0; {
					codeIdx++
					switch code[codeIdx] {
					case '[':
						bracks++
					case ']':
						bracks--
					}
				}
			}
		case ']':
			if mem[tapeIdx] != '0' {
				for bracks := -1; bracks < 0; {
					codeIdx--
					switch code[codeIdx] {
					case '[':
						bracks++
					case ']':
						bracks--
					}
				}
			}
		}
	}
	return string(mem)
}
