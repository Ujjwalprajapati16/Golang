package main

// iterating over data structures

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	sum := 0

	// index , value
	for _, num := range nums {
		fmt.Println(num)
		sum = sum + num
	}

	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	// map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// for keys only
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on string
	// unicode code point rune
	// starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
