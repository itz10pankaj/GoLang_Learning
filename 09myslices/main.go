package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 23
	highScore[1] = 24
	highScore[2] = 25
	highScore[3] = 26
	// highScore[4] = 27 // give error
	highScore = append(highScore, 29, 28)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value form slice based on index
	index := 2
	highScore = append(highScore[:index], highScore[index+1:]...)
	fmt.Println(highScore)
}
