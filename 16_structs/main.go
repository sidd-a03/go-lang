package main

import (
	"fmt"
	"time"
	"uuid"
)

type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func newOrder(id string, amount float32, status string, createdAt time.Time) *order {
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: createdAt,
	}

	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	myNewOrder := order{
		id: uuid.NewV4().String(),
		amount:    52.02,
		status:    "done",
		createdAt: time.Now(),
		customer: customer{
			name:  "Achinta",
			phone: "861******1",
		},
	}

	// initial setup
	myOrder := newOrder(uuid.NewV4().String(), 52.02, "done", time.Now())

	myOrder.changeStatus("failed")

	fmt.Println(myOrder.getAmount()) // 0

	fmt.Println(*myOrder)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
	fmt.Println(myNewOrder)
}
