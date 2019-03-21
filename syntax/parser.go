package syntax

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/enabokov/language/lexis"
)

type bigToken struct {
	class    string
	value    string
	operator string
	left     *bigToken
	right    *bigToken
	function *lexis.Token
	args     []bigToken
	cond     string
	_else    string
	prog     []*bigToken
}

func maybeBinary(input lexis.TokenStream, left *bigToken, prec int) *bigToken {
	token := IsOperator(input, "")
	fmt.Println(`SDASDAS`, token)
	fmt.Println(input.Peek())
	if token != nil {
		currentPrec := Precedence[token.Value]
		if currentPrec > prec {
			input.Next()

			class := "binary"
			if token.Value == "=" || token.Value == ":=" {
				class = "assign"
			}

			parsed := parseAtom(input)
			fmt.Println("==============Patsed:", parsed)
			right := maybeBinary(input, parsed, currentPrec)
			return maybeBinary(
				input,
				&bigToken{
					class:    class,
					operator: token.Value,
					left:     left,
					right:    right,
				},
				prec,
			)
		}
	}

	return left
}

func delimited(input lexis.TokenStream, start string, stop string, separator string, parser func(input lexis.TokenStream) *bigToken) (tokens []bigToken) {
	var first = true

	SkipPunctuation(input, start)
	for !input.Eof() {
		if IsPunctuation(input, stop) != nil {
			break
		}

		if first {
			first = false
		} else {
			SkipPunctuation(input, separator)
		}

		if IsPunctuation(input, stop) != nil {
			break
		}

		parsed := parser(input)
		tokens = append(tokens, *parsed)
	}
	SkipPunctuation(input, stop)
	return tokens
}

func parseCall(input lexis.TokenStream, token *lexis.Token) *bigToken {
	return &bigToken{
		class:    "call",
		function: token,
		args:     delimited(input, "(", ")", ",", parseExpression),
	}
}

func parseVarname(input lexis.TokenStream) *bigToken {
	name := input.Next()
	if name.Class != lexis.ClassVariable {
		log.Fatalln(input.Croak(fmt.Sprintf("Expecting variable name, but not %s", name)))
	}

	return &bigToken{
		class: "var",
		value: name.Value,
	}
}

func maybeCall(input lexis.TokenStream, token *lexis.Token) *bigToken {
	if IsPunctuation(input, "(") != nil {
		return parseCall(input, token)
	}

	return &bigToken{class: "call", function: token, args: nil}
}

func parseExpression(input lexis.TokenStream) *bigToken {
	parsed := parseAtom(input)
	token := maybeBinary(input, parsed, 0)
	fmt.Println(token)

	var class string
	var value string

	if token != nil {
		class = token.class
		value = token.operator
	}

	return maybeCall(input, &lexis.Token{Class: class, Value: value})
}

func parseBool(input lexis.TokenStream) *bigToken {
	return &bigToken{
		class: lexis.ClassBool,
		value: fmt.Sprint(input.Next().Value == "true"),
	}
}

func parseIf(input lexis.TokenStream) *bigToken {
	SkipKeyword(input, "if")
	cond := parseExpression(input)
	if IsPunctuation(input, "{") == nil {
		SkipKeyword(input, "then")
	}

	// then := parseExpression(input)

	var ret *bigToken
	if IsKeyword(input, "else") != nil {
		input.Next()
		ret = &bigToken{
			class: "if",
			cond:  cond.value,
			_else: parseExpression(input).value,
		}
	}

	return ret
}

func parseDefer(input lexis.TokenStream) *bigToken {
	SkipKeyword(input, "defer")
	variable := parseExpression(input)
	input.Next()

	delimited(input, "(", ")", ",", parseVarname)
	parseExpression(input)

	return &bigToken{
		class: "defer",
		value: variable.value,
	}
}

func parseAtom(input lexis.TokenStream) *bigToken {
	if IsPunctuation(input, "(") != nil {
		input.Next()
		exp := parseExpression(input)
		SkipPunctuation(input, ")")
		return exp
	}

	if IsPunctuation(input, "{") != nil {
		return parseProgram(input)
	}

	if IsKeyword(input, "if") != nil {
		return parseIf(input)
	}

	if IsKeyword(input, "true") != nil || IsKeyword(input, "false") != nil {
		return parseBool(input)
	}

	if IsKeyword(input, "defer") != nil {
		return parseDefer(input)
	}

	// if IsKeyword(input, "lambda") != nil {
	// input.Next()
	// return parseLambda()
	// }

	token := input.Next()
	if token.Class == lexis.ClassVariable || token.Class == lexis.ClassNumber || token.Class == lexis.ClassString {
		return &bigToken{class: "token", function: token, args: nil}
	}

	Unexpected(input)

	return maybeCall(input, token)
}

func parseProgram(input lexis.TokenStream) *bigToken {
	program := delimited(input, "{", "}", ":", parseExpression)
	if len(program) == 0 {
		return &bigToken{class: lexis.ClassBool, value: "false"}
	}

	out, err := json.Marshal(program[0])
	if err != nil {
		panic(err)
	}

	if len(program) == 1 {
		return &bigToken{class: lexis.ClassProgram, value: string(out)}
	}

	return nil
}

func ParseTopLevel(input lexis.TokenStream) *bigToken {
	var prog []*bigToken
	for !input.Eof() {
		prog = append(prog, parseExpression(input))
		if !input.Eof() {
			fmt.Println("Next line, ", prog)
			// SkipPunctuation(input, ";")
		}
	}

	return &bigToken{class: "prog", prog: prog}
}
