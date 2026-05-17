package main

import "fmt"

func main() {
	// while loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// for loop

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	// range
	for i := range 3 {
		fmt.Println(i)
	}
}
