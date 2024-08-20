package main

import "fmt"

//iterating over data structure
func main() {
	//Slices

	// nums := []int{1, 2, 3, 4, 5}
	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println(num, i) //so the i is for the index
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	//Maps

	m := map[string]string{"name": "Vishal", "lastName": "Sharma"}
	for key, val := range m {
		fmt.Println(key, val)

	}
}
