package base64

import (
	"bytes"
	"strings"
)

const encodeStd string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const encodeURL string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
const encodedPad string = "="

func Encode(plain []byte) string {
	encoded := ""
	padNum := 0

	if n := len(plain) % 3; n != 0 {
		pad := make([]byte, 3-n)
		plain = append(plain, pad...)
		padNum = len(pad)
	}

	for i := 0; i < len(plain); i += 3 {
		tmp := [4]byte{
			(plain[i] & 0b11111100) >> 2,
			((plain[i] & 0b00000011) << 4) | ((plain[i+1] & 0b11110000) >> 4),
			((plain[i+1] & 0b00001111) << 2) | ((plain[i+2] & 0b11000000) >> 6),
			(plain[i+2] & 0b00111111),
		}

		encoded += string([]byte{encodeStd[tmp[0]], encodeStd[tmp[1]], encodeStd[tmp[2]], encodeStd[tmp[3]]})
	}

	if padNum != 0 {
		encoded = encoded[:len(encoded)-padNum]
		for i := 0; i < padNum; i++ {
			encoded += encodedPad
		}
	}

	return encoded
}

func Decode(encoded string) []byte {
	plain := []byte{}
	encodeStdByte := []byte(encodeStd)
	padNum := strings.Count(encoded, encodedPad)

	if len(encoded)%4 != 0 {
		panic("length of encoded string is invalid")
	}

	for i := 0; i < len(encoded); i += 4 {
		index := []int{
			bytes.IndexByte(encodeStdByte, encoded[i]),
			bytes.IndexByte(encodeStdByte, encoded[i+1]),
			bytes.IndexByte(encodeStdByte, encoded[i+2]),
			bytes.IndexByte(encodeStdByte, encoded[i+3]),
		}

		tmp := []byte{
			byte((index[0] << 2) | ((index[1] & 0b110000) >> 4)),
			byte(((index[1] & 0b001111) << 4) | ((index[2] & 0b111100) >> 2)),
			byte(((index[2] & 0b000011) << 6) | index[3]),
		}

		plain = append(plain, tmp...)
	}

	return plain[:len(plain)-padNum]
}
