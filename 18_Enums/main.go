package main

import "fmt"

// Enums are a set of named integer constants
// enumerated types are defined using the keyword "const" and "iota" keyword

// iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.

type OrderStatus int

const (
	Pending OrderStatus = iota
	Processing
	Delivered
	Cancelled
)

func changeOrderstatus(orderStatus OrderStatus) {
	switch orderStatus {
	case Pending:
		fmt.Println("Order is Pending")
	case Processing:
		fmt.Println("Order is Processing")
	case Delivered:
		fmt.Println("Order is Delivered")
	case Cancelled:
		fmt.Println("Order is Cancelled")
	default:
		fmt.Println("Invalid Order Status")
	}
}

func main() {
	changeOrderstatus(Pending)
	changeOrderstatus(Processing)
	changeOrderstatus(Delivered)
	changeOrderstatus(Cancelled)
	changeOrderstatus(5)
}
