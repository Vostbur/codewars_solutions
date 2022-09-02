package kata

// The final version basE91 codec is taken from https://github.com/Equim-chan/base91-go (BSD 3-Clause License)
// I wrote a similar code, but I understood the algorithm using this example

var (
	enctab = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~'")

	dectab = map[byte]byte{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7,
		'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15,
		'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
		'Y': 24, 'Z': 25, 'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31,
		'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39,
		'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47,
		'w': 48, 'x': 49, 'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55,
		'4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61, '!': 62, '#': 63,
		'$': 64, '%': 65, '&': 66, '(': 67, ')': 68, '*': 69, '+': 70, ',': 71,
		'.': 72, '/': 73, ':': 74, ';': 75, '<': 76, '=': 77, '>': 78, '?': 79,
		'@': 80, '[': 81, ']': 82, '^': 83, '_': 84, '`': 85, '{': 86, '|': 87,
		'}': 88, '~': 89, '\'': 90,
	}
)

func Encode(d []byte) (encoded []byte) {
	var b, n uint32 = 0, 0

	for i := 0; i < len(d); i++ {
		b |= uint32(d[i]) << n
		n += 8

		if n <= 13 {
			continue
		}

		v := b & 0x1fff
		if v > 88 {
			b >>= 13
			n -= 13
		} else {
			v = b & 0x3fff
			b >>= 14
			n -= 14
		}

		encoded = append(encoded, enctab[v%91], enctab[v/91|0])
	}

	if n != 0 {
		encoded = append(encoded, enctab[b%91])
		if n > 7 || b > 90 {
			encoded = append(encoded, enctab[b/91])
		}
	}

	return
}

func Decode(d []byte) (decoded []byte) {
	var b, n uint32 = 0, 0
	v := -1

	for i := 0; i < len(d); i++ {
		p := dectab[d[i]]
		if p > 90 {
			continue
		}
		if v < 0 {
			v = int(p)
			continue
		}
		v += int(p) * 91
		b |= uint32(v) << n
		if v&0x1fff > 88 {
			n += 13
		} else {
			n += 14
		}
		for {
			decoded = append(decoded, uint8(b))
			b >>= 8
			n -= 8
			if n <= 7 {
				break
			}
		}
		v = -1
	}

	if v > -1 {
		decoded = append(decoded, uint8(b|uint32(v)<<n))
	}

	return
}
