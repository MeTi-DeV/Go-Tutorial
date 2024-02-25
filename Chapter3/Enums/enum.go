package main

import "encoding/json"

type Order struct {
	id     int
	price  int
	status OrderStatus
}

type OrderStatus string

const (
	Created OrderStatus = "Created"
	Pending OrderStatus = "Pending"
	Success OrderStatus = "Success"
	Failed  OrderStatus = "Failed"
	Refund  OrderStatus = "Refund"
)

func main() {
	println(Created)
	println(Pending)
	println(Failed)
	println(Refund)

	order1 := Order{
		id:     1,
		price:  100,
		status: Created,
	}
	order2 := Order{
		id:     2,
		price:  300,
		status: Pending,
	}
	order3 := Order{
		id:     3,
		price:  400,
		status: Failed,
	}
	order1Json, _ := json.Marshal(order1)
	order2Json, _ := json.Marshal(order2)
	order3Json, _ := json.Marshal(order3)
	println(string(order1Json))
	println(string(order2Json))
	println(string(order3Json))

}
