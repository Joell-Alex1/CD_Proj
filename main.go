package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cd_proj/lexer"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter code: ")
	code, _ := reader.ReadString('\n')
	code = strings.TrimSpace(code) // remove newline

	tokens := lexer.LexicalAnalysis(code)

	fmt.Printf("%-10s | %-12s\n", "Lexeme", "Type")
	fmt.Println("---------------------------")
	for _, tok := range tokens {
		fmt.Printf("%-10s | %-12s\n", tok.Lexeme, tok.Type)
	}
}
