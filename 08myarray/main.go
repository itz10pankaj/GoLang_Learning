package main

import "fmt"

func main() {
	fmt.Println("Study About array")

	// var fruitList [4]string
	var fruitList = [4]string{"Apple", "Mango", "bnanan", "grapes"}
	// fruitList[0] = "Apple"
	// fruitList[1] = "Mango"
	// fruitList[2] = "bnanan"
	// fruitList[3] = "grapes"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Length of Fruit List is", len(fruitList))
}
