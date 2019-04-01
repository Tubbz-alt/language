package main

import (
	"fmt"
	"os"

	"github.com/enabokov/language/lexis"
	"github.com/enabokov/language/syntax"
)

func main() {
	filename := os.Args[1]
	os.Setenv("BNF_FILE_PATH", "/Users/eduardnabokov/Documents/work/cloudmade/bnf/bnf.yml")
	tokenStream := lexis.Analyze(filename)
	fmt.Println(syntax.Analyze(tokenStream))
}
