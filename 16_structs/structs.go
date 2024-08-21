package main

import (
	"fmt"
	"time"
)

// order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision

}

func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	order := order{
		id:     "1",
		amount: 50,
		status: "received",
	}

	order.createdAt = time.Now()
	order.changeStatus("return")
	fmt.Println("Order", order)
	fmt.Println("Order_amount", order.amount)
}
