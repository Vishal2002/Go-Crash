package main

import "fmt"

//only have for loop as construct for looping
func main() {

	//while loop

	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	//infinite loop

	// for {
	// 	fmt.Println("I am in Infinite Loop")
	// }

	//for loop
	table := 2
	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d\n", table, i, table*i)
	}

}
