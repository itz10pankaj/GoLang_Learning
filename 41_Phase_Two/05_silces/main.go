package main

import "fmt"

// -> Dynamic Arrays
func main() {
	// uninitialized Slice in nil
	// var nums []int
	// fmt.Println((nums))
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2)
	fmt.Println((nums))
	fmt.Println(cap(nums))
}
