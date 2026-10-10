package main

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)

func processNum(numChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	num := <-numChan
	fmt.Println("Processing number", num)
	time.Sleep(time.Second)
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("process....")
}

func emailSender(emailChan <-chan string, done chan <-bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	// emailChannel := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChannel, done)

	// for email := range 100 {
	// 	emailChannel <- fmt.Sprintf("%d@gmail.com", email)
	// }
	// close(emailChannel)

	// // fmt.Println(<-emailChannel)
	// // fmt.Println(<-emailChannel)

	// fmt.Println("done sending..")
	// <-done

	// un-buffered

	// done := make(chan bool)
	// go task(done)

	// <- done // blocking

	// var wg sync.WaitGroup
	// // messageChannel := make(chan string)

	// // messageChannel <- "ping" // channel -> blocking

	// // msg := <- messageChannel

	// numChan := make(chan int)
	// wg.Add(1)
	// go processNum(numChan, &wg)

	// // for {
	// // 	numChan <- rand.Intn(100)
	// // }

	// numChan <- rand.Intn(100)

	// wg.Wait()
}
