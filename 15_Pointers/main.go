package main

import "fmt"

//Pointers
// Values passed as the copy
// If we want to change the value of the variable, we need to pass the address of the variable.
func changeNumber(num int) {
	num = 5
	fmt.Println("Inside function:", num)
}

// pass by refrence
// &num is the address of the variable
// *int is the type of the variable
func changeNumber2(num *int) {
	*num = 5
	fmt.Println("Inside function:", *num)
}

func main() {
	num := 1
	changeNumber(num)
	fmt.Println("Outside function:", num)
	fmt.Println("Address of num:", &num)

	changeNumber2(&num)
	fmt.Println("Outside function:", num)
}
