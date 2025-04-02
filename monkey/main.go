package main

import (
	"fmt"
	"monkey/lexer"
	"monkey/token"
)

func main() {
	input := "let five = 5;"

	l := lexer.New(input)

	for {
		tok := l.NextToken()
		fmt.Println(tok)
		// fmt.Printf("Type %s, Literal: %s\n", tok.Type, tok.Literal)

		if tok.Type == token.EOF {
			break
		}
	}
}
