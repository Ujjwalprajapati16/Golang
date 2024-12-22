package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	// variadic functions
	// A variadic function is a function that accepts a variable number of arguments.

	// Here’s a function that will take an arbitrary number of ints as arguments.
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // [1 2 3 4 5 6 7 8 9 10] 55

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(nums...)) // [1 2 3 4 5 6 7 8 9 10] 55
}
