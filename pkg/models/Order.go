package models

type Order struct {
	id               int32
	discount, amount float64
	paymentMethod    string
	rating, duration int32
	cuisine string
	time int32
	verified bool
	customer Customer
	restaurant Restaurant
}
