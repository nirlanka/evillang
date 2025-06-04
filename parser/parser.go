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
			// Format: ref Type name = value;
			rest := strings.TrimPrefix(line, "ref ")
			tokens := strings.Fields(rest) // ["String", "foo", "=", "\"abc\";"]

			if len(tokens) >= 4 && tokens[2] == "=" {
				typename := tokens[0]
				name := tokens[1]
				value := strings.TrimSuffix(strings.Join(tokens[3:], " "), ";")

				statements = append(statements, &ast.RefDeclaration{
					TypeName: typename,
					Name:     name,
					Value:    value,
				})
			}
		}
	}

	return &ast.Program{Statements: statements}
}
