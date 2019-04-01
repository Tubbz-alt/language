package main

import (
	"fmt"
	"os"

	"github.com/enabokov/language/lexis"
	"github.com/enabokov/language/semantics"
	"github.com/enabokov/language/syntax"
)

func main() {
	filename := os.Args[1]
	tokenStream := lexis.Analyze(filename)
	ast := syntax.Analyze(tokenStream)
	fmt.Println(semantics.Analyze(ast))
}
