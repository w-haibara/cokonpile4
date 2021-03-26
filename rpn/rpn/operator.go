package rpn

var DefaultOps = Operator{
	"+": Add,
	"-": Sub,
	"*": Mul,
	"/": Div,
}

func Add(s Stack) Stack {
	var (
		v1 float64
		v2 float64
	)
	s, v1 = s.PopFloat()
	s, v2 = s.PopFloat()
	return s.PushFloat(v2 + v1)
}

func Sub(s Stack) Stack {
	var (
		v1 float64
		v2 float64
	)
	s, v1 = s.PopFloat()
	s, v2 = s.PopFloat()
	return s.PushFloat(v2 - v1)
}

func Mul(s Stack) Stack {
	var (
		v1 float64
		v2 float64
	)
	s, v1 = s.PopFloat()
	s, v2 = s.PopFloat()
	return s.PushFloat(v2 * v1)
}

func Div(s Stack) Stack {
	var (
		v1 float64
		v2 float64
	)
	s, v1 = s.PopFloat()
	s, v2 = s.PopFloat()
	return s.PushFloat(v2 / v1)
}
