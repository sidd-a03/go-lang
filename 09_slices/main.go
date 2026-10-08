package main

import "fmt"

// slice --> dynamic array
// most used construct in golang
// + usefull methods

func main() {
	// un-initialized slice is nil
	var nums []int

	fmt.Println(nums == nil) // true
	fmt.Println(len(nums)) // returns 0

	var arr = make([]int, 2, 3) // initialize arr with [0 0]
	fmt.Println(arr, len(arr), arr == nil)
	fmt.Println(cap(arr)) // capacity -> max numbers of element can fit

	arr = append(arr, 1,2,3)
	fmt.Println(arr, cap(arr))
}