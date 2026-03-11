package rules

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func CheckEnglish(pass *analysis.Pass, msg string, lit *ast.BasicLit, _ *ast.CallExpr, _ int) {

	for _, r := range msg {

		if unicode.IsLetter(r) {
			if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
				continue
			}

			pass.Reportf(
				lit.Pos(),
				"log message %q should be in English only",
				msg,
			)
			return
		}
	}
}
