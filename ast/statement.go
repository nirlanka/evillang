package ast

import (
	"fmt"

	"github.com/nirlanka/evillang/lexer"
)

// Generic statement
type Statement interface {
	statementNode()
	String() string
}

type RefStatement struct {
	TypeName string
	// TODO: Later keep a type definition reference too

	Name  string
	Value string // all values (number, string)
}

func (r *RefStatement) statementNode() {}

func (r *RefStatement) String() string {
	targetType := lexer.PrimitiveTypes[r.TypeName]

	return fmt.Sprintf("/** @type {%s} */\nconst %s = new %s().set(%s);", targetType, r.Name, r.TypeName, r.Value)
}
