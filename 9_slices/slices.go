package main

import (
	"fmt"
	"slices"
)

func main() {
	//uninitialized slice is nil
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2) // make(type,len,capacity)
	//capacity is max number of elements can fit as this is slice so capacity can change as we append more
	fmt.Println(cap(nums)) //cap->2
	nums = append(nums, 2) // [0,0,2]
	nums = append(nums, 3) // [0,0,2,3]
	nums = append(nums, 3) // [0,0,2,3,3]
	fmt.Println(cap(nums))
	fmt.Println(nums)

	//slice operator
	var dam = []int{1, 2, 3, 4}
	fmt.Println(dam[1:4])

	//methods  --there are lots of method of slices like Equal,Concat,Reverse and lots
	var slice1 = []int{1, 2, 3, 4}
	var slice2 = []int{1, 2, 2, 4}
	fmt.Println(slices.Equal(slice1, slice2))

}
