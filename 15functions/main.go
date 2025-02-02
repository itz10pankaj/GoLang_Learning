package main

import "fmt"

func greeterTwo() {
	fmt.Println("Another function")
}
func main() {
	fmt.Println("Functions")
	// greeter()
	// greeterTwo()

	// result := adder(3, 4)
	// fmt.Println("Result is:", result)

	// result := proAdder(1, 2, 3, 4)
	// fmt.Println("Result is:", result)

	result, Mymessage := proAdder(1, 2, 3, 4)
	fmt.Println(Mymessage, result)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// func proAdder(values ...int) int {
// 	total := 0
// 	for _, val := range values {
// 		total += val
// 	}
// 	return total
// }

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Resukt is "
}

func greeter() {
	fmt.Println("Hello from Golang")
}
