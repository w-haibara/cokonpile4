package kenshiro

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

const (
	V1 string = "あ"
	V2 string = "た"
)

func ToFloat(atata string) float32 {
	fmt.Println("==== ToFloat ====")
	fmt.Printf("atata : %v\n", atata)

	if utf8.RuneCountInString(atata) != 32 {
		panic("length of input string must be 32")
	}

	bytes := make([]byte, 4)

	c := 0
	for _, v := range atata {
		var bit byte

		switch string(v) {
		case V1:
			bit = 0b0
		case V2:
			bit = 0b1
		default:
			panic(fmt.Sprintf("unknown rune: [%v] at %v", string(v), c))
		}

		bytes[c/8] |= bit << (7 - (c % 8))

		c++
	}

	print("bits  : ")
	for _, v := range bytes {
		fmt.Printf("%08v ", strconv.FormatInt(int64(v), 2))
	}
	println()

	bits := binary.BigEndian.Uint32(bytes)
	f := math.Float32frombits(bits)

	fmt.Printf("float : %v\n", f)

	fmt.Println("====         ====")

	return f
}

func ToKenshiro(f float32) string {
	fmt.Println("==== ToKenshiro ====")

	fmt.Printf("float : %v\n", f)

	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	print("bits  : ")
	atata := ""
	bytes := buf.Bytes()
	for _, v := range bytes {
		fmt.Printf("%08v ", strconv.FormatInt(int64(v), 2))
		for i := 0; i < 8; i++ {
			switch int((v >> (7 - i)) & 0b00000001) {
			case 0:
				atata += V1
			case 1:
				atata += V2
			}
		}
	}
	println()

	fmt.Printf("atata : %v\n", atata)

	fmt.Println("====            ====")

	return atata
}
