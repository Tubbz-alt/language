package lexis

import (
	"regexp"
	"strconv"
)

func isKeyword(keywords []string, lexeme string) bool {
	for _, key := range keywords {
		if key == lexeme {
			return true
		}
	}

	return false
}

func isDigit(lexeme string) bool {
	if _, err := strconv.Atoi(lexeme); err != nil {
		return false
	}

	return true
}

func isIdentifierStart(lexeme string) bool {
	return regexp.MustCompile(`[a-zA-Z_]`).MatchString(lexeme)
}

func isIdentifier(lexeme string) bool {
	return isIdentifierStart(lexeme) || regexp.MustCompile(`[0-9-]`).MatchString(lexeme)
}

func isOperator(lexeme string) bool {
	return regexp.MustCompile(`[\+\-\*\/%=&|<>!]`).MatchString(lexeme)
}

func isPunctuation(lexeme string) bool {
	return regexp.MustCompile(`[,;\(\){}\[\]]`).MatchString(lexeme)
}

func isWhitespace(lexeme string) bool {
	return regexp.MustCompile(`[[:space:]]`).MatchString(lexeme)
}
