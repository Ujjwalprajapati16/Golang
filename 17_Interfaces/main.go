package main

import "fmt"

// Interfaces are a way to define a contract that a type should follow.

type paymenter interface {
	pay(amount float32)
	refund(amount float32, accepted bool)
}

type payment struct {
	// gatway stripe
	gatway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)

	// stripePaymentGW := stripe{}
	// stripePaymentGW.pay(amount)

	p.gatway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Amount paid using razorpay: ", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Amount paid using stripe: ", amount)
}

type fakePaymentGatway struct{}

func (f fakePaymentGatway) pay(amount float32) {
	fmt.Println("Amount paid using fake payment gateway: ", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Amount paid using paypal: ", amount)
}

func (p paypal) refund(amount float32, accepted bool) {
	fmt.Println("Amount refunded using paypal: ", amount)
}

func main() {
	// razorpayPaymentGW := razorpay{}
	// stripePaymentGW := stripe{}
	// fakeGatway := fakePaymentGatway{}
	paypalGatway := paypal{}
	paymentGateway := payment{
		gatway: paypalGatway,
	}
	paymentGateway.makePayment(100)
}
