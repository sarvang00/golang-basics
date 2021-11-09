package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proResult, proMsg := proAdder(1, 2, 3, 8)
	fmt.Println("Pro result is: ", proResult)
	fmt.Println("Pro message is: ", proMsg)

}

func greeter() {
	fmt.Println("Namaste from Golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Hello Pro"
}
