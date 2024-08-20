package main

import (
	"fmt"
	"math"
)

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	result := a - b
	return int(math.Abs(float64(result)))
}

// func getLanguages() (string, string, string) {
// 	return "golang", "javascript", "python"
// }

func processIt(fn func(int, int) int, x int, y int) int {
	return fn(x, y)
}

func main() {
	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)
	// result := add(3, 5)
	// fmt.Println(result)
	result := processIt(add, 3, 5)
	fmt.Println("Result:", result)
	result2 := processIt(sub, 3, 5)
	fmt.Println("Result:", result2)

}
