package main

import (
	"fmt"
	"sync"
)

type Post struct {
	Views int
	mu    sync.Mutex
}

func (P *Post) incPost(wg *sync.WaitGroup) {
	defer func() {
		P.mu.Unlock()
		wg.Done()
	}()

	P.mu.Lock()
	P.Views++

}

func main() {
	var wg sync.WaitGroup

	myPost := Post{
		Views: 0,
	}

	for range 100 {
		wg.Add(1)
		go myPost.incPost(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.Views)
}
