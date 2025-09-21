package lexer

import (
	"cd_proj/models"
	"regexp"
)

// Define sets like Python
var Operators = []string{"+", "-", "/", "*", "%", "="}
var Punctuation = []string{";", ",", "{", "}", "(", ")", "[", "]"}
var Keywords = []string{
	"False", "await", "else", "import", "pass",
	"None", "break", "except", "in", "raise",
	"True", "class", "finally", "is", "return",
	"and", "continue", "for", "lambda", "try",
	"as", "def", "from", "nonlocal", "while",
	"assert", "del", "global", "not", "with",
	"async", "elif", "if", "or", "yield", "print", "printf",
}

// helper function
func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func LexicalAnalysis(code string) []models.Token {
	// Regex will split into words, numbers, and single chars like ()[]{};
	re := regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*|\d+|[=+\-*/%{}()\[\];,]`)
	matches := re.FindAllString(code, -1)

	var tokens []models.Token
	for _, m := range matches {
		switch {
		case contains(Keywords, m):
			tokens = append(tokens, models.Token{Lexeme: m, Type: models.Keyword})
		case contains(Operators, m):
			tokens = append(tokens, models.Token{Lexeme: m, Type: models.Operator})
		case contains(Punctuation, m):
			tokens = append(tokens, models.Token{Lexeme: m, Type: models.Punctuation})
		case regexp.MustCompile(`^\d+$`).MatchString(m):
			tokens = append(tokens, models.Token{Lexeme: m, Type: models.Number})
		default:
			tokens = append(tokens, models.Token{Lexeme: m, Type: models.Identifier})
		}
	}

	return tokens
}
