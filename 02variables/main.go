package main

import "fmt"

const LoginTable string = "hello login" // Capital L makes it Public

func main() {
	// // fmt.Println("Hello, World!")
	// var username string = "Pankaj"
	// fmt.Println(username)
	// fmt.Printf("variable is of type: %T \n", username)

	// 	var isLoggedIn bool = true
	// 	fmt.Println(isLoggedIn)
	// 	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	// var smallVal uint8 = 255
	// fmt.Println(smallVal)
	// fmt.Printf("variable is of type: %T \n", smallVal)

	// var FloatVal float64 = 255.73482998574859878237
	// fmt.Println(FloatVal)
	// fmt.Printf("variable is of type: %T \n", FloatVal)

	// default values and some alaise
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	//implict type
	// var webSite = "Pankaj" // laxer ne apne aap decide ker liya
	// fmt.Println(webSite)
	// fmt.Printf("variable is of type: %T \n", webSite)
	// webSite = 5 // error

	// no var style
	// numberOfUser := 300 // this method is only alow in some method not outside of any method
	// fmt.Println(numberOfUser)

	fmt.Println(LoginTable)

}
