package lexer

import "fmt"

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

	l.curStr = ""

	start := l.curPos
	for l.curCh != '"' && l.curCh != 0 {
		l.readChar()
	}

	l.curStr = l.input[start:l.curPos]

	l.readChar() // skip closing quote
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
		return Token{Species: STRING, Literal: l.curStr}
	case 0:
		return Token{Species: EOF, Literal: ""}
	default:
		return Token{} // TODO: Replace this with identifier handling
	}
}

// Helpers

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
