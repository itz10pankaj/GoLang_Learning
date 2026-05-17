package main

import "fmt"

// type payment struct {
// }

// func (p payment) makepayment(amount float32) {
// 	razorPayPaymentGw := razorpay{}

// 	razorPayPaymentGw.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	// logic to make payemnt
// 	fmt.Println("mAKING PAYMENT USING RAZORPAY", amount)
// }
// func main() {
// 	MyPayment := payment{}
// 	MyPayment.makepayment(333.34)
// }

// The above method id correct but if we have to change raporpay to slices we have to make chanes in
// makepayment  which is not good

// type payment struct {
// 	gateway razorpay
// }

// func (p payment) makepayment(amount float32) {

// 	p.gateway.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	// logic to make payemnt
// 	fmt.Println("mAKING PAYMENT USING RAZORPAY", amount)
// }
// func main() {
// 	Gateway := razorpay{}
// 	MyPayment := payment{
// 		gateway: Gateway,
// 	}
// 	MyPayment.makepayment(333.34)
// }

// IN the method method also we have to pass gateway in struct so this also do not solve problem

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makepayment(amount float32) {

	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("mAKING PAYMENT USING RAZORPAY", amount)
}

type stripe struct{}

func (r stripe) pay(amount float32) {
	fmt.Println("mAKING PAYMENT USING stripe", amount)
}

type fakePayment struct{}

func (r fakePayment) pay(amount float32) {
	fmt.Println("mAKING PAYMENT USING fakePayment", amount)
}

func main() {
	Gateway := stripe{}
	newPayment := payment{
		gateway: Gateway,
	}
	newPayment.makepayment(222)

}
