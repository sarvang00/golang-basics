package main

import "fmt"

// A public variable
const LoginToken string = "nvjdsnv"

func main() {
	var username string = "sarvang"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.4558223441
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default values and aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Another variable is of type: %T \n", anotherVar)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// not var style (using walrus) --> cannot use outside of a method
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	// Printing a public variable
	fmt.Println(LoginToken)
	fmt.Printf("LoginToken variable is of type: %T \n", LoginToken)
}
