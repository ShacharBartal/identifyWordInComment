package main

import (
	"github.com/ShacharBartal/identifyWordInComment/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
