package main

import (
	"fmt"

	"github.com/nirlanka/evillang/compiler"
	"github.com/nirlanka/evillang/parser"
)

func main() {
	input := `
		ref a = 42;
		ref b = "hi";
	`

	ast := parser.Parse(input)
	output := compiler.Compile(ast)

	fmt.Println("// Compiled JS:")
	fmt.Println(output)
}
