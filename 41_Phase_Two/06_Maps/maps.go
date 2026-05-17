package main

import (
	"fmt"
	"maps"
)

func main() {

	// Maps
	// creating map
	// m := make(map[string]string)
	// m["P"]="Pankaj"
	// m["A"]= "Atul"

	// fmt.Println(m)
	// fmt.Println(m["jj"]) // it will return zero value

	// fmt.Println(len(m))
	// delete(m,"P")
	// fmt.Println(m)

	// clear(m)
	// fmt.Println(m)

	// Second method to create map
	// m:=map[string]int{"Price":344,"age":23}
	// fmt.Println(m)

	// k,ok := m["Price"]

	// if ok {
	//     fmt.Println(k)
	// }else {
	//     fmt.Println("ELement not found")
	// }

	m2 := map[string]int{"Price": 384, "age": 23}
	m1 := map[string]int{"Price": 384, "age": 23}
	fmt.Println(maps.Equal(m1, m2))

}
