package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cd_proj/lexer"
	"cd_proj/parser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter code: ")
	code, _ := reader.ReadString('\n')
	code = strings.TrimSpace(code)

	// Lexical analysis
	tokens := lexer.LexicalAnalysis(code)
	fmt.Printf("\n%-10s | %-12s\n", "Lexeme", "Type")
	fmt.Println("---------------------------")
	for _, tok := range tokens {
		fmt.Printf("%-10s | %-12s\n", tok.Lexeme, tok.Type)
	}

	// Example grammar: S -> A + B, A -> a, B -> b
	g := parser.Grammar{
		NonTerminals: []string{"S", "A", "B"},
		Terminals:    []string{"a", "b", "+"},
		StartSymbol:  "S",
		Productions: map[string][][]string{
			"S": {{"A", "+", "B"}},
			"A": {{"a"}},
			"B": {{"b"}},
		},
	}

	// Map tokens to terminals expected by grammar
	var tokenSymbols []string
	for _, tok := range tokens {
		switch tok.Lexeme {
		case "a", "b", "+": // terminals in grammar
			tokenSymbols = append(tokenSymbols, tok.Lexeme)
		default:
			tokenSymbols = append(tokenSymbols, tok.Lexeme)
		}
	}

	// Compute FIRST, FOLLOW, Table
	first := parser.ComputeFirst(g)
	follow := parser.ComputeFollow(g, first)
	table := parser.BuildParsingTable(g, first, follow)

	// Parse
	parser.LL1ParseTable(tokenSymbols, g, table)
}
