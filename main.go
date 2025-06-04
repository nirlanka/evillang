package main

import (
	"fmt"

	"github.com/nirlanka/evillang/compiler"
	"github.com/nirlanka/evillang/parser"
)

func main() {
	input := `
		ref TBaseInteger a = 42;
		ref TBaseString b = "hi";
	`

	ast := parser.Parse(input)
	output := compiler.Compile(ast)

	fmt.Println("// Compiled JS:")
	fmt.Println("import * from 'eviltypes';")
	fmt.Println(output)
}
