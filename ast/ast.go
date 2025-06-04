package ast

type Node interface {
	String() string
}

type Program struct {
	Statements []Node
}

type RefDeclaration struct {
	TypeName string
	Name     string
	Value    string
}

func (r *RefDeclaration) String() string {
	return "/** @type {" + r.TypeName + "} */\n" + "const " + r.Name + " = new " + r.TypeName + "().set(" + r.Value + ");"
}
