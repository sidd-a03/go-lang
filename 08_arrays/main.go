package main

import "fmt"

func main() {
	// - fixed size, that is predictable
	// - Memory optimization
	// - Constant time access
	// default values 0
	var nums [4]int

	// array length
	fmt.Println(len(nums))

	nums[0] = 1

	fmt.Println(nums)

	var booleans [4]bool
	fmt.Println(booleans) // [false false false false]

	var stringss [2]string
	fmt.Println(stringss)

	arr := [3]int {1,2,3} // short-hand declaration
	fmt.Println(arr)

	d_arr := [2][2]int {{1, 2}, {3, 4}}
	fmt.Println(d_arr)
}