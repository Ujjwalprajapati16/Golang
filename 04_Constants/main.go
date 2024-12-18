package main

import "fmt"

// const can be decalred outside main
const age = 30

func main() {
	const name = "golang" // after assigning const it cannot be changed
	const version = 1.8

	fmt.Println(name, version)

	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
