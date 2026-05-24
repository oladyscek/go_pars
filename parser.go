package main

import "strconv"

func pres(op string) int {
	switch op {
	case "+", "-": // have the lowest pres
		return 1
	case "*", "/": // have the higher pres
		return 2
	case "^", "!", "√": // have the highest pres
		return 3
	}
	return 9999
}

func isOperator(inp string) bool {
	var operators = map[string]struct{}{
		"+": {}, "-": {}, "*": {}, "/": {},
		"^": {}, "!": {}, "√": {},
	}
	_, ok := operators[inp] // ok returns true if there is such key in operators
	return ok
}

func root(input []string) int {
	lowestpres := 9999
	root_pres := -1
	depth := 0 // operator in () is not root

	for i, tok := range input {
		if tok == "(" {
			depth++
			continue
		}
		if tok == ")" {
			depth--
			continue
		}
		if depth == 0 {
			if isOperator(tok) {
				prec := pres(tok)       // pres of current operator
				if prec <= lowestpres { // if lover
					lowestpres = prec
					root_pres = i
				}
			}
		}
	}
	return root_pres
}

func parse(input []string) Node {
	if len(input) == 1 { // if just 1 num return it
		val, _ := strconv.ParseFloat(input[0], 64)
		return Number{
			Value: val,
		}
	}

	if input[0] == "(" && input[len(input)-1] == ")" { // could be like (2+2), i know it mean that user is fool but we should parse it
		return parse(input[1 : len(input)-1]) // P.S. BUG!! case (2+3)+(4+5) will crash
	}

	root_ind := root(input) // find the root
	left_node := parse(input[:root_ind]) // parse left operand recursivly 
	right_node := parse(input[root_ind+1:]) // parse right operand recursivly

	return BinaryExpr{ 
		op:    input[root_ind],
		Left:  left_node,
		Right: right_node,
	}
}
