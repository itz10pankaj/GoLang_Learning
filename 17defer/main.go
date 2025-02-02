package main

import "fmt"

func main() {

	//Last on First Out

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Defer")
	// for i := 0; i < 3; i++ {
	// 	defer fmt.Println(i)
	// }
}
