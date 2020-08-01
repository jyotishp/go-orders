package models

type Item struct {
	Id int32
	Name, Cuisine string
	Discount, Amount float32
}

type ItemFilter struct {
	RestaurantId int32
	Min float32
	Max float32
}
