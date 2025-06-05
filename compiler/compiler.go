package compiler

import (
	"strings"

	"github.com/nirlanka/evillang/ast"
)

func Compile(program *ast.Program) string {
	var out strings.Builder
	out.WriteString("\n")

	for _, stment := range program.Statements {
		out.WriteString(stment.String())
		out.WriteString("\n")
	}

	return out.String()
}
