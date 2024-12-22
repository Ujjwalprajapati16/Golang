package main

import "fmt"

// functions -> reusable blocks of code
// func name_of_function(parameters) return_type { code }
func add(a int, b int) int {
	return a + b
}

// multiple return values
func getLanguage() (string, string) {
	return "Go", "Python"
}

// passing functions as arguments
func processit(fn func(int, int) int) {
	result := fn(10, 20)
	fmt.Println(result)
}

func processit2() func(a int) int {
	return func(a int) int {
		return a * a
	}
}

func main() {
	result := add(1, 2)
	fmt.Println(result)

	lang1, lang2 := getLanguage()
	fmt.Println(lang1, lang2)

	// anonymous functions
	processit(func(a int, b int) int { return a + b })

	// closures
	f := processit2()
	fmt.Println(f(10))
}
