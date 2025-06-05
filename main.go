package main

import (
	"fmt"

	"github.com/nirlanka/evillang/compiler"
	"github.com/nirlanka/evillang/lexer"
	"github.com/nirlanka/evillang/parser"
)

func main() {
	source := `
		ref string foo = "abc";
		ref int x = 42;
		ref float y = 1.3;
		ref int z = -99_999;
	`

	lex := lexer.New(source)
	par := parser.New(lex)

	prog := par.ParseProgram()
	js := compiler.Compile(prog)

	fmt.Println(js)
}
