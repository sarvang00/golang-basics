package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// no inheritance/super/parent/child in Go

	sarvang := User{"Sarvang", "gamechanger.cim@gmail.com", true, 21}
	fmt.Println(sarvang)
	fmt.Printf("Sarvang's details are: %+v\n", sarvang)
	fmt.Printf("User %v's email id is: %v\n", sarvang.Name, sarvang.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
