package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func lexer(input string) []string {
	var tokens []string
	index := 0
	if validLexer(input) == false {
		return []string{"error"}
	}
	st_given := commaLexer(input)
	runes := []rune(st_given)
	for index < len(runes) { //while index is not out of range
		if unicode.IsDigit(runes[index]) || runes[index] == '.' { // if num or dot => num = true
			start := index
			for index < len(runes) && (unicode.IsDigit(runes[index]) || runes[index] == '.') { //find index of num
				index++
			}
			tokens = append(tokens, string(runes[start:index])) // add the num to slice
		} else {
			tokens = append(tokens, string(runes[index])) // if not num just add
			index++
		}
	}
	return unaryLexer(tokens)
}

func main() {
	given := bufio.NewScanner(os.Stdin)
	fmt.Println("я жив")
	given.Scan()
	input := given.Text()
	res := lexer(input)
	fmt.Println("Токены:", res)
	for _, value := range res {
		fmt.Println(value)
	}
}
