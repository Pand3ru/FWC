package main

import (
	"fmt"
	"strconv"
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
			Operator := tokens[i].Value
			if len(OperatorStack) > 0 {
				for j := len(OperatorStack) - 1; j >= 0; j-- {
					if GetPrecedence(Operator) < GetPrecedence(OperatorStack[j]) {
						v, err := Pop(&OperatorStack)
						if err != nil {
							fmt.Println("If that error happens then the Tokenizer is broken. #1")
						}
						Push(&OutputStack, v)
					}
				}
			}
			Push(&OperatorStack, Operator)
		case 1:
			value, err := strconv.ParseFloat(TokenValue, 64) // Convert the string to a float64
			if err != nil {
				fmt.Println("If that error happens then the Tokenizer is broken. #2")
				return nil
			}
			Push(&OutputStack, value)
		case 2:
			TokenValue := tokens[i].Value
			if TokenValue == "(" {
				Push(&OperatorStack, TokenValue)
			} else if TokenValue == ")" {
				if len(OperatorStack) > 0 {
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
