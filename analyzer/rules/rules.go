package rules

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

type Rule func(pass *analysis.Pass, msg string, lit *ast.BasicLit, call *ast.CallExpr, msgIndex int)

var All = []Rule{
	CheckLowerCase,
	CheckEnglish,
	CheckSpecialChars,
	CheckSensitive,
}

var AllWithoutSensitive = []Rule{
	CheckLowerCase,
	CheckEnglish,
	CheckSpecialChars,
}
