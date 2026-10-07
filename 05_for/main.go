package main

import "fmt"

// no while loop in go lang
// for -> only construct in go for looping

func main() {
	// while loop using foor
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infinite loop

	/*
		for {
			println("1")
		}		
	*/

	// classic for loop

	for j := 0; j <= 3; j++ {
		if j == 2 {
			continue
		}
		fmt.Println(j)
	}

	// range -> exclude 3
	for j := range 3 {
		fmt.Println(j)
	}
}