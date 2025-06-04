package parser

import (
	"github.com/nirlanka/evillang/ast"
	"github.com/nirlanka/evillang/lexer"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}

	// Load two tokens to init - curToken, peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.curToken.Type != lexer.EOF {
		statement := p.parseRefDeclaration() // TODO

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseRefDeclaration() *ast.RefDeclaration {
	if p.curToken.Type != lexer.REF {
		return nil
	}

	// Except: TYPE
	p.nextToken()
	if p.curToken.Type != lexer.TYPE {
		return nil
	}
	typename := p.curToken.Literal

	// Expect: IDENT
	p.nextToken()
	if p.curToken.Type != lexer.IDENT {
		return nil
	}
	name := p.curToken.Literal

	// Expect: =
	p.nextToken()
	if p.curToken.Type != lexer.ASSIGN {
		return nil
	}

	// Expect: STRING or NUMBER
	p.nextToken()
	value := p.curToken.Literal

	// Expect: ;
	p.nextToken()
	if p.curToken.Type != lexer.SEMICOLON {
		return nil
	}

	return &ast.RefDeclaration{
		TypeName: typename,
		Name:     name,
		Value:    value,
	}
}
