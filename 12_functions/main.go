package main

import "fmt"

func add(a, b int) int {
	return a+b
}

func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

func processIt() func() int {
	return func() int {
		return 200
	}
}

func main() {
	sum := add(2,3)
	fmt.Println(sum)

	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)

	fn := func () int {
		return 2
	}

	fmt.Println(fn())

	fn2 := processIt()

	fmt.Println(fn2())
}