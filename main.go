package main

import (
	"fmt"

	"github.com/nirlanka/evillang/compiler"
	"github.com/nirlanka/evillang/lexer"
	"github.com/nirlanka/evillang/parser"
)

func main() {
	// source := `
	// 	ref String foo = "abc";
	// 	ref Integer x = 42;
	// 	ref Float y = 1.2;
	// 	ref ApiApp app = 42;
	// 	`
	source := `
		ref string foo = "abc";
		ref int x = 42;
		ref ApiApp app = 42;

		x = 5;
		`

	// TODO: Ignore // lines
	// TODO: Handle Float values

	lex := lexer.New(source)

	par := parser.New(lex)
	prog := par.ParseProgram()

	js := compiler.Compile(prog)
	fmt.Println(js)
}
