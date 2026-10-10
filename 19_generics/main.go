package main

import "fmt"

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type Stack[T int | string] struct {
	elements []T
}

func main() {
	lang := []string{"golang", "typescript"}
	printSlice(lang)

	myStack := Stack[int]{
		elements: []int{1, 2, 3},
	}

	fmt.Println(myStack)
}
