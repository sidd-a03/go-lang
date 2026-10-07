package main

import "fmt"

func main() {
	const name string = "achinta"

	const (
		port int = 5000
		host string = "localhost"
	)

	fmt.Println(port, host)
}