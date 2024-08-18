package main

import (
	"fmt"
	"maps"
)

func main() {

	//map is similar to hashmap,dictionary

	//creating a map in go
	m := make(map[string]string)

	m["name"] = "Vishal"
	m["surname"] = "Sharma"
	fmt.Println(m["name"])
	fmt.Println(m["surname"])

	//if key value doesnot exist then it return zero value
	//int->0, string ->"" , bool -> false

	//methods of map
	// Equals, clone and others

	m1 := map[string]string{"name": "Vishal", "surname": "Sharma"}
	m2 := map[string]string{"name": "Vishal", "surname": "Sharma"}
	fmt.Println(maps.Equal(m1, m2))

}
