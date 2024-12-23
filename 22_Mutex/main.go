package main

import (
	"fmt"
	"sync"
)

// Mutex
// Mutex is a lock that we can set to protect a shared resource.
// It is used to prevent race conditions when multiple goroutines are trying to access the same resource.

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()

	p.views++
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myPost.increment(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}
