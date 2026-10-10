package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func anything(param ...any) any {
	return param
}

func main() {
	fmt.Println(1, 2, 3, 4, "hello")
	nums := []int{3, 4, 5, 6}
	result := sum(nums...)

	fmt.Println(result)
}
