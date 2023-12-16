package main

import "fmt"

func main() {
	input := "187-162"
	// tokens := tokenizer_ParseString(input)
	tokens := ReplaceAndTokenize(input, 3)
	value, _ := CalculatePostfix(&tokens)
	fmt.Println(value)
}
