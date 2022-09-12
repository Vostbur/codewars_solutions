package kata

import (
	"regexp"
	"strings"
)

var MORSE_CODE = map[string]string{
	"-.-.--":    "!",
	".-..-.":    `"`,
	"...-..-":   "$",
	".-...":     "&",
	".----.":    "'",
	"-.--.":     "(",
	"-.--.-":    ")",
	".-.-.":     "+",
	"--..--":    ",",
	"-....-":    "-",
	".-.-.-":    ".",
	"-..-.":     "/",
	"-----":     "0",
	".----":     "1",
	"..---":     "2",
	"...--":     "3",
	"....-":     "4",
	".....":     "5",
	"-....":     "6",
	"--...":     "7",
	"---..":     "8",
	"----.":     "9",
	"---...":    ":",
	"-.-.-.":    ";",
	"-...-":     "=",
	"..--..":    "?",
	".--.-.":    "@",
	".-":        "A",
	"-...":      "B",
	"-.-.":      "C",
	"-..":       "D",
	".":         "E",
	"..-.":      "F",
	"--.":       "G",
	"....":      "H",
	"..":        "I",
	".---":      "J",
	"-.-":       "K",
	".-..":      "L",
	"--":        "M",
	"-.":        "N",
	"---":       "O",
	".--.":      "P",
	"--.-":      "Q",
	".-.":       "R",
	"...":       "S",
	"-":         "T",
	"..-":       "U",
	"...-":      "V",
	".--":       "W",
	"-..-":      "X",
	"-.--":      "Y",
	"--..":      "Z",
	"..--.-":    "_",
	"...---...": "SOS",
}

func TimeUnit(s [][]byte) (u int) {
	for i, v := range s {
		if i == 0 || u > len(v) {
			u = len(v)
		}
	}
	return
}

func DecodeBits(bits string) (r string) {
	bits = strings.Trim(bits, "0")
	re := regexp.MustCompile(`1+|0+`)
	unit := TimeUnit(re.FindAll([]byte(bits), -1))
	for _, word := range strings.Split(bits, strings.Repeat("0000000", unit)) {
		for _, letter := range strings.Split(word, strings.Repeat("000", unit)) {
			for _, char := range strings.Split(letter, strings.Repeat("0", unit)) {
				switch char {
				case strings.Repeat("111", unit):
					r += "-"
				case strings.Repeat("1", unit):
					r += "."
				}
			}
			r += " "
		}
		r += "   "
	}
	return r[:len(r)-1]
}

func DecodeMorse(morseCode string) (r string) {
	for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
		for _, letter := range strings.Split(word, " ") {
			r += MORSE_CODE[letter]
		}
		r += " "
	}
	return r[:len(r)-1]
}
