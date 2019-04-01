package syntax

import (
	"fmt"

	"github.com/enabokov/language/lexis"
)

var priorities = map[string]int{
	`=`:  1,
	`||`: 2,
	`&&`: 3,
	`<`:  7, `>`: 7, `<=`: 7, `>=`: 7, `==`: 7, `!=`: 7,
	`+`: 10, `-`: 10,
	`*`: 20, `/`: 20, `%`: 20,
}

func isPackage(input lexis.TokenStream, token *lexis.Token) bool {
	if token.Class == lexis.ClassKeyword && token.Value == "package" {
		fmt.Println("check if package")
		nextToken := input.Peek()
		if nextToken.Class == lexis.ClassVariable {
			return true
		}
	}

	return false
}

func isImport(input lexis.TokenStream, token *lexis.Token) bool {
	if token.Class == lexis.ClassKeyword && token.Value == "import" {
		fmt.Println("check if import --", token.Value)
		nextToken := input.Peek()
		if nextToken.Class == lexis.ClassString {
			return true
		}
	}

	return false
}

func isFunction(input lexis.TokenStream, token *lexis.Token) bool {
	if token.Class == lexis.ClassKeyword && token.Value == "def" {
		fmt.Println("check if function --", token.Value)
		nextToken := input.Peek()
		if nextToken.Class == lexis.ClassVariable {
			return true
		}
	}

	return false
}

func isVariable(input lexis.TokenStream, token *lexis.Token) bool {
	if token.Class == lexis.ClassKeyword && token.Value == "var" {
		fmt.Println("check if variable --", token.Value)
		nextToken := input.Peek()
		if nextToken.Class == lexis.ClassVariable {
			return true
		}
	}

	return false
}

func isCaller(input lexis.TokenStream, token *lexis.Token) bool {
	if token.Class == lexis.ClassVariable && input.Peek().Class == lexis.ClassCall {
		fmt.Println("check if caller --", token.Value)
		nextToken := input.Peek()
		if nextToken.Class == lexis.ClassCall {
			return true
		}
	}

	return false
}

func isAssignment(input lexis.TokenStream, token *lexis.Token) bool {
	if token.Class == lexis.ClassVariable && input.Peek().Class == lexis.ClassOperator && input.Peek().Value == `=` {
		fmt.Println("check if assignment --", token.Value)
		return true
	}

	return false
}
