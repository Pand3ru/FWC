package main

import "fmt"

func main() {
	something := tokenizer_ParseString("-(45*23)+76/32 % 3+-2")

	for _, v := range something {
		fmt.Printf("Type: %v | Value: %v \n", v.Type, v.Value)
	}

	in := functions_infix(something)

	for i := 0; i < len(in); i++ {
		if token, ok := in[i].(Token); ok {
			fmt.Printf("%v ", token.Value)
		} else {
			// Handle the case where the type assertion fails
			fmt.Println("Type assertion failed")
		}
	}
	fmt.Println()

	fmt.Println(CalculatePostfix(&in))

}
