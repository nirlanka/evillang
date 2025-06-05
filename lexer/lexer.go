package lexer

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	input string

	nextPos int // ~ curPos + 1
	curPos  int // currently reading

	curCh byte

	curStr string
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar() // set init position

	return l
}

// Set pos and ch

func (l *Lexer) readChar() {
	if l.curPos >= len(l.input) {
		l.curCh = 0 // ACII NUL --> end of input
	} else {
		l.curCh = l.input[l.curPos]
	}

	l.curPos = l.nextPos
	l.nextPos++
}

func (l *Lexer) resetStr() {
	l.curStr = ""
}

func (l *Lexer) readStr() {
	l.readChar() // skip opening quote

	l.resetStr()

	start := l.curPos
	for l.curCh != '"' && l.curCh != 0 {
		l.readChar()
	}

	l.curStr = l.input[start:l.curPos]

	l.readChar() // skip closing quote
}

func (l *Lexer) readDecimalNumber() {
	l.resetStr()

	start := l.curPos

	// Positive / negative / decimal-point / digit:

	decimalDotCount := 0

	if isDecimalpoint(l.curCh) {
		decimalDotCount++
		l.readChar()
	}

	positiveNegativeCount := 0

	if isPositiveSym(l.curCh) || isNegativeSym(l.curCh) {
		positiveNegativeCount++
		l.readChar()
	}

	// The rest:

	for (decimalDotCount <= 1) &&
		(isDigit(l.curCh) || isDecimalpoint(l.curCh) || isUnderscore(l.curCh)) {
		if isDecimalpoint(l.curCh) {
			decimalDotCount++
		}

		l.readChar()
	}

	l.curStr = l.input[start:l.curPos]
}

func (l *Lexer) readIdentifier() {
	l.resetStr()

	start := l.curPos
	for isLetter(l.curCh) || isDigit(l.curCh) || isUnderscore(l.curCh) {
		l.readChar()
	}

	l.curStr = l.input[start:l.curPos]
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.curCh) {
		s := string(l.curCh)
		fmt.Println(s)
		l.readChar()
	}
}

// Tokens

func (l *Lexer) ReadToken() Token {
	l.skipWhitespace()
	l.resetStr()

	switch l.curCh {
	case '=':
		l.readChar()
		return Token{Species: ASSIGN, Literal: "="}
	case ';':
		l.readChar()
		return Token{Species: SEMICOLON, Literal: ";"}
	case '"':
		l.readStr()
		return Token{Species: INLINE_STRING, Literal: l.curStr}
	case 0:
		return Token{Species: EOF, Literal: ""}
	default:
		if isLetter(l.curCh) || isUnderscore(l.curCh) {
			l.readIdentifier()
			ident := l.curStr

			if tokSpecies, ok := Keywords[ident]; ok {
				return Token{Species: tokSpecies, Literal: ident}
			}

			if isStartsUppercase(ident) || isStartsUnderscoreAndUppercase(ident) {
				return Token{Species: TYPE, Literal: ident}
			}

			return Token{Species: IDENT, Literal: ident}
		} else if isDigit(l.curCh) {
			l.readDecimalNumber()
			return Token{Species: DECIMAL_NUMBER, Literal: l.curStr}
		} else {
			return Token{Species: ILLEGAL, Literal: string(l.curCh)}
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
