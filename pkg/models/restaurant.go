package models

type Restaurant struct {
	Id int32
	RestaurantName string
	Address Address
	Items []Item
}
