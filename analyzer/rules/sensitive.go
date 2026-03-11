package rules

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var sensitiveKeywords = []string{
	"password",
	"pass",
	"token",
	"secret",
	"apikey",
	"api_key",
	"pwd",
	"credential",
}

func containsSensitiveKeyword(s string) bool {
	lower := strings.ToLower(s)

	for _, keyword := range sensitiveKeywords {
		if strings.Contains(lower, keyword) {
			return true
		}
	}

	return false
}

func CheckSensitive(pass *analysis.Pass, msg string, lit *ast.BasicLit, call *ast.CallExpr, msgIndex int) {
	lower := strings.ToLower(msg)

	for _, keyword := range sensitiveKeywords {
		if strings.Contains(lower, keyword) {
			pass.Reportf(
				lit.Pos(),
				"log message contains potentially sensitive keyword %q",
				keyword,
			)
			break
		}
	}

	for i, arg := range call.Args {
		if i == msgIndex {
			continue
		}

		ident, ok := arg.(*ast.Ident)
		if !ok {
			continue
		}

		if containsSensitiveKeyword(ident.Name) {
			pass.Reportf(
				ident.Pos(),
				"log argument %q may contain sensitive data",
				ident.Name,
			)
		}
	}
}
