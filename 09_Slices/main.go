package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic array
// most used construct in go
// + useful methods
func main() {
	// uninitialised slice is nil
	var nums []int
	fmt.Println(nums)

	fmt.Println(nums == nil) // true
	fmt.Println(len(nums))
	// initialise slice
	nums = []int{1, 2, 3}
	fmt.Println(nums)

	// append to slice
	nums = append(nums, 4)
	fmt.Println(nums)

	// make(data type, length, capacity)
	var nums1 = make([]int, 3, 5) // make returns slice
	fmt.Println(nums1)

	// capacity -> no. of elements that can be stored in the slice
	fmt.Println(cap(nums)) // 3

	// append
	nums1 = append(nums1, 5)
	fmt.Println(nums1)      // [0 0 0 5]
	fmt.Println(cap(nums1)) // 5

	// empty slice
	var nums2 = make([]int, 0, 5)
	fmt.Println(nums2)
	fmt.Println(cap(nums2))
	fmt.Println(len(nums2))

	// add by indexing
	nums1[0] = 1
	nums1[1] = 2
	fmt.Println(nums1)

	// copy function
	var nums3 = make([]int, 2)
	// copy(destination, source)
	copy(nums3, nums1)
	fmt.Println(nums3)

	// slice operator
	var nums4 = []int{1, 2, 3, 4, 5, 6}
	// nums4[start:end] end is exclusive
	fmt.Println(nums4[2:4])
	fmt.Println(nums4[2:])
	fmt.Println(nums4[:4])

	// delete from slice
	nums4 = append(nums4[:2], nums4[3:]...)
	fmt.Println(nums4)

	var nums5 = []int{1, 2, 3, 4, 5, 6}
	// nums5[start:end:cap(nums5)] end is exclusive
	fmt.Println(nums5[2:4:cap(nums5)])
	fmt.Println(nums5[2:])
	fmt.Println(nums5[:4])

	//slice package
	fmt.Println(slices.Equal(nums4, nums5))

	// multi-dimensional slice
	var nums6 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums6)
}
