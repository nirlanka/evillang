package lexer

type Token struct {
	Species TokenSpecies
	Literal string
}

type TokenSpecies string

// TokenType values:
const (
	ILLEGAL        TokenSpecies = "ILLEGAL"
	EOF                         = "EOF"
	IDENT                       = "IDENT"
	INLINE_STRING               = "INLINE_STRING"
	DECIMAL_NUMBER              = "DECIMAL_NUMBER"
	ASSIGN                      = "="
	SEMICOLON                   = ";"
	REF                         = "REF"
	CUSTOM_TYPE                 = "CUSTOM_TYPE"
	PRIMITIVE_TYPE              = "PRIMITIVE_TYPE"
)

var Keywords = map[string]TokenSpecies{
	"ref":    REF,
	"string": PRIMITIVE_TYPE,
	"int":    PRIMITIVE_TYPE,
	"float":  PRIMITIVE_TYPE,
	"bool":   PRIMITIVE_TYPE,
	"map":    PRIMITIVE_TYPE,
}

var PrimitiveTypes = map[string]string{
	"string": "TString",
	"int":    "TInteger",
	"float":  "TFloat",
	"bool":   "TBoolean",
	"map":    "TMap",
}
