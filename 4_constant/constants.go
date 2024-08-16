package main

import "fmt"

func main() {
	const dessert = "Jalebi🥨🥨"
	// dessert="Gulab Jamun" throws an error

	//we can also assign multiple constant vakues
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(dessert)
	fmt.Println(port, host)
}
