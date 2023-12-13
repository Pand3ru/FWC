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
	var returnStack []Token
	lastWasOperator := false
	evenParentheses := 0
	value := ""

	for i := 0; i < len(input); i++ {
		lastWasComma := false
		char := string(input[i])

		switch char {
		case "+", "-", "*", "/", "%", "^":

			if char == "-" {
				if i+1 < len(input) {
					nextChar := string(input[i+1])
					if i != 0 && nextChar != "(" {
						prevChar := string(input[i-1])
						if isDigit(nextChar) && !isDigit(prevChar) && prevChar != ")" {
							value += "-"
							continue
						}
					} else if i == 0 && nextChar != "(" {
						value += "-"
						continue
					}
				}
			}
			if lastWasOperator {
				fmt.Println("Syntax error! Multiple operators not allowed")
				return nil
			}
			lastWasOperator = true
			token := Token{
				Type:  0, // Type 0 for Operator
				Value: char,
			}
			returnStack = append(returnStack, token)
		case "(", ")":
			evenParentheses++
			token := Token{
				Type:  2, // Type 2 for parentheses
				Value: char,
			}
			if char == ")" {
				lastWasOperator = false
			}
			returnStack = append(returnStack, token)
		case " ":
			continue
		case ",", ".":
			if lastWasComma {
				fmt.Println("Syntax error! Multiple decimal points not allowed")
				return nil
			}
			lastWasComma = true
		default:
			value += char
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
			if string(value[0]) == "-" {
				i--
			}
			token := Token{
				Type:  1, // Type 1 for meth
				Value: value,
			}
			lastWasOperator = false
			value = ""
			returnStack = append(returnStack, token)
		}
	}
	k := len(returnStack) - 1
	v := returnStack[k]
	if v.Type == 0 || evenParentheses%2 != 0 {
		fmt.Println("Syntax Error!")
		return nil
	}
	return returnStack
}
