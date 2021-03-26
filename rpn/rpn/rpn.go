package rpn

import (
	"fmt"
	"strconv"
	"strings"
)

type Expr string

func ToExpr(str string) Expr {
	return Expr(strings.TrimSpace(str))
}

func (e Expr) Tokenize() Token {
	return Token(strings.Fields(string(e)))
}

func (e Expr) Sprint() string {
	return fmt.Sprintf("expr: [%v]", e)
}

func (e Expr) Print() {
	fmt.Println(e.Sprint())
}

type Token []string

func (t Token) Sprint() string {
	str := "token: "
	for _, v := range t {
		str += fmt.Sprintf("[%v] ", v)
	}
	return str
}

func (t Token) Print() {
	fmt.Println(t.Sprint())
}

func (t Token) Consume() (Token, string) {
	return t[1:], t[0]
}

type Stack []string

func (s Stack) Push(elm string) Stack {
	return append(Stack{elm}, s...)
}

func (s Stack) PushFloat(elm float64) Stack {
	return s.Push(fmt.Sprintf("%v", elm))
}

func (s Stack) Pop() (Stack, string) {
	return s[1:], s[0]
}

func (s Stack) PopFloat() (Stack, float64) {
	s2, res := s.Pop()
	f, err := strconv.ParseFloat(res, 64)
	if err != nil {
		panic(err)
	}
	return s2, f
}

func (s Stack) Sprint() string {
	return fmt.Sprint("stack:", s)
}

func (s Stack) Print() {
	fmt.Println(s.Sprint())
}

type Operator map[string](func(Stack) Stack)

func (op Operator) Eval(exp string) (string, error) {
	e := ToExpr(exp)
	e.Print()

	fmt.Println("\n=== tokenize ===\n")
	t := e.Tokenize()
	t.Print()

	fmt.Println("\n=== parse ===\n")
	var s Stack
	for range t {
		var val string
		t, val = t.Consume()
		if _, err := strconv.ParseFloat(val, 64); err == nil {
			fmt.Println("[number]", val)
			s = s.Push(val)
			s.Print()
			println()
		} else {
			fmt.Println("[symbol]", val)
			v, ok := op[val]
			if !ok {
				return "", fmt.Errorf("unknown sybol: %v", val)
			}
			s = v(s)
			s.Print()
			println()
		}
	}

	if len(s) != 1 {
		return "", fmt.Errorf("stack length is not 1; %v", s.Sprint())
	}

	return s[0], nil
}
