package main

import (
	"fmt"
	"sample/fizzbuzz"
)

func main() {
	fb := fizzbuzz.Default()
	for i := uint(1); i <= 15; i++ {
		fmt.Println(fb.Say(i))
	}
}
