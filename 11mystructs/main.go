package main

import "fmt"

func main() {
	fmt.Println("structs")

	// no inheritance in golang
	hitesh := User{"Hitesh", "hitesh@go.com", true, 16}

	fmt.Println(hitesh)
	fmt.Printf("Name of user is %v\n", hitesh.Name)
	fmt.Printf("Detail of user is %+v", hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
