package main

import "fmt"

// Define a generic function to print slices
func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {

	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"apple", "banana", "cherry"}
	// boolSlice := []bool{true, false, true}

	// Print the slices using the generic function
	fmt.Println("Integer slice:")
	printSlice(intSlice)

	fmt.Println("String slice:")
	printSlice(stringSlice)

	// 	fmt.Println("Boolean slice:")
	// 	printSlice(boolSlice)
}
