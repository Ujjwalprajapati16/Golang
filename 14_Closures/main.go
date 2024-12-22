package main

import "fmt"

// Closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	counter1 := counter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())

	counter2 := counter()
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())
}
