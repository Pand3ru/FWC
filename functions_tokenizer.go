package main

import (
	"fmt"
	"strconv"
)

// import

type Token struct {
	Type  int32 // 0: Operator (+,-) 1: Operand (1,2) 2: ()
	Value string
}

func isDigit(in string) bool {
	_, err := strconv.Atoi(in)
	return err == nil
}

func tokenizer_ParseString(input string) []Token {
	var ReturnStack []Token
	lastWasOperator := false

	for i := 0; i < len(input); i++ {
		lastWasComma := false
		char := string(input[i])

		switch char {
		case "+", "-", "*", "/", "mod", "%", "^":
			if lastWasOperator {
				fmt.Println("Syntax error! Multiple operators not allowed")
				return nil
			}
			lastWasOperator = true
			token := Token{
				Type:  0, // Type 0 for Operator
				Value: char,
			}
			ReturnStack = append(ReturnStack, token)
		case "(", ")":
			token := Token{
				Type:  2, // Type 2 for parentheses
				Value: char,
			}
			ReturnStack = append(ReturnStack, token)
		case " ":
			continue
		case ",", ".":
			if lastWasComma {
				fmt.Println("Syntax error! Multiple decimal points not allowed")
				return nil
			}
			lastWasComma = true
		default:
			value := char
			j := i + 1
			for ; j < len(input); j++ {
				nextChar := string(input[j])
				if !isDigit(nextChar) && nextChar != "," && nextChar != "." {
					break
				}
				if nextChar == "," || nextChar == "." {
					if lastWasComma {
						fmt.Println("Syntax error! Multiple decimal points not allowed")
						return nil
					}
					lastWasComma = true
				}
				value += nextChar
			}
			i = i + len(value) - 1
			token := Token{
				Type:  1, // Type 1 for meth
				Value: value,
			}
			ReturnStack = append(ReturnStack, token)
		}
	}

	return ReturnStack
}
