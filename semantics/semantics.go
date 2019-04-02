package semantics

import (
	"github.com/enabokov/language/syntax"
)

func Analyze(ast syntax.TokenProgram) error {
	scan(ast)
	// fmt.Printf("%# v", pretty.Formatter(ast))
	return nil
}
