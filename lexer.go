package main

import "unicode"

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
		'=': {}, '(': {}, ')': {}, '^': {}, ' ': {},
		'!': {}, '√': {}, 'π': {}, '.': {}, ',': {},
	}
	for _, val := range input {
		if _, ok := validOps[val]; ok || unicode.IsDigit(val) { // if letter or another valid letter
			continue // check next symbol
		} else {
			return false // if not valid
		}
	}
	return true
}

