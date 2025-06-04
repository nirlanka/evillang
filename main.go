package main

import (
	"fmt"

	"github.com/nirlanka/evillang/compiler"
	"github.com/nirlanka/evillang/lexer"
	"github.com/nirlanka/evillang/parser"
)

func main() {
	// input := `
	// 	ref TBaseInteger a = 42;
	// 	ref TBaseString b = "hi";
	// `

	// ast := parser.Parse(input)
	// output := compiler.Compile(ast)

	// fmt.Println("// Compiled JS:")
	// fmt.Println("import * from 'eviltypes';")
	// fmt.Println(output)

	//// TEMP
	// input := `ref String foo = "abc";`

	// lex := lexer.New(input)

	// for tok := lex.NextToken(); tok.Type != lexer.EOF; tok = lex.NextToken() {
	// 	fmt.Printf("%+v\n", tok)
	// }

	//// TEMP
	source := `
	ref String foo="abc";
	ref Number x=42;
	`

	lex := lexer.New(source)
	par := parser.New(lex)
	prog := par.ParseProgram()

	js := compiler.Compile(prog)
	fmt.Println(js)
}
