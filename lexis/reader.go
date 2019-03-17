package lexis

import "fmt"

type Token struct {
	Class string
	Value string
}

func readWhile(input stream, predicate func(lexeme string) bool) (lexeme string) {
	for !input.eof() && predicate(input.peek()) {
		lexeme += input.next()
	}
	fmt.Println("LEXEME", lexeme)
	return lexeme
}

func readNumber(input stream) *Token {
	var hasDot = false
	number := readWhile(input,
		func(lexeme string) bool {
			if lexeme == "." {
				if hasDot {
					return false
				}
				hasDot = true
				return true
			}

			return isDigit(lexeme)
		},
	)

	return &Token{
		Class: "num",
		Value: number,
	}
}

func readIdentifier(input stream, keywords []string) *Token {
	id := readWhile(input, isIdentifier)

	var class = "var"
	if isKeyword(keywords, id) {
		class = "kw"
	}

	return &Token{
		Class: class,
		Value: id,
	}
}

func readEscaped(input stream, end string) string {
	var (
		escaped = false
		lexeme  string
	)

	input.next()
	for !input.eof() {
		ch := input.next()
		if escaped {
			lexeme += ch
			escaped = false
		} else if ch == "\\" {
			escaped = true
		} else if ch == end {
			break
		} else {
			lexeme += ch
		}
	}

	return lexeme
}

func readString(input stream) *Token {
	return &Token{
		Class: "str",
		Value: readEscaped(input, `"`),
	}
}

func skipComment(input stream) {
	readWhile(input,
		func(lexeme string) bool {
			return lexeme != "\n"
		},
	)

	input.next()
}

func readNext(input stream, keywords []string) (token *Token) {
	for {
		readWhile(input, isWhitespace)
		if input.eof() {
			return nil
		}

		ch := input.peek()
		fmt.Println(ch)

		switch ch {
		case `#`:
			skipComment(input)
			continue
		case `"`:
			return readString(input)
		}

		if isDigit(ch) {
			return readNumber(input)
		}

		if isIdentifierStart(ch) {
			return readIdentifier(input, keywords)
		}

		if isPunctuation(ch) {
			return &Token{
				Class: "punc",
				Value: input.next(),
			}
		}

		if isOperator(ch) {
			return &Token{
				Class: "op",
				Value: readWhile(input, isOperator),
			}
		}

		input.croak(fmt.Sprintf("Can't handle character: %s", ch))
		return nil
	}
}
