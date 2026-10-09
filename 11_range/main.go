package main

import (
	"fmt"
)

// range keyword --> iterating over data structure
func main() {
	nums := []int{6, 7, 8}

	for i, num := range nums {
		fmt.Println(i, num)
	}

	mpp := map[string]string{
		"fname": "achinta",
		"lname": "siddhanta",
	}

	for key, val := range mpp {
		fmt.Println(key, val)
	}

	// unicode code point rune
	// starting byte of rune
	// 255 -> 1 byte
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
