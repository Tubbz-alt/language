package syntax

import (
	"fmt"

	"github.com/enabokov/language/lexis"
)

func Analyze(input lexis.TokenStream) bool {
	result := ParseTopLevel(input)
	fmt.Println("RESULTS")
	fmt.Println(*result.prog[0].right)

	return true
}
