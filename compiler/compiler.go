package compiler

import (
	"strings"

	"github.com/nirlanka/evillang/ast"
)

func Compile(program *ast.Program) string {
	var out strings.Builder

	for _, statement := range program.Statements {
		out.WriteString(statement.String())
		out.WriteString("\n")
	}

	return out.String()
}
