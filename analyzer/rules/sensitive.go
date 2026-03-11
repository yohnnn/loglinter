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

var customSensitivePatterns []string

func SetSensitivePatterns(patterns []string) {
	customSensitivePatterns = customSensitivePatterns[:0]

	for _, p := range patterns {
		p = strings.TrimSpace(strings.ToLower(p))
		if p == "" {
			continue
		}

		customSensitivePatterns = append(customSensitivePatterns, p)
	}
}

func effectiveSensitivePatterns() []string {
	if len(customSensitivePatterns) == 0 {
		return sensitiveKeywords
	}

	patterns := make([]string, 0, len(sensitiveKeywords)+len(customSensitivePatterns))
	patterns = append(patterns, sensitiveKeywords...)
	patterns = append(patterns, customSensitivePatterns...)

	return patterns
}

func containsSensitiveKeyword(s string) bool {
	lower := strings.ToLower(s)

	for _, keyword := range effectiveSensitivePatterns() {
		if strings.Contains(lower, keyword) {
			return true
		}
	}

	return false
}

func CheckSensitive(pass *analysis.Pass, msg string, lit *ast.BasicLit, call *ast.CallExpr, msgIndex int) {
	lower := strings.ToLower(msg)

	for _, keyword := range effectiveSensitivePatterns() {
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
