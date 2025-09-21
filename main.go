package main

import (
	"fmt"
	"cd_proj/lexer"
)

func main() {
	code := "print ( [hello] ) printf 123 + 45"
	tokens := lexer.LexicalAnalysis(code)

	fmt.Printf("%-10s | %-12s\n", "Lexeme", "Type")
	fmt.Println("---------------------------")
	for _, tok := range tokens {
		fmt.Printf("%-10s | %-12s\n", tok.Lexeme, tok.Type)
	}
}
