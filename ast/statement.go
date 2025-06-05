package ast

import "fmt"

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
	// TODO: Add type alias resolution/type-name transformations if needed

	return fmt.Sprintf("/** @type {%s} */\nconst %s = new %s().set(%s);", r.TypeName, r.Name, r.TypeName, r.Value)
}
