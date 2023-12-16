package main

import (
	"fmt"
	"strings"
)

func functions_infix(tokens []Token) Stack {
	var OperatorStack Stack
	var OutputStack Stack

	// Type 0: Operator
	// Type 1: Zahlen
	// Type 2: Klammern

	for i := 0; i < len(tokens); i++ {
		TokenType := tokens[i].Type
		TokenValue := strings.Replace(tokens[i].Value, ",", ".", -1)

		switch TokenType {
		case 0:
			// Handling right associativity for '^'
			for len(OperatorStack) > 0 && OperatorStack[len(OperatorStack)-1] != "(" {
				topOp := OperatorStack[len(OperatorStack)-1].(Token).Value // Assuming the OperatorStack stores Tokens
				if GetPrecedence(TokenValue) < GetPrecedence(topOp) || (TokenValue != "^" && GetPrecedence(TokenValue) == GetPrecedence(topOp)) {
					v, err := Pop(&OperatorStack)
					if err != nil {
						fmt.Println("Error in pop operation.")
						return nil
					}
					Push(&OutputStack, v)
				} else {
					break
				}
			}
			Push(&OperatorStack, tokens[i])
		case 1, 4:
			Push(&OutputStack, tokens[i])
		case 2:
			if TokenValue == "(" {
				Push(&OperatorStack, TokenValue)
			} else if TokenValue == ")" && len(OperatorStack) > 0 {

				for j := len(OperatorStack) - 1; j >= 0; j-- {
					if OperatorStack[j] != "(" {
						v, err := Pop(&OperatorStack)
						if err != nil {
							fmt.Println("If that error happens then the Tokenizer is broken. #3")
							return nil
						}
						Push(&OutputStack, v)
					} else {
						Pop(&OperatorStack)
						break
					}
				}

			}
		}
	}
	if len(OperatorStack) != 0 {
		for i := len(OperatorStack) - 1; i >= 0; i-- {
			v, err := Pop(&OperatorStack)
			if err != nil {
				fmt.Println("If that error happens then the Tokenizer is broken. #4")
				return nil
			}
			Push(&OutputStack, v)
		}
	}
	return OutputStack
}

func GetPrecedence(val interface{}) int8 {
	switch val {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "%":
		return 3
	case "^":
		return 4
	}
	return -1
}
