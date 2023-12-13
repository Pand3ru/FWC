package main

import "fmt"

func main() {
	something := tokenizer_ParseString("-(45*23)+76/32 % 3+-2")

	for _, v := range something {
		fmt.Printf("Type: %v | Value: %v \n", v.Type, v.Value)
	}

	in := functions_infix(something)

	for i := 0; i < len(in); i++ {
		fmt.Printf("%v ", in[i])
	}
	fmt.Println()
}
