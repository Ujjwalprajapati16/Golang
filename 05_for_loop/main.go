package main

import "fmt"

// for -> only construct in go for looping

func main() {
	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//inifinte loop
	// for {
	// 	fmt.Println("loop")
	// }

	// classic for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		if i == 2 {
			continue
		}

		if i == 3 {
			break
		}
	}

	// range - exclude the given number
	for i := range 3 {
		fmt.Println(i)
	}
}
