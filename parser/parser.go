package parser

import (
	"strings"

	"github.com/nirlanka/evillang/ast"
)

func Parse(input string) *ast.Program {
	lines := strings.Split(input, "\n")

	var statements []ast.Node

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "ref ") {
			parts := strings.Split(line, "=")

			name := strings.TrimSpace((strings.TrimPrefix(parts[0], "ref")))
			value := strings.TrimSpace(strings.TrimSuffix(parts[1], ";"))

			statements = append(statements, &ast.RefDeclaration{
				Name:  name,
				Value: value,
			})
		}
	}

	return &ast.Program{Statements: statements}
}
