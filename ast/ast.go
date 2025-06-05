package ast

import (
	"fmt"
	"strings"

	"github.com/nirlanka/evillang/lexer"
)

// Abstract entities:
// ==================

type Node interface {
	String() string
}

type Expression interface {
	expressionNode()
	String() string
}

type Statement interface {
	statementNode()
	String() string
}

type Program struct {
	Statements []Node
}

// Entities:
// =========

// Identifier:

type Identifier struct {
	Token lexer.Token
	Value string
}

func (id *Identifier) String() string {
	return id.Value
}

// Integer:

type IntegerLiteral struct {
	Token lexer.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {} // TODO: ?

func (il *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}

// Infix:

type InfixExpression struct {
	Token    lexer.Token // operator tokens, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {} // TODO: ?

func (ie *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
}

// Prefix:

type PrefixExpression struct {
	Token    lexer.Token // e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {} // TODO: ?

func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", pe.Operator, pe.Right.String())
}

// Call expression:

type CallExpression struct {
	Token     lexer.Token // the `(` token
	Function  Expression  // identifier / nested identifier e.g. obj.method
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {} // TODO: ?

func (ce *CallExpression) String() string {
	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	return fmt.Sprintf("%s(%s)", ce.Function.String(), strings.Join(args, ","))
}

// Statements:
// ===========

// Expression statement:

type ExpressionStatement struct {
	Expression Expression
}

func (as *ExpressionStatement) statementNode() {} // TODO: ?

func (as *ExpressionStatement) String() string {
	return as.Expression.String()
}

// Ref Declaration statement:

type RefDeclarationStatement struct {
	TypeName string
	Name     string
	Value    string
}

var TypeKeywordTargets = map[string]string{
	"string": "TBaseString",
	"int":    "TBaseInteger",
	"float":  "TBaseFloat",
	"bool":   "TBaseBoolean",
	"object": "TBaseObject",
}

func (as *RefDeclarationStatement) statementNode() {} // TODO: ?

func (r *RefDeclarationStatement) String() string {
	var typeName string
	if tokType, ok := TypeKeywordTargets[r.TypeName]; ok {
		typeName = tokType
	} else {
		typeName = r.TypeName
	}

	return "/** @type {" + r.TypeName + "} */\n" + "const " + r.Name + " = new " + typeName + "().set(" + r.Value + ");"
}

// Asssignment statement:

type AssignmentStatement struct {
	Name  *Identifier
	Value Expression
}

func (as *AssignmentStatement) statementNode() {} // TODO: ?

func (as *AssignmentStatement) String() string {
	return fmt.Sprintf("%s.set(%s);", as.Name.String(), as.Value.String())
}
