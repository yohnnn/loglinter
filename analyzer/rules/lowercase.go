package rules

import (
	"go/ast"
	"strconv"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func CheckLowerCase(pass *analysis.Pass, msg string, lit *ast.BasicLit, _ *ast.CallExpr, _ int) {

	if msg == "" {
		return
	}

	r := []rune(msg)[0]

	if unicode.IsUpper(r) {
		runes := []rune(msg)
		runes[0] = unicode.ToLower(runes[0])
		fixedMsg := string(runes)

		pass.Report(analysis.Diagnostic{
			Pos:     lit.Pos(),
			End:     lit.End(),
			Message: "log message " + strconv.Quote(msg) + " should start with a lowercase letter",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "convert first letter to lowercase",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(strconv.Quote(fixedMsg)),
						},
					},
				},
			},
		})
	}
}
