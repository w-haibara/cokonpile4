package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  uint
	out string
}

func Test_Render(t *testing.T) {
	testCases := []TestCase{
		{
			in:  3,
			out: "Fizz",
		},
		{
			in:  5,
			out: "Buzz",
		},
		{
			in:  15,
			out: "FizzBuzz",
		},
		{
			in:  1,
			out: "1",
		},
	}

	for _, testCase := range testCases {
		fb := Default()
		str := fb.Say(testCase.in)
		assert.Equal(t, testCase.out, str)
	}
}
