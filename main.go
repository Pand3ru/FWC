package main

import "fmt"

func main() {
	something := tokenizer_ParseString("-x^2+x+1")
	in := functions_infix(something)

	for _, v := range something {
		fmt.Printf("Type: %v | Value: %v \n", v.Type, v.Value)
	}


	for i := 0; i < len(in); i++ {
		if token, ok := in[i].(Token); ok {
			fmt.Printf("%v ", token.Value)
		} else {
			// Handle the case where the type assertion fails
			fmt.Println("Type assertion failed")
		}
	}
	fmt.Println()
//	sol, err := CalculatePostfix(&in)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println(sol)

}
