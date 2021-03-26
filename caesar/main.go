package main

import (
	"fmt"
	"sample/caesar"
)

const n int = 20

func main() {
	plain := "This is a plain text!"
	fmt.Println("plain text :", plain)

	cipher := caesar.Enc(n, plain)
	fmt.Println("cipher     :", cipher)

	fmt.Println("decrypted  :", caesar.Dec(n, cipher))
}
