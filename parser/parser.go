package parser

import (
	"github.com/nirlanka/evillang/ast"
	"github.com/nirlanka/evillang/lexer"
)

type Parser struct{}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{} // TODO: Set args

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	return program
}
