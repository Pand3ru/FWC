package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func CalculatePostfix(s *Stack) (float64, error) {
	var OutputStack Stack
	slice := *s

	for i := 0; i < len(slice); i++ {
		TokenType := slice[i].(Token).Type
		switch TokenType {
		case 1: // Assuming 1 represents operand
			Push(&OutputStack, slice[i])
		case 0: // Assuming 0 represents operator
			Operator := slice[i].(Token).Value
			Operand2, err := Pop(&OutputStack)
			if err != nil {
				return 0, err
			}
			Operand1, err := Pop(&OutputStack)
			if err != nil {
				return 0, err
			}
			val, err := Calculate(Operator, Operand1.(Token), Operand2.(Token))
			if err != nil {
				return 0, err
			}
			//fmt.Printf("\n=====\nTokentype: %v, TokenValue: %v, Operator: %v, Operand1: %v, Operand2: %v, OutputValue: %v \n=====\n", TokenType, slice[i].(Token).Value, Operator, Operand1, Operand2, val)
			Push(&OutputStack, Token{Value: fmt.Sprintf("%f", val)})
		}
	}

	result, err := Pop(&OutputStack)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(result.(Token).Value, 64)
}

func Calculate(a string, b Token, c Token) (float64, error) {
	op := a
	A, err := strconv.ParseFloat(b.Value, 64)
	if err != nil {
		return 0, err
	}

	B, err := strconv.ParseFloat(c.Value, 64)
	if err != nil {
		return 0, err
	}

	switch op {
	case "+":
		return A + B, nil
	case "-":
		return A - B, nil
	case "*":
		return A * B, nil
	case "^":
		return math.Pow(A, B), nil
	case "/":
		if B != 0 {
			return A / B, nil
		}
		return 0, errors.New("divide by zero")
	case "%":
		Aint := int(A)

		Bint := int(B)

		if Bint != 0 {
			return float64(Aint % Bint), nil
		}
		return 0, errors.New("divide by zero")
	default:
		return 0, errors.New("invalid operator")
	}
}

func ReplaceAndTokenize(input string, value int) Stack {
	input = strings.Replace(input, "x", fmt.Sprint(value), -1)
	tokens := tokenizer_ParseString(input)
	return functions_infix(tokens)
}
