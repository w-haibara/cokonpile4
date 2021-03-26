package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(sougaku(10, 20, 520))
}

func sougaku(start, end, price int) string {
	str := "消費税が変更された場合の定価\n"

	for i := start; i <= end; i++ {
		str += "消費税が" +
			strconv.Itoa(i) +
			"%になった場合 : " +
			strconv.Itoa(int(math.Floor(float64(price)*((100.0+float64(i))/100.0)))) +
			"円\n"
	}
	return str
}
