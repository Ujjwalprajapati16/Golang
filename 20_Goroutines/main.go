package main

import (
	"fmt"
	"sync"
)

// goroutine -> light weight thread
// go keyword is used to create a goroutine

func task(id int) {
	fmt.Println("Task", id, "started")
}

// wait group -> sync goroutines
// sync.WaitGroup
// Add() -> add number of goroutines
// Done() -> decrement the count
// Wait() -> wait for all goroutines to finish
// defer -> used to call Done() after the function finishes

func taskWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", id, "started")
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	// go task(i)

	// 	go func(i int) {
	// 		fmt.Println("Task", i, "started")
	// 	}(i)
	// }
	// fmt.Println("All tasks started")

	// wait for goroutines to finish
	// time.Sleep(2 * time.Second)

	// using wait group
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go taskWithWaitGroup(i, &wg)
	}

	wg.Wait()
}
