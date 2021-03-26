package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sample/rpn"
	"strings"
)

func main() {
	op := rpn.DefaultOps
	op["inc"] = inc
	op["dec"] = dec
	op["sqrt"] = sqrt
	op["exit"] = exit

	for {
		fmt.Print("> ")

		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		expr := strings.TrimSpace(stdin.Text())
		if expr == "" {
			continue
		}

		res, err := op.Eval(expr)
		if err != nil {
			fmt.Println("[error:", err, "]")
			break
		}
		fmt.Println("result:", res)
	}
}

func inc(s rpn.Stack) rpn.Stack {
	var v float64
	s, v = s.PopFloat()
	return s.PushFloat(v + 1)
}

func dec(s rpn.Stack) rpn.Stack {
	var v float64
	s, v = s.PopFloat()
	return s.PushFloat(v - 1)
}

func sqrt(s rpn.Stack) rpn.Stack {
	var v float64
	s, v = s.PopFloat()
	return s.PushFloat(math.Sqrt(v))
}

func exit(s rpn.Stack) rpn.Stack {
	os.Exit(0)
	return rpn.Stack{}
}
