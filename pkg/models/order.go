package models

type Order struct {
	Id               int32
	Discount, Amount float64
	PaymentMethod    string
	Rating, Duration int32
	Cuisine string
	Time int32
	Verified bool
	Customer Customer
	Restaurant Restaurant
}
