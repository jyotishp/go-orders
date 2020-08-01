package models

type Order struct {
	Id               int32
	Discount, Amount float32
	PaymentMethod    string
	Rating, Duration int32
	Cuisine string
	Time int32
	Verified bool
	Customer Customer
	Restaurant RestaurantNoItems
	Items []Item
}

type OrderIp struct {
	Id               int32
	Discount, Amount float32
	PaymentMethod    string
	Rating, Duration int32
	Cuisine string
	Time int32
	Verified bool
	CustomerId int32
	RestaurantId int32
	Items []int32
}
