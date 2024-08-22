package main

import "fmt"

// PaymentGateway interface defines the behavior for payment gateways
type PaymentGateway interface {
	Pay(amount float32)
}

// Payment struct now accepts any PaymentGateway
type Payment struct {
	gateway PaymentGateway
}

// MakePayment uses the injected PaymentGateway to make the payment
func (p Payment) MakePayment(amount float32) {
	p.gateway.Pay(amount)
}

// RazorPay implements PaymentGateway
type RazorPay struct{}

func (r RazorPay) Pay(amount float32) {
	fmt.Println("Making payment with RazorPay:", amount)
}

// PayPal implements PaymentGateway
type PayPal struct{}

func (p PayPal) Pay(amount float32) {
	fmt.Println("Making payment with PayPal:", amount)
}

func main() {
	// Using RazorPay
	razorPayment := Payment{gateway: RazorPay{}}
	razorPayment.MakePayment(200)

	// Using PayPal
	payPalPayment := Payment{gateway: PayPal{}}
	payPalPayment.MakePayment(150)
}
