package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"identifyWordInComment/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
