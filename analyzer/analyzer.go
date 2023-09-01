package analyzer

import (
	"bytes"
	"flag"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

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
			be, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			nodeStr := render(pass.Fset, be)
			if !strings.HasPrefix(nodeStr, "panic(") {

				return true
			}

			pass.Reportf(be.Pos(), "Illegal panic invocation: %q", nodeStr)
			return true
		})
	}

	return nil, nil
}

// render returns the pretty-print of the given node
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
