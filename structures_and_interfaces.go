package main

type Node interface {
	Eval() float64
}

type Number struct {
	Value float64
}

type ExprNode struct {
	Op       string
	Children []Node
}

type BinaryExpr struct {
	op    string
	Left  Node
	Right Node
}

type UnaryExpr struct {
	op      string
	Operand Node
}
