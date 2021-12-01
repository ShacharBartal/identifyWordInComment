package analyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"strings"
)

var (
	defaultKeywords = []string{"TODO", "todo", "nolint", "dominos"}
)

var Analyzer = &analysis.Analyzer{
	Name:             "identifyWordAtComment",
	Doc:              "Checks if customs words are contained in comments.",
	Run:              run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error){
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	
	nodeFilter := []ast.Node{
		(*ast.CommentGroup)(nil),
		(*ast.File)(nil),
	}
	
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		switch node.(type) {
		case *ast.CommentGroup:
			for _, co := range node.(*ast.CommentGroup).List {
				for _, keyword := range defaultKeywords {
					if strings.Contains(co.Text, keyword) {
						pass.Reportf(node.Pos(), "word `%s` is contained", keyword)
					}
				}
			}
		case *ast.File:
			for _, cos := range node.(*ast.File).Comments {
				for _, c := range cos.List {
					for _, keyword := range defaultKeywords {
						if strings.Contains(c.Text, keyword) {
							pass.Reportf(c.Pos(), "word `%s` is contained", keyword)
						}
					}
				}
			}
		}
	})
	
	return nil, nil
}