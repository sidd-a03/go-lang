package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	role := "admin"
	hasPermission := true

	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}

	if Age := 18; Age >= 18 {
		fmt.Println("Adult", Age)
	}

	// go doesn't has ternary operator
}