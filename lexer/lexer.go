package lexer

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	input string

	// nextPos ~ readPos + 1
	nextPos int // ready to read
	readPos int // read position

	nextCh byte // read character

	curStr string
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar() // set init position

	return l
}

// Set pos and ch

func (l *Lexer) readChar() {
	s := string(l.nextCh)
	fmt.Print(s) // DEBUG

	if l.readPos >= len(l.input) {
		l.nextCh = 0 // ASCII NUL --> end of input
	} else {
		l.nextCh = l.input[l.readPos]
	}

	l.nextPos = l.readPos
	l.readPos++
}

func (l *Lexer) resetStr() {
	l.curStr = ""
}

func (l *Lexer) readStrText() {
	l.readChar() // skip opening quote

	l.resetStr()

	start := l.nextPos
	for l.nextCh != '"' && l.nextCh != 0 {
		l.readChar()
	}

	l.curStr = l.input[start:l.nextPos]

	l.readChar() // skip closing quote
}

func (l *Lexer) readDecimalNumber() {
	l.resetStr()

	start := l.nextPos

	// Positive / negative / decimal-point / digit:

	decimalDotCount := 0

	if isDecimalpoint(l.nextCh) {
		decimalDotCount++
		l.readChar()
	}

	positiveNegativeCount := 0

	if isPositiveSym(l.nextCh) || isNegativeSym(l.nextCh) {
		positiveNegativeCount++
		l.readChar()
	}

	// The rest:

	for (decimalDotCount <= 1) &&
		(isDigit(l.nextCh) || isDecimalpoint(l.nextCh) || isUnderscore(l.nextCh)) {
		if isDecimalpoint(l.nextCh) {
			decimalDotCount++
		}

		l.readChar()
	}

	l.curStr = l.input[start:l.nextPos]
}

func (l *Lexer) readIdentifier() {
	l.resetStr()

	start := l.readPos - 1
	for isLetter(l.nextCh) || isDigit(l.nextCh) || isUnderscore(l.nextCh) {
		l.readChar()
	}

	l.curStr = l.input[start:l.nextPos]
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.nextCh) {
		l.readChar()
	}
}

// Tokens

func (l *Lexer) ReadToken() Token {
	l.skipWhitespace()
	l.resetStr()

	switch l.nextCh {
	case '=':
		l.readChar()
		return Token{Species: ASSIGN, Literal: "="}
	case ';':
		l.readChar()
		return Token{Species: SEMICOLON, Literal: ";"}
	case '"':
		l.readStrText()
		return Token{Species: INLINE_STRING, Literal: `"` + l.curStr + `"`}
	case 0:
		return Token{Species: EOF, Literal: ""}
	default:
		if isLetter(l.nextCh) || isUnderscore(l.nextCh) {
			l.readIdentifier()
			ident := l.curStr

			if tokSpecies, exists := Keywords[ident]; exists {
				return Token{Species: tokSpecies, Literal: ident}
			}

			if isPrimitiveType(ident) {
				return Token{Species: PRIMITIVE_TYPE, Literal: ident}
			}

			if isCustomTypeLike(ident) {
				return Token{Species: CUSTOM_TYPE, Literal: ident}
			}

			return Token{Species: IDENT, Literal: ident}
		} else if isPositiveSym(l.nextCh) ||
			isNegativeSym(l.nextCh) ||
			isDecimalpoint(l.nextCh) ||
			isDigit(l.nextCh) {
			l.readDecimalNumber()
			return Token{Species: DECIMAL_NUMBER, Literal: l.curStr}
		} else {
			return Token{Species: ILLEGAL, Literal: string(l.nextCh)}
		}
	}
}

// Helpers

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isUnderscore(ch byte) bool {
	return ch == '_'
}

func isStartsLowercase(s string) bool {
	if s == "" {
		return false
	}

	r := []rune(s)[0]
	return unicode.IsLower(r)
}

func isCustomTypeLike(ident string) bool {
	return isStartsUppercase(ident) || isStartsUnderscoreAndUppercase(ident)
}

func isPrimitiveType(ident string) bool {
	if tokSpecies, exists := Keywords[ident]; exists && tokSpecies == PRIMITIVE_TYPE {
		return true
	}

	return false
}

func isStartsUnderscoreAndLowercase(s string) bool {
	if s == "" {
		return false
	}

	rr := []rune(s)
	return rr[0] == '_' && unicode.IsLower(rr[1])
}

func isStartsUppercase(s string) bool {
	if s == "" {
		return false
	}

	r := []rune(s)[0]
	return unicode.IsUpper(r)
}

func isStartsUnderscoreAndUppercase(s string) bool {
	if s == "" {
		return false
	}

	rr := []rune(s)
	return rr[0] == '_' && unicode.IsUpper(rr[1])
}

func isDecimalpoint(ch byte) bool {
	return ch == '.'
}

func isPositiveSym(ch byte) bool {
	return ch == '+'
}

func isNegativeSym(ch byte) bool {
	return ch == '-'
}
