package lexer

type Token struct {
	Species TokenSpecies
	Literal string
}

type TokenSpecies string

// TokenType values:
const (
	ILLEGAL   TokenSpecies = "ILLEGAL"
	EOF                    = "EOF"
	IDENT                  = "IDENT"
	STRING                 = "STRING"
	NUMBER                 = "NUMBER"
	ASSIGN                 = "="
	SEMICOLON              = ";"
	REF                    = "REF"
	TYPE                   = "TYPE"
)

var Keywords = map[string]TokenSpecies{
	"ref":    REF,
	"string": TYPE,
	"int":    TYPE,
	"float":  TYPE,
	"bool":   TYPE,
	"object": TYPE,
}
