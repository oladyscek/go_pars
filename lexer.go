package main

import (
	"strings"
	"unicode"
)

func commaLexer(st_input string) string {
	sl := make([]rune, 0, len(st_input))
	for _, val := range st_input {
		if val == ',' { // comma to dot
			sl = append(sl, '.')
			continue
		}
		if val == ' ' { //ignore spaces
			continue
		} else { //just normal num
			sl = append(sl, val)
		}
	}
	str := string(sl)
	return str
}

func validLexer(input string) bool {
	var validOps = map[rune]struct{}{ // valid letters except numbers
		'+': {}, '-': {}, '*': {}, '/': {},
		'(': {}, ')': {}, '^': {}, ' ': {},
		'!': {}, '√': {}, 'π': {}, '.': {}, ',': {},
	}
	for _, val := range input {
		if _, ok := validOps[val]; ok || unicode.IsDigit(val) { // if number or another valid letter
			continue // check next symbol
		} else {
			return false // if not valid
		}
	}
	return true
}

func unaryLexer(input []string) []string {
	sls := make([]string, 0, len(input)) //output slice
	i := 0
	var unary = map[string]struct{}{ // set of unary triggers (before unary minus)
		"+": {}, "-": {}, "*": {}, "/": {},
		"(": {}, "^": {},
	}
	for i < len(input) { 
		if i == 0 && input[i] == "-" { // if first is minus => unar
			var sb strings.Builder 
			sb.WriteString("-") // building a new token
			sb.WriteString(input[i+1])
			sls = append(sls, sb.String()) // add new token
			i += 2 // +2 because index 0 and 1 are the unary minus and the number
		} else { // if index is not 0
			if _, ok := unary[input[i-1]]; input[i] == "-" && ok { // if minus and token before in unary triggers 
				var sb strings.Builder
				sb.WriteString("-") // building a new token
				sb.WriteString(input[i+1])
				sls = append(sls, sb.String()) // add new token
				i += 2
			} else { // if num just add and go to next token
				sls = append(sls, input[i]) 
				i++
			}
		}
	}
	return sls
}
