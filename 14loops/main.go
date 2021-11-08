package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in  Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// By index
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// By element index range
	for idx := range days {
		fmt.Println(days[idx])
	}

	// By element
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("The value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at LCO Pro")
}
