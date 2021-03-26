package base64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  string
	out string
}

func Test_Render(t *testing.T) {
	testCases := []TestCase{
		{
			in:  "",
			out: "",
		},
		{
			in:  "f",
			out: "Zg==",
		},
		{
			in:  "fo",
			out: "Zm8=",
		},
		{
			in:  "foo",
			out: "Zm9v",
		},
		{
			in:  "foob",
			out: "Zm9vYg==",
		},
		{
			in:  "fooba",
			out: "Zm9vYmE=",
		},
		{
			in:  "foobar",
			out: "Zm9vYmFy",
		},
	}

	for _, testCase := range testCases {
		str := Encode([]byte(testCase.in))
		assert.Equal(t, testCase.out, str)
		str = string(Decode(str))
		assert.Equal(t, str, testCase.in)
	}
}
