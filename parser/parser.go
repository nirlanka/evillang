package parser

import (
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
func (p *Parser) parseStatement() ast.Statement {
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

func (p *Parser) parseRefStatement() *ast.RefStatement {
	// Expects REF
	if p.curToken.Species != lexer.REF {
		return nil
	}

	// Expects: TYPE
	p.readToken()
	if p.curToken.Species != lexer.TYPE {
		return nil
	}
	typename := p.curToken.Literal

	// Expects: IDENT
	p.readToken()
	if p.curToken.Species != lexer.IDENT {
		return nil
	}
	name := p.curToken.Literal

	// Expects: =
	p.readToken()
	if p.curToken.Species != lexer.ASSIGN {
		return nil
	}

	// Value:
	// TODO: Handle single-word value initially, then switch to expressions later

	// Expects: STRING or NUMBER
	p.readToken()
	value := p.curToken.Literal

	// Expects: ;
	p.readToken()
	if p.curToken.Species != lexer.SEMICOLON {
		return nil
	}

	return &ast.RefStatement{
		TypeName: typename,
		Name:     name,
		Value:    value,
	}
}

// Program

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.curToken.Species != lexer.EOF {
		stment := p.parseStatement() // generic statement parser
		program.Statements = append(program.Statements, stment)
	}

	return program
}
