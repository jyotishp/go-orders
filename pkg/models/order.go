package models

type Order struct {
	Id               int32
	Discount, Amount float32
	PaymentMethod    string
	Rating, OrderDuration int32
	Cuisine string
	OrderTime int32
	Verified bool
	Customer Customer
	Restaurant Restaurant
	OrderItems []Item
}
