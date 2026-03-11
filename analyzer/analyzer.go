package analyzer

import (
	"go/ast"
	"go/types"
	"strconv"

	"github.com/yohnnn/loglinter/analyzer/rules"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type Config struct {
	EnableSensitive bool
}

func DefaultConfig() Config {
	return Config{
		EnableSensitive: true,
	}
}

var slogMethods = map[string]int{
	"Info":         0,
	"Error":        0,
	"Warn":         0,
	"Debug":        0,
	"InfoContext":  1,
	"ErrorContext": 1,
	"WarnContext":  1,
	"DebugContext": 1,
}

var zapMethods = map[string]int{
	"Info":  0,
	"Error": 0,
	"Warn":  0,
	"Debug": 0,
	"Fatal": 0,
	"Panic": 0,
}

func New() *analysis.Analyzer {
	return NewWithConfig(DefaultConfig())
}

func NewWithConfig(cfg Config) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "loglint",
		Doc:  "checks log messages for formatting and security issues",
		Run: func(pass *analysis.Pass) (any, error) {
			return run(pass, cfg)
		},
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}
}

func run(pass *analysis.Pass, cfg Config) (any, error) {

	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {

		call := n.(*ast.CallExpr)

		handleCall(pass, call, cfg)

	})

	return nil, nil
}

func handleCall(pass *analysis.Pass, call *ast.CallExpr, cfg Config) {

	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	obj := pass.TypesInfo.Uses[sel.Sel]
	if obj == nil {
		return
	}

	fn, ok := obj.(*types.Func)
	if !ok {
		return
	}

	pkg := fn.Pkg()
	if pkg == nil {
		return
	}

	switch pkg.Path() {

	case "log/slog":
		msgIndex, ok := slogMethods[fn.Name()]
		if !ok {
			return
		}

		checkLog(pass, call, msgIndex, cfg)

	case "go.uber.org/zap":
		msgIndex, ok := zapMethods[fn.Name()]
		if !ok {
			return
		}

		checkLog(pass, call, msgIndex, cfg)
	}
}

func checkLog(pass *analysis.Pass, call *ast.CallExpr, msgIndex int, cfg Config) {

	if len(call.Args) <= msgIndex {
		return
	}

	arg := call.Args[msgIndex]

	lit, ok := arg.(*ast.BasicLit)
	if !ok {
		return
	}

	msg, err := strconv.Unquote(lit.Value)
	if err != nil {
		return
	}

	activeRules := rules.All
	if !cfg.EnableSensitive {
		activeRules = rules.AllWithoutSensitive
	}

	for _, rule := range activeRules {
		rule(pass, msg, lit, call, msgIndex)
	}
}
