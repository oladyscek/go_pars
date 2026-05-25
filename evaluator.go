package main

import "math"

type Node interface { 
	Eval() float64
}

type Number struct {
	Value float64
}

func (n Number) Eval() float64 {
	return n.Value
}

type BinaryExpr struct {
	op    string
	Left  Node
	Right Node
}

type UnaryErpr struct {
	op string
	Operand Node
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
	}
	return 0
}
func factorial(inp float64) float64 {
	if inp == 1{
		return 1
	}
	return inp * (factorial(inp-1))
}

func (u UnaryErpr) Eval() float64 {
	switch u.op {
	case "!":
		return factorial(u.Operand.Eval())
	case "√":
		return math.Sqrt(u.Operand.Eval())
	}
	return 0
}
