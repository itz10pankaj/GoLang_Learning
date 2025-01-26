package main

import "fmt"

func main() {
	fmt.Println("hello maps")

	rollNo := make(map[int]string)
	rollNo[1] = "Pankaj"
	rollNo[2] = "Nishant"
	rollNo[3] = "Ashu"

	// fmt.Println(rollNo)
	// fmt.Println("Roll no 1 is", rollNo[1])

	delete(rollNo, 3)
	// fmt.Println(rollNo)

	//loops
	for key, value := range rollNo {
		fmt.Printf("For key %v, and value is %v\n", key, value)
	}
}
