package main

import (
	"github.com/malteehrlen/adams"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(adams.NewAnalyzer())
}
