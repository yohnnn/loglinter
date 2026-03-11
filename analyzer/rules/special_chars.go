package rules

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func CheckSpecialChars(pass *analysis.Pass, msg string, lit *ast.BasicLit, _ *ast.CallExpr, _ int) {
	for _, r := range msg {

		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
			continue
		}

		pass.Reportf(
			lit.Pos(),
			"log message %q should not contain special punctuation",
			msg,
		)

		return
	}
}
