package ast

type Node interface {
	String() string
}

type Program struct {
	Statements []Node
}

type RefDeclaration struct {
	Name  string
	Value string
}

func (c *RefDeclaration) String() string {
	return "const " + c.Name + " = " + c.Value + ";"
}
