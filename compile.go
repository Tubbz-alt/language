package main

import (
	"fmt"
	"os"

	"github.com/enabokov/language/lexis"
)

func main() {
	filename := os.Args[1]
	lexemes := lexis.Analyze(filename)
	fmt.Println(lexemes)
}
