package models

type Restaurant struct {
	Id int32
	Name string
	Address Address
	Items []Item
}

type RestaurantNoItems struct {
	Id int32
	Name string
	Address Address
}