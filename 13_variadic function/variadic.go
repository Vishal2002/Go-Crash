package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num

	}
	return total
}

// variadic function can take a variable number of arguments
func main() {
	total := sum(1, 23, 4, 43, 23, 13, 31, 2, 1, 2, 12, 1, 2)
	fmt.Println(total)
}
