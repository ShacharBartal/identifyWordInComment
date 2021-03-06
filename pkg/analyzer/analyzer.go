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
	Name:     "identifyWordAtComment",
	Doc:      "Checks if customs words are contained in comments.",
	Run:      Run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func Run(pass *analysis.Pass) (interface{}, error) {
	inspectorI := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CommentGroup)(nil),
		(*ast.File)(nil),
	}

	inspectorI.Preorder(nodeFilter, func(node ast.Node) {
		switch node.(type) {
		case *ast.CommentGroup:
			checkForKeyWord(node.(*ast.CommentGroup), pass)
		case *ast.File:
			for _, fComments := range node.(*ast.File).Comments {
				checkForKeyWord(fComments, pass)
			}
		}
	})

	return nil, nil
}

func checkForKeyWord(cg *ast.CommentGroup, pass *analysis.Pass, keyWords ...string) {
	if len(keyWords) == 0 {
		keyWords = defaultKeywords
	}
	for _, c := range cg.List {
		for _, keyWord := range keyWords {

			if strings.Contains(c.Text, keyWord) {
				pass.Reportf(cg.Pos(), "word `%s` is contained", keyWord)
			}
		}
	}
}
