package main

import (
	"fmt"
	"time"
)

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func ProcessNum(numchan chan int) {
// 	for num := range numchan {
// 		fmt.Println("Processed Number:", num)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func sum(result chan int, num1 int, num2 int) {
// 	result <- num1 + num2
// }

func EmailSender(emailchan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailchan {
		fmt.Println("Sending Email to", email)
		time.Sleep(time.Second)
	}
}
func main() {
	// rand.Seed(time.Now().UnixNano())

	// numchan := make(chan int)

	// go ProcessNum(numchan)

	// for i := 0; i < 10; i++ {
	// 	randomNum := rand.Intn(100)
	// 	numchan <- randomNum
	// }

	// close(numchan)

	// time.Sleep(time.Second * 2)

	// result := make(chan int)

	// go sum(result, 2, 3)
	// fmt.Println(<-result)

	// Lets Implement a queue system for Emailing Service

	emailchan := make(chan string, 100)
	done := make(chan bool)

	go EmailSender(emailchan, done)

	for i := 0; i < 5; i++ {
		emailchan <- fmt.Sprintf("%d@gamil.com", i)
	}
	fmt.Println("Done")
	close(emailchan)
	<-done

}
