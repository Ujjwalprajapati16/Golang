package main

import "fmt"

// func printSLice(item []int) {
// 	for _, i := range item {
// 		println(i)
// 	}
// }

// T any -> T is a type parameter
func printSLice[T any](item []T) {
	for _, i := range item {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func printSLice2[T int | string | bool](item []T) {
	for _, i := range item {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

type Stack[T any] struct {
	elements []T
}

func main() {
	printSLice([]int{1, 2, 3, 4, 5})               // 1 2 3 4 5
	printSLice([]string{"a", "b", "c", "d", "e"})  // a b c d e
	printSLice([]float64{1.1, 2.2, 3.3, 4.4, 5.5}) // 1.1 2.2 3.3 4.4 5.5

	myStack := Stack[int]{elements: []int{1, 2, 3, 4, 5}}
	printSLice(myStack.elements) // 1 2 3 4 5

	myStack2 := Stack[string]{elements: []string{"a", "b", "c", "d", "e"}}
	printSLice(myStack2.elements) // a b c d e

	printSLice2([]int{1, 2, 3, 4, 5}) // 1 2 3 4 5
}
