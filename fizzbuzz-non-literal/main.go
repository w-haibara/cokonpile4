package main

const (
	i0 = iota
	i1
	i2
	i3
	i4
	i5
	i6
	i7
	i8
	i9
	i10
)

var (
	f        = rune(i10*i10 + i2)
	i        = rune(i10*i10 + i5)
	z        = rune(i10*i10 + i10*i2 + i2)
	b        = rune(i10*i10 - i2)
	u        = rune(i10*i10 + i10 + i7)
	fizz     = string([]rune{f, i, z, z})
	buzz     = string([]rune{b, u, z, z})
	fizzbuzz = fizz + buzz
)

func main() {
	for i := i1; i <= i10+i5; i++ {
		switch {
		case i%(i3*i5) == i0:
			println(fizzbuzz)
		case i%i3 == i0:
			println(fizz)
		case i%i5 == i0:
			println(buzz)
		default:
			println(i)
		}
	}
}
