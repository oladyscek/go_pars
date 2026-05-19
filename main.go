package main

import (
	"bufio"
	"fmt"
	"os"
)

func not_main() {
	given := bufio.NewScanner(os.Stdin)
	fmt.Println("введите выражение")
	given.Scan()
	input := given.Text()
	res := lexer(input)
	fmt.Println("Токены:", res)
}
