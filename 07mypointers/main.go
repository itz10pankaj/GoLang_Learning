package main

import "fmt"

func main() {
	fmt.Println("hello world")

	// var ptr *int
	// fmt.Println("value of ptr", ptr)
	// fmt.Printf("type of ptr is %T", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("value of ptr", ptr)
	fmt.Println("value of ptr", *ptr)
	fmt.Printf("type of ptr is %T \n", ptr)

	*ptr = *ptr + 2
	fmt.Println(myNumber)

}
