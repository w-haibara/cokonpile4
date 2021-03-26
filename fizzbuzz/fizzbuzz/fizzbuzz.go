package fizzbuzz

import (
	"strconv"
)

type Config struct {
	N1   uint
	N2   uint
	Fizz string
	Buzz string
}

func Default() Config {
	return Config{
		N1:   3,
		N2:   5,
		Fizz: "Fizz",
		Buzz: "Buzz",
	}
}

func (c Config) Say(n uint) string {
	v1 := n%c.N1 == 0
	v2 := n%c.N2 == 0
	switch {
	case v1 && v2:
		return c.Fizz + c.Buzz
	case v1:
		return c.Fizz
	case v2:
		return c.Buzz
	}

	return strconv.Itoa(int(n))
}
