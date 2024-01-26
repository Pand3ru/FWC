package main

import "errors"

type Stack []interface{}

func Push(s *Stack, value interface{}) {
	*s = append(*s, value)
}

func Pop(s *Stack) (value interface{}, err error) {
	if len(*s) == 0 {
		return nil, errors.New("Stack is empty")
	}
	index := len(*s) - 1
	value = (*s)[index]
	*s = (*s)[:index]
	return value, nil
}

func Peek(s *Stack) interface{} {
	if len(*s) == 0 {
		return nil
	}
	index := len(*s) - 1
	value := (*s)[index]
	return value
}
