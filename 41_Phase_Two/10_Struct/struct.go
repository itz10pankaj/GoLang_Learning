package main

import (
	"fmt"
	"time"
)

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o *order) changeStatus(status string) {
	o.status = "Paid"
}
func (o order) getAmount() float32 {
	return o.amount
}
func createOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}
func main() {
	fmt.Println("Start small. Ship something.")
	// First Method to make instance of a struct
	myorder := order{
		id:     "1",
		amount: 50.6,
		status: "Done",
	}

	// second method
	myorder2 := createOrder("2", 45.6, "Deli")
	fmt.Println(myorder2)

	myorder.createdAt = time.Now()
	myorder.changeStatus("Paid")
	fmt.Println(myorder.getAmount(), myorder.status)
}
