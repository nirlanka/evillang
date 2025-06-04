package lexer

import "unicode"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII NUL -> end of input
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok := Token{Type: ASSIGN, Literal: string(l.ch)}
		l.readChar()
		return tok
	case ';':
		tok := Token{Type: SEMICOLON, Literal: string(l.ch)}
		l.readChar()
		return tok
	case '"':
		return l.readString()
	case 0:
		return Token{Type: EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()

			if tokType, ok := Keywords[ident]; ok {
				return Token{Type: tokType, Literal: ident}
			}

			return Token{Type: IDENT, Literal: ident}
		} else if isDigit(l.ch) {
			return l.readNumber()
		} else {
			return Token{Type: ILLEGAL, Literal: string(l.ch)}
		}
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() Token {
	l.readChar() // skip opening quote

	start := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}

	lit := l.input[start:l.position]

	l.readChar() // skip closing quote

	return Token{Type: STRING, Literal: `"` + lit + `"`}
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[start:l.position]
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() Token {
	start := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return Token{Type: NUMBER, Literal: l.input[start:l.position]}
}
