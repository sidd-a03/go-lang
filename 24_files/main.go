package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	f, err := os.Open("example.txt")

	if err != nil {
		// log the error
		panic(err)
	}
	
	defer f.Close()
	
	fileInfo, err := f.Stat()

	if err != nil {
		// log the error
		panic(err)
	}

	fmt.Println("file name", fileInfo.Name())


	// read file
	buf := make([]byte, 50)

	_, err = f.Read(buf)

	if err != nil {
		panic(err)
	}

	for i := range len(buf) {
		fmt.Println("data is", string(buf[i]))
	} */

	/*
	f, err := os.ReadFile("example.txt") // read the content all at once in the memory

	if err != nil {
		panic(err)
	}

	fmt.Println(string(f)) */

	dir, err := os.Open(".")

	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1)

	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}

}
