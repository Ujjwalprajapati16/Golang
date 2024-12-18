package main

import "fmt"

func main() {
	var name string = "Golang"
	var version float64 = 1.8
	fmt.Println(name, version)

	//infer
	var age = 20
	fmt.Println(age)

	var isAdult = true
	fmt.Println(isAdult)

	//short declaration
	name2 := "Golang"
	version2 := 1.8
	fmt.Println(name2, version2)

	var price float32 = 50.5
	var quantity int = 5
	var total = price * float32(quantity)
	fmt.Println(total)
}
