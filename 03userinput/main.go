package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user,hello"
	fmt.Println(welcome)

	fmt.Println("ENTER READING FOR OUR PIZZA:")

	// COMMA OK SYNTAX || ERRROR OK
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving rating", input)
	fmt.Printf("Type of input %T \n", input)
}
