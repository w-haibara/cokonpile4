package kuku

import (
	"strconv"
)

func pad(n int, plength int, pstr string) string {
	s := strconv.Itoa(n)
	for i := len(s); i < plength; i++ {
		s = pstr + s
	}
	return s
}

func Table(n int) string {
	l := len(strconv.Itoa(n * n))
	res := ""

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			res += pad(i*j, l, " ") + " "
		}
		res += "\n"
	}
	return res
}
