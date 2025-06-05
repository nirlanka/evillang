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

var Keywords = map[string]TokenType{
	"ref":    REF,
	"string": TYPE,
	"int":    TYPE,
	"float":  TYPE,
	"bool":   TYPE,
	"object": TYPE,
}
