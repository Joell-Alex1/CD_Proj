package models

type TokenType string

const (
	Keyword     TokenType = "Keyword"
	Operator    TokenType = "Operator"
	Punctuation TokenType = "Punctuation"
	Identifier  TokenType = "Identifier"
	Number      TokenType = "Number"
)

type Token struct {
	Lexeme string
	Type   TokenType
}
