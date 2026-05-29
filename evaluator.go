package main

import "math"

func (n Number) Eval() float64 {
	return n.Value
}

func (b BinaryExpr) Eval() float64 {
	switch b.op {
	case "+":
		return b.Left.Eval() + b.Right.Eval()
	case "-":
		return b.Left.Eval() - b.Right.Eval()
	case "*":
		return b.Left.Eval() * b.Right.Eval()
	case "/":
		return b.Left.Eval() / b.Right.Eval()
	case "^":
		return math.Pow(
			b.Left.Eval(),
			b.Right.Eval(),
		) 
	}
	panic("unknown operator")
}
func factorial(inp float64) float64 {
	if inp < 0 {
		panic("negative factorial")
	}
	if inp != math.Trunc(inp) {
		panic("float factorial")
	}
	if inp == 1 || inp == 0 {
		return 1
	}
	return inp * (factorial(inp - 1))
}

func (u UnaryExpr) Eval() float64 {
	switch u.op {
	case "!":
		return factorial(u.Operand.Eval())
	case "√":
		return math.Sqrt(u.Operand.Eval())
	}
	return 0
}
