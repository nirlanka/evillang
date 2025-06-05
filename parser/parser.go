package parser

import (
	"fmt"

	"github.com/nirlanka/evillang/ast"
	"github.com/nirlanka/evillang/lexer"
)

const (
	_ int = iota // TODO: Why?

	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // e.g. myFunc(X)
	BANG        // !
)

var precedences = map[lexer.TokenType]int{
	lexer.EQ:       EQUALS,
	lexer.NOT_EQ:   EQUALS,
	lexer.LT:       LESSGREATER,
	lexer.GT:       LESSGREATER,
	lexer.PLUS:     SUM,
	lexer.MINUS:    SUM,
	lexer.SLASH:    PRODUCT,
	lexer.ASTERISK: PRODUCT,
	lexer.LPAREN:   CALL,
	lexer.RPAREN:   CALL, // TODO: Is this needed?
	lexer.BANG:     BANG,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token

	errors []string

	prefixParseFns map[lexer.TokenType]prefixParseFn
	infixParseFns  map[lexer.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}

	// p.registerPrefix(lexer.IDENT, p.parseIdentifier)
	// p.registerPrefix(lexer.NUMBER, p.parseIntegerLiteral)
	// p.registerPrefix(lexer.BANG, p.parsePrefixExpression)
	// p.registerPrefix(lexer.MINUS, p.parsePrefixExpression)
	// p.registerPrefix(lexer.LPAREN, p.parseGroupedExpression)

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
		statement := p.parseRefDeclaration()
		// TODO: Add other types than ref-declarations

		if statement != nil {
			// fmt.Println("Parsed statement:", statement)
			program.Statements = append(program.Statements, statement)
		} else {
			p.nextToken()
		}
	}

	return program
}

// Handle statement type
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case lexer.REF:
		return p.parseRefDeclaration()
	case lexer.IDENT:
		// TODO: Check if peekToken has to change to peekTokenIs()
		if p.peekToken.Type == lexer.ASSIGN {
			return p.parseAssignmentStatement() // TODO: Implement
		}
		return p.parseExpressionStatement() // TODO: Implement
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseRefDeclaration() *ast.RefDeclarationStatement {
	if p.curToken.Type != lexer.REF {
		return nil
	}

	// Except: TYPE
	p.nextToken()
	if p.curToken.Type != lexer.TYPE {
		fmt.Printf("Parse error: expected TYPE after 'ref', got %q (%s)", p.curToken.Literal, p.curToken.Type)
		// TODO: Add error messages in other places too

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

	return &ast.RefDeclarationStatement{
		TypeName: typename,
		Name:     name,
		Value:    value,
	}
}

func (p *Parser) parseAssignmentStatement() *ast.AssignmentStatement {
	statement := &ast.AssignmentStatement{}

	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	p.nextToken() // consume IDENT
	if p.peekToken.Type != lexer.ASSIGN {
		fmt.Println("Parse error: expected peek token", lexer.ASSIGN, "but had", p.peekToken.Type)
		return nil
	}

	p.nextToken() // move to expression
	statement.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == lexer.SEMICOLON {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{}

	statement.Expression = p.parseExpression(LOWEST)

	if p.peekToken.Type == lexer.SEMICOLON {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		fmt.Println("Parse error: prefix parse failed over", p.curToken.Type)
		return nil
	}

	leftExpr := prefix()

	for p.peekToken.Type != lexer.SEMICOLON && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]

		if infix == nil {
			return leftExpr
		}

		p.nextToken() // move to infix operator
		leftExpr = infix(leftExpr)
	}

	return leftExpr
}

func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken() // move peek into curToken, fetch new peek
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
