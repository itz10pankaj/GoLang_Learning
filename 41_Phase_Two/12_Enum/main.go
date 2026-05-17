package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	confirmed
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Chaneing the order sattus", status)
}

func main() {
	changeOrderStatus(confirmed)
}
