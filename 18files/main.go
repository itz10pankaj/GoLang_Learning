package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handling")

	content := "This need to go in a file"

	file, err := os.Create("./sample.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is", length)
	defer file.Close()
	readFile("./sample.txt")
}

func readFile(filname string) {
	databyte, err := os.ReadFile(filname)
	if err != nil {
		panic(err)
	}

	// fmt.Println(databyte)
	fmt.Println(string(databyte))
}
