package lexis

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/enabokov/language/bnf"
)

func readFile(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	return string(file)
}

func Analyze(filename string) []token {
	bnfConfig := bnf.Read()

	sourceCode := readFile(filename)
	lexemes := getLexemes(sourceCode)
	tokens := getTokens(lexemes, bnfConfig)
	return tokens
}
