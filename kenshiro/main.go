package main

import (
	"fmt"
	"sample/kenshiro"
)

const f float32 = 25.625

func main() {
	fmt.Printf("float : %v\n\n", f)

	atata := kenshiro.ToKenshiro(f)
	fmt.Printf("ToKenshiro : %v\n\n", atata)

	f2 := kenshiro.ToFloat(atata)
	fmt.Printf("ToFloat : %v\n", f2)
}
