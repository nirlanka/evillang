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

var TypeKeywordTargets = map[string]string{
	"string": "TBaseString",
	"int":    "TBaseInteger",
	"float":  "TBaseFloat",
	"bool":   "TBaseBoolean",
	"object": "TBaseObject",
}

func (r *RefDeclaration) String() string {
	var typeName string
	if tokType, ok := TypeKeywordTargets[r.TypeName]; ok {
		typeName = tokType
	} else {
		typeName = r.TypeName
	}

	return "/** @type {" + r.TypeName + "} */\n" + "const " + r.Name + " = new " + typeName + "().set(" + r.Value + ");"
}
