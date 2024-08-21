package main

import "fmt"

//by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

//by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {

	num := 1
	fmt.Println("Memory address", &num)
	changeNum(&num)
	fmt.Println("After changeNum", num)

}