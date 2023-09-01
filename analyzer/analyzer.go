package analyzer

import (
	"bytes"
	"flag"
	"go/ast"
	"go/printer"
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

            var buf bytes.Buffer
            if err := printer.Fprint(&buf, pass.Fset, be); err != nil {
                panic(err)
            }
            nodeStr := buf.String()
			if !strings.HasPrefix(nodeStr, "panic(") {

				return true
			}

			pass.Reportf(be.Pos(), "Illegal panic invocation: %q", nodeStr)
			return true
		})
	}

	return nil, nil
}
