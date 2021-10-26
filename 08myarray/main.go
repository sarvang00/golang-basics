package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	// These are rarely used

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of fruit list is: ", len(fruitList))

	var veggieList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie list is: ", veggieList)
	fmt.Println("Veggie list length is: ", len(veggieList))
}
