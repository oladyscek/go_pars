package main

func pres(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^", "!", "√":
		return 3
	}
	return 9999
}

func isOperator(inp string) bool {
	var operators = map[string]struct{}{
		"+": {}, "-": {}, "*": {}, "/": {},
		"^": {}, "!": {}, "√": {},
	}
	_, ok := operators[inp]
	return ok
}

func root(input []string) int {
	lowestpres := 999
	root_pres := -1
	depth := 0

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
				prec := pres(tok)
				if prec <= lowestpres {
					lowestpres = prec
					root_pres = i
				}
			}
		}
	}
	return root_pres
}
