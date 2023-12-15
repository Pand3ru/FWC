package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	Type  int8 // 0: Operator (+,-) 1: Operand (1,2) 2: () 3: f(x) 4: Variable
	Value string
}

// Tries to convert string into digit
func isDigit(in string) bool {
	_, err := strconv.Atoi(in)
	return err == nil || (in >= "a" && in <= "z")
}

// Function that tokenizes a string.
func tokenizer_ParseString(input string) []Token {
	// Declare values
	var returnStack []Token
	lastWasOperator := false
	evenParentheses := 0
	value := ""
	// Iterate trough entire input
	for i := 0; i < len(input); i++ {

		// More Variables
		lastWasComma := false
		char := string(input[i])

		// Make sure there are no spaces
		input = strings.Replace(input, " ", "", -1)

		switch char {
		case "+", "-", "*", "/", "%", "^":

			if char == "-" {
				nextChar := ""
				if i+1 < len(input) {
					nextChar = string(input[i+1])
				}
				prevChar := ""
				if i > 0 {
					prevChar = string(input[i-1])
				}

				if nextChar == "(" {
					if i == 0 || (prevChar != "" && !isDigit(prevChar) && prevChar != ")") {
						// Handling negative parenthetical expressions (e.g., "-(2 + 3)")
						token1 := Token{Type: 1, Value: "-1"}
						token2 := Token{Type: 0, Value: "*"}
						returnStack = append(returnStack, token1, token2)
						continue
					}
				} else if i == 0 || (prevChar != "" && !isDigit(prevChar) && prevChar != ")") {
					// Negative sign (Unary Minus)
					value += "-"
					continue
				}
				// Else treat as a subtraction operator (Binary Minus)
				token := Token{Type: 0, Value: char}
				returnStack = append(returnStack, token)
				continue
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
		case ",", ".":
			if lastWasComma {
				fmt.Println("Syntax error! Multiple decimal points not allowed")
				return nil
			}
			lastWasComma = true
		default:	
			value += char
			if char >= "a" && char <="z" {
				fmt.Printf("value: %v, char %v\n", char, value)
				token := Token{
					Type: 4,
					Value: value,
				}
				returnStack = append(returnStack, token)
				value = ""
				lastWasOperator = false
				break
			}
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
