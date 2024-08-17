package main

import "fmt"

func main() {
	var num [5]string = [5]string{"a", "b", "c", "d", "e"}
	number := []interface{}{"a", 5, true, "c", "d"}

	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(len(number))
	fmt.Println(number[0])

	//2-d arry

	matrix := [2][2]int{{1, 2}, {1, 3}}

	fmt.Println(matrix)

	// fixed size and if we want to use the dynamic array we use slices
	//array have constant time access

}
