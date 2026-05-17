package main

import (
	"fmt"
)

//	func PrintSliceInt(items []int) {
//		for i, item := range items {
//			fmt.Println("Item at", i, "is", item)
//		}
//	}
func PrintSliceString(items []string) {
	for i, item := range items {
		fmt.Println("Item at", i, "is", item)
	}
}
func PrintSlice[T any](items []T) {
	for i, item := range items {
		fmt.Println("Item at", i, "is", item)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	// nums := []int{1, 2, 3}
	nums := []string{"A", "B", "C"}
	PrintSlice(nums)

	myStruct := stack[string]{
		elements: []string{"A", "B", "C"},
	}
	PrintSlice(myStruct.elements)
}
