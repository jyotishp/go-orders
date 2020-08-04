package db

type ItemIp struct {
	RestaurantId int32
	Name string
	ItemId int32
	Cuisine string
	Discount float32
	Amount float32
}

type ItemUp struct {
	Cuisine string `json:":c"`
	Discount float32 `json:":d"`
	Amount float32 `json:":a"`
}

type Customer struct {
	Id int32
	Name string
	Address Address
}

type CustomerNoId struct {
	Name string
	Address Address
}

type Address struct {
	Line1 string
	Line2 string
	City string
	State string
}

type Item struct {
	Id int32
	Name string
	Cuisine string
	Discount float32
	Amount float32
}

type OrderCustomer struct {
	Name string
	Address Address
}

type Restaurant struct {
	Id int32
	Name string
	OrderCount int32
	Address Address
	Items []Item
}

type RestaurantNoId struct {
	Name string `json:":n"`
	//OrderCount int32 `json:":oc"`
	Address Address `json:":a"`
	Items []Item `json:":itms"`
}

type RestaurantNoItems struct {
	Id int32
	Name string
	OrderCount int32
	Address Address
}

type Order struct {
	Id int32
	Discount float32
	Amount float32
	PaymentMethod string
	Rating int32
	Duration int32
	Cuisine string
	Time int32
	Verified bool
	Customer Customer
	Restaurant RestaurantNoItems
	Items []Item
}

type OrderIp struct {
	Id int32
	Discount float32 `json:":d"`
	Amount float32 `json:":amt"`
	PaymentMethod string `json:":pm"`
	Rating int32 `json:":r"`
	Duration int32 `json:":dn"`
	Cuisine string `json:":c"`
	Time int32 `json:":t"`
	Verified bool `json:":v"`
	Customer Customer `json:":ctmr"`
	Restaurant RestaurantNoItems `json:":rt"`
	Items []Item `json:":itms"`
}

type OrderNoId struct {
	Discount float32 `json:":d"`
	Amount float32 `json:":amt"`
	PaymentMethod string `json:":pm"`
	Rating int32 `json:":r"`
	Duration int32 `json:":dn"`
	Cuisine string `json:":c"`
	Time int32 `json:":t"`
	Verified bool `json:":v"`
	Customer Customer `json:":ctmr"`
	Restaurant RestaurantNoItems `json:":rt"`
	Items []Item `json:":itms"`
}