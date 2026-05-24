package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	given := bufio.NewScanner(os.Stdin)
	fmt.Println("введите выражение")
	given.Scan()
	input := given.Text()
	tokens := lexer(input)
	ast := parse(tokens)
	fmt.Println(ast.Eval())
}
