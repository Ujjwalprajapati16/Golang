package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are not an adult")
	}

	// else if
	age = 16

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a child")
	}

	var role = "admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions {
		fmt.Println("You can access the dashboard")
	}

	if role == "admin" && hasPermissions {
		fmt.Println("You can access the dashboard")
	}

	// shorthand or we can declare a variable inside if
	if age := 15; age >= 18 {
		fmt.Println("You are an adult", age)
	} else if age >= 12 {
		fmt.Println("You are a teenager", age)
	} else {
		fmt.Println("You are a child", age)
	}

	// go does not have ternary, you will have to use noraml if else
}
