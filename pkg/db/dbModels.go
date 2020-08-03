package db

type ItemIp struct {
	RestaurantId int32
	Name string
	ItemId int32
	Cuisine string
	Discount float32
	Amount float32
}
//
//type ItemUp struct {
//	Cuisine string `json:":ic"`
//	Discount float32 `json:":idc"`
//	Amount float32 `json:":iamt"`
//}
//
//type Customer struct {
//	Id int32
//	Name string `json:":cn"`
//	Address Address `json:":cadr"`
//}
//
//type CustomerNoId struct {
//	Name string `json:":cn"`
//	Address Address `json:":cadr"`
//}
//
//type Address struct {
//	Line1 string `json:"Line1"`
//	Line2 string `json:"Line2"`
//	City string `json:"City"`
//	State string `json:"State"`
//}
//
//type Item struct {
//	Id int32
//	Name string `json:"Name"`
//	Cuisine string `json:"Cuisine"`
//	Discount float32 `json:"Discount"`
//	Amount float32 `json:"Amount"`
//}
//
//type OrderCustomer struct {
//	Name string `json:"Name"`
//	Address Address `json:"Address"`
//}
//
//type Restaurant struct {
//	Id int32
//	Name string `json:":rn"`
//	Address Address `json:":radr"`
//	Items []Item `json:":ritms"`
//}
//
//type NewRestaurant struct {
//	Name string `json:":rn"`
//	Address Address `json:":radr"`
//	Items []Item `json:":ritms"`
//}
//
//type RestaurantNoItems struct {
//	Id int32
//	Name string `json:"Name`
//	Address Address `json:"Address"`
//}
//
//type Order struct {
//	Id int32
//	Discount float32 `json:":od"`
//	Amount float32 `json:":oamt"`
//	PaymentMethod string `json:":opm"`
//	Rating int32 `json:":or"`
//	Duration int32 `json:":odtn"`
//	Cuisine string `json:":oc"`
//	Time int32 `json:":otm"`
//	Verified bool `json:":ov"`
//	Customer OrderCustomer `json:":octmr"`
//	Restaurant RestaurantNoItems `json:":ortrnt"`
//	Items []Item `json:":oitms"`
//}
//
//type OrderNoId struct {
//	Discount float32 `json:":od"`
//	Amount float32 `json:":oamt"`
//	PaymentMethod string `json:":opm"`
//	Rating int32 `json:":or"`
//	Duration int32 `json:":odtn"`
//	Cuisine string `json:":oc"`
//	Time int32 `json:":otm"`
//	Verified bool `json:":ov"`
//	Customer OrderCustomer `json:":octmr"`
//	Restaurant RestaurantNoItems `json:":ortrnt"`
//	Items []Item `json:":oitms"`
//}
//
//type OrderIp struct {
//	Id               int32
//	Discount float32
//	Amount float32
//	PaymentMethod    string
//	Rating, Duration int32
//	Cuisine string
//	Time int32
//	Verified bool
//	CustomerId int32
//	RestaurantId int32
//	Items []int32
//}

//type ItemIp struct {
//	RestaurantId int32
//	Name string
//	ItemId int32
//	Cuisine string
//	Discount float32
//	Amount float32
//}

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
	Address Address
	Items []Item
}

type RestaurantNoId struct {
	Name string `json:":n"`
	Address Address `json:":a"`
	Items []Item `json:":itms"`
}

type NewRestaurant struct {
	Name string
	Address Address
	Items []Item
}

type RestaurantNoItems struct {
	Id int32
	Name string
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


//type OrderNoId struct {
//	Discount float32
//	Amount float32
//	PaymentMethod string
//	Rating int32
//	Duration int32
//	Cuisine string
//	Time int32
//	Verified bool
//	Customer OrderCustomer
//	Restaurant RestaurantNoItems
//	Items []Item
//}

//type OrderIp struct {
//	Id               int32
//	Discount float32
//	Amount float32
//	PaymentMethod    string
//	Rating, Duration int32
//	Cuisine string
//	Time int32
//	Verified bool
//	CustomerId int32
//	RestaurantId int32
//	Items []int32
//}
