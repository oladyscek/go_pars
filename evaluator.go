package main

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
