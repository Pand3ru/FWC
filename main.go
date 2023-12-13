package main

import "fmt"

func main() {
	something := tokenizer_ParseString("4.32 ++ (1*2,223)")

	for _, v := range something {
		fmt.Printf("Type: %v | Value: %v \n", v.Type, v.Value)
	}
}
