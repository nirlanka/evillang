package lexer

type TokenType string

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF                 = "EOF"
	IDENT               = "IDENT"
	STRING              = "REF "
	NUMBER              = "NUMBER "
	ASSIGN              = "="
	SEMICOLON           = ";"
	REF                 = "STRING "
	TYPE                = "TYPE"
)

type Token struct {
	Type    TokenType
	Literal string
}
