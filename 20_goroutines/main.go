package main

import (
	"fmt"
	"time"
)


func task(id int)  {
	fmt.Println("Doing task", id)
}


func main() {
	for i := range 11 {
		go task(i)
	}

	time.Sleep(time.Second * 2)
}