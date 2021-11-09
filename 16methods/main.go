package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")

	fmt.Println("Structs in Golang")
	// no inheritance/super/parent/child in Go

	sarvang := User{"Sarvang", "gamechanger.cim@gmail.com", true, 21}
	fmt.Println(sarvang)
	fmt.Printf("Sarvang's details are: %+v\n", sarvang)
	fmt.Printf("User %v's email id is: %v\n", sarvang.Name, sarvang.Email)

	// Using our method
	sarvang.GetStatus()

	sarvang.NewMail()

	fmt.Println("Email: ", sarvang.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// below is a method; copy of object (User) is passed
func (u User) GetStatus() {
	fmt.Println("User status: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
