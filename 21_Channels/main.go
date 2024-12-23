package main

import (
	"fmt"
	"time"
)

// channel -> make(chan type)
// channel -> make(chan type, buffer)

// send only channel -> chan<- type
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("process number ", num)
		time.Sleep(1 * time.Second)
	}
}

// receive only channel -> <-chan type
func sum(result chan int, num1 int, num2 int) {
	result <- num1 + num2
}

// goroutine syncronizer
func task(done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println("Processing task......")
}

// difference between gouroutine and channel
// gouroutine -> parallel execution
// channel -> communication between gouroutine
// channel -> unidirectional

func emailSender(emailChannel chan string, done chan bool) {
	defer func() {
		done <- true
		close(emailChannel)
	}()
	for email := range emailChannel {
		fmt.Println("Sending email to ", email)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// messageChannel := make(chan string)

	// messageChannel <- "Hello World"  // blocking operation
	// message := <-messageChannel
	// fmt.Println(message)

	// numChan := make(chan int, 3)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// time.Sleep(1 * time.Second)

	result := make(chan int)
	go sum(result, 10, 20)
	res := <-result // blocking operation
	fmt.Println("Sum is ", res)

	done := make(chan bool)
	go task(done)
	<-done // blocking operation

	// buffered channels
	emailChannel := make(chan string, 3)
	emailChannel <- "pikachu@pokemon.com"
	emailChannel <- "Ash@pokemon.com"
	emailChannel <- "Misty@pokemon.com"
	// emailChannel <- "email4" // deadlock

	fmt.Println(<-emailChannel)
	fmt.Println(<-emailChannel)
	fmt.Println(<-emailChannel)

	for i := 0; i < 5; i++ {
		go func(i int) {
			emailChannel <- "Example" + fmt.Sprint(i) + "@gmail.com"
		}(i)
	}

	done = make(chan bool)
	go emailSender(emailChannel, done)
	// close(emailChannel) // close the channel
	<-done
}
