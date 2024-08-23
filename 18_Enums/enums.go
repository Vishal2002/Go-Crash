package main

import "fmt"

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Processing
// 	Prepared
// 	Delivered
// ) //integer

type OrderStatus string

const (
	Received   OrderStatus = "received"
	Processing             = "processing"
	Prepared               = "prepared"
	Delivered              = "delivered"
)

//enumerated type

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}
func main() {

	changeOrderStatus(Received)

}
