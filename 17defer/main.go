package main

import "fmt"

func main() {
	defer fmt.Println("Hello World")
	defer fmt.Println("Hello World 2")
	fmt.Println("Welcome to derfer in Golang")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
