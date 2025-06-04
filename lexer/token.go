package lexer

type TokenType string

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF                 = "EOF"
	IDENT               = "IDENT"
	STRING              = "STRING"
	NUMBER              = "NUMBER "
	ASSIGN              = "="
	SEMICOLON           = ";"
	REF                 = "REF "
	TYPE                = "TYPE"
)

type Token struct {
	Type    TokenType
	Literal string
}
