package main

import (
	"fmt"
	"sample/base64"
)

func main() {
	plain := "ABCDEFG"
	fmt.Println("plain data :", plain)

	encoded := base64.Encode([]byte(plain))
	fmt.Println("encoded    :", encoded)

	fmt.Println("decoded    :", string(base64.Decode(encoded)))
}
