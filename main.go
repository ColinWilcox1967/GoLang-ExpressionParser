

package main

import (
	"fmt"
	"./expressionparser"
)

func main() {

	// Example expression: (2 + 3) * 5
	expr := "(2 + 3) * 5"
	lexer := expressionparser.NewLexer(expr)
	parser := expressionparser.NewParser(lexer)

	// Parse the expression
	ast, err := parser.Parse()
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}

	// Evaluate the expression
	result, err := expressionparser.Eval(ast)
	if err != nil {
		fmt.Println("Error evaluating expression:", err)
		return
	}

	fmt.Println("Result:", result)
}

// end of file
