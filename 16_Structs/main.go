package main

import (
	"fmt"
	"time"
)

// struct is a user-defined type that contains named fields
// struct fields can be of any type, including other structs

// struct embadding
type Customer struct {
	name    string
	address string
	orders  []Order
}

// struct is a collection of fields
type Order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time // nano second precision
	customer  Customer
}

// struct constructor
func newOrder(id string, amount float64) Order {
	return Order{
		id:        id,
		amount:    amount,
		status:    "pending",
		createdAt: time.Now(),
	}
}

// struct methods
// (O *Order) is a receiver of the method
func (O *Order) changeStatus(status string) {
	O.status = status
}

func (O Order) print() string {
	return fmt.Sprintf("Order ID: %s, Amount: %.2f, Status: %s, Created At: %v", O.id, O.amount, O.status, O.createdAt.Format(time.RFC3339))
}

func main() {
	// create a new struct
	order := Order{
		id:        "123",
		amount:    100.00,
		status:    "pending",
		createdAt: time.Now(),
	}

	// update struct fields
	order.status = "completed"

	// access struct fields
	fmt.Println(order.id)

	// print the struct
	fmt.Println(order)

	myOrder := Order{
		id:        "456",
		amount:    200.00,
		status:    "Delivered",
		createdAt: time.Now(),
	}

	// call struct method
	myOrder.changeStatus("completed")

	fmt.Println(myOrder)

	fmt.Println(myOrder.print())

	// create a new struct using constructor
	newOrder := newOrder("789", 300.00)
	fmt.Println(newOrder.print())
	newOrder.changeStatus("completed")
	fmt.Println(newOrder.print())

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// struct embadding
	customer := Customer{
		name:    "John Doe",
		address: "123 Main St",
		orders: []Order{
			Order{id: "1", amount: 100.00, status: "pending", createdAt: time.Now()},
			Order{id: "2", amount: 200.00},
			Order{id: "3", amount: 300.00},
		},
	}

	fmt.Println(customer)

	newCustomer := Customer{
		name:    "Jane Doe",
		address: "456 Main St",
	}

	newCustomer.orders = append(newCustomer.orders, Order{id: "4", amount: 400.00})
	fmt.Println(newCustomer)
}
