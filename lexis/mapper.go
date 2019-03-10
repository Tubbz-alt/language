package lexis

import (
	"fmt"
	"strings"

	"github.com/enabokov/language/bnf"
)

type token struct {
	class string
	value string
}

func getLexemes(sourceCode string) (lexemes []string) {
	return strings.Fields(sourceCode)
}

func getTokens(lexemes []string, bnfConfig bnf.BNF) []token {
	var tokens []token
	for _, lexeme := range lexemes {
		if checkValueInArray(lexeme, bnfConfig.Keywords) {
			fmt.Println("Keyword:", lexeme)
		} else if checkValueInArray(lexeme, bnfConfig.PossibleType) {
			fmt.Println("Type:", lexeme)
		} else if checkValueInArray(lexeme, bnfConfig.Punctuation) {
			fmt.Println("Punctuation:", lexeme)
		} else if checkValueIsString(lexeme) {
			fmt.Println("String:", lexeme)
		} else if checkValueIsNumber(lexeme) {
			fmt.Println("Number:", lexeme)
		} else {
			fmt.Println("Unknown:", lexeme)
		}
	}

	return tokens
}
