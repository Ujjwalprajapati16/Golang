package main

import "fmt"

// numbered squence of specific length
func main() {
	var nums [4]int

	//array length
	fmt.Println(len(nums))

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	// access single element
	fmt.Println(nums[0])

	// print whole array
	fmt.Println(nums)

	var vals [4]bool
	fmt.Println(vals) // all values are false

	var name [3]string
	fmt.Println(name) // all values are empty

	var nums3 [4]int
	fmt.Println(nums3) // all values are 0

	var nums2 = [4]int{1, 2, 3, 4}
	fmt.Println(nums2)

	nums4 := [4]int{1, 2, 3, 4}
	fmt.Println(nums4)

	//2D array
	nums5 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums5)

	// - fixed size, that is predictable
	// - Memory optimisation
	// - Constant time access
}
