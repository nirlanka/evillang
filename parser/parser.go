package parser

import (
	"fmt"
	"strings"

	"github.com/nirlanka/evillang/ast"
	"github.com/nirlanka/evillang/lexer"
)

type Parser struct {
	lexer *lexer.Lexer

	curToken  lexer.Token
	nextToken lexer.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}

	// load initial (2) tokens
	p.readToken()
	p.readToken()

	return p
}

// Set token in state

func (p *Parser) readToken() {
	p.curToken = p.nextToken
	p.nextToken = p.lexer.ReadToken()
}

// Statements

// Returns Statement, not pointer, so it doesn't conflict with
// specific variation of Statement
func (p *Parser) parseStatement() (ast.Statement, *string) {
	switch p.curToken.Species {
	case lexer.REF:
		return p.parseRefStatement()
	// case lexer.IDENT:
	//// 	TODO: Handler other statements
	default:
		// return p.parseExpressionStatement()
		// TODO: Implement above

		return p.parseRefStatement() // TODO: Remove dummy
	}
}

func (p *Parser) parseRefStatement() (*ast.RefStatement, *string) {
	// Expects REF
	if p.curToken.Species != lexer.REF {
		return nil, p.error([]string{lexer.REF}, string(p.curToken.Species))
	}

	// Expects: TYPE
	p.readToken()
	if p.curToken.Species != lexer.CUSTOM_TYPE && p.curToken.Species != lexer.PRIMITIVE_TYPE {
		return nil, p.error([]string{lexer.CUSTOM_TYPE, lexer.PRIMITIVE_TYPE}, string(p.curToken.Species))
	}
	typename := p.curToken.Literal

	// Expects: IDENT
	p.readToken()
	if p.curToken.Species != lexer.IDENT {
		return nil, p.error([]string{lexer.IDENT}, string(p.curToken.Species))
	}
	name := p.curToken.Literal

	// Expects: =
	p.readToken()
	if p.curToken.Species != lexer.ASSIGN {
		return nil, p.error([]string{lexer.ASSIGN}, string(p.curToken.Species))
	}

	// Value:
	// TODO: Handle single-word value initially, then switch to expressions later

	// Expects: STRING or NUMBER
	p.readToken()
	value := p.curToken.Literal

	// Expects: ;
	p.readToken()
	if p.curToken.Species != lexer.SEMICOLON {

		return nil, p.error([]string{lexer.SEMICOLON}, string(p.curToken.Species))
	}

	p.readToken()

	return &ast.RefStatement{
		TypeName: typename,
		Name:     name,
		Value:    value,
	}, nil
}

func (p *Parser) error(expected []string, got string) *string {
	err := fmt.Sprintf("Parse Error: Expected [%s], but got (%s)", strings.Join(expected, ", "), got)

	return &err
}

// Program

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.curToken.Species != lexer.EOF {
		stment, err := p.parseStatement() // generic statement parser
		if err != nil {
			fmt.Println()
			panic(fmt.Sprintf("\nError: \n- %s\n]", *err))
		}
		program.Statements = append(program.Statements, stment)
	}

	return program
}
