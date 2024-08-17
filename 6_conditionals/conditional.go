package main

import "fmt"

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("Your are Adult")
	// } else if age >= 12 {
	// 	fmt.Println("You are Teenager")
	// } else {
	// 	fmt.Println("You are Kid")
	// }

	//You can also initialize the variable in if construct
	if age := 15; age >= 18 {
		fmt.Println("You are Adult", age)
	} else if age >= 12 {
		fmt.Println("You are Teenager", age)
	} else {
		fmt.Println("You are Kid", age)
	}

	//there is no ternary operator in Go

}
