package adams

import (
	"flag"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var flagSet flag.FlagSet

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "adams",
		Doc:   "Checks whether any calls to Panic() are made.",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			switch fun := call.Fun.(type) {
			case *ast.Ident:
				if "panic" == fun.Name {
					pass.Reportf(call.Pos(), "Illegal panic")
				}
			}
			return true
		})
	}
	return nil, nil
}
