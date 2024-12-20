package main

import (
	"fmt"
	"maps"
)

// maps -> key:value pairs
func main() {
	// creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "John"
	m["surname"] = "Doe"

	fmt.Println(m)

	// getting an element
	name, ok := m["name"]
	if ok {
		fmt.Println(name)
	} else {
		fmt.Println("Key does not exists")
	}

	// If key does not exists then it returns zero value
	// if key does not exists then it returns false
	fmt.Println(m["age"])

	mp := make(map[string]int)

	mp["a"] = 1
	mp["b"] = 2
	mp["c"] = 3

	fmt.Println(mp["a"])

	// deleting an element
	delete(mp, "b")
	fmt.Println(mp["b"])

	// clear
	clear(mp)
	fmt.Println(mp)

	// creating map without make - map with literal
	mp2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(mp2)

	// Maps package has functions for working with maps
	fmt.Println(maps.Keys(mp2))
	fmt.Println(maps.Values(mp2))
	fmt.Println(maps.Equal(mp2, mp))

	// Iterating over map
	for key, value := range mp2 {
		fmt.Println(key, value)
	}
}
