package caesar

import (
	"unicode"
)

func Enc(n int, plain string) string {
	ciper := ""

	for _, c := range plain {
		tmp := int(c)
		if unicode.IsLetter(c) {
			tmp = int(c) + n

			if (int(c) < int('Z') && tmp > int('Z')) ||
				(int(c) > int('a') && tmp > int('z')) {
				tmp -= 26
			}
		}
		ciper += string(tmp)
	}

	return ciper
}

func Dec(n int, ciper string) string {
	plain := ""

	for _, c := range ciper {
		tmp := int(c)
		if unicode.IsLetter(c) {
			tmp = int(c) - n

			if (int(c) < int('Z') && tmp < int('A')) ||
				(int(c) > int('a') && tmp < int('a')) {
				tmp += 26
			}
		}
		plain += string(tmp)
	}

	return plain
}
