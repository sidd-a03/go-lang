package main

import (
	"fmt"
	"slices"
)

// slice --> dynamic array
// most used construct in golang
// + usefull methods

func main() {
	// un-initialized slice is nil
	// var nums []int
	nums := []int {}

	fmt.Println(nums == nil) // true
	fmt.Println(len(nums)) // returns 0

	var arr = make([]int, 2, 3) // initialize arr with [0 0]
	fmt.Println(arr, len(arr), arr == nil)
	fmt.Println(cap(arr)) // capacity -> max numbers of element can fit

	arr = append(arr, 1,2,3)
	fmt.Println(arr, cap(arr))

	// copy slice
	var arr2 = make([]int, len(arr))

	arr2 = append(arr2, 2)
	copy(arr2, arr)
	fmt.Println(arr, arr2)

	// slice operator
	nums2 := []int {1,2,3}
	fmt.Println(nums2[:1])

	// slice package
	nums3 := []int {1,2}
	nums4 := []int {1,2}

	fmt.Println(slices.Equal(nums3, nums4))
}