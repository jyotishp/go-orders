package db

import (
	"github.com/jyotishp/go-orders/pkg/models"
)

type dbItem struct {
	IName string `json:":in"`
	Cuisine string `json:":ic"`
	Discount float32 `json:":idc"`
	Amount float32 `json:":iamt"`
}

type dbAddress struct {
	Line1 string `json:":l1"`
	Line2 string `json:":l2"`
	City string `json:":cty"`
	State string `json:":st"`
}

type dbCustomer struct {
	CName string `json:":cn"`
	Address dbAddress `json:":cadr"`
}

type dbRestaurant struct {
	RName string `json:":rn"`
	Address dbAddress `json:":radr"`
	Items []dbItem `json:":ritms"`
}

type dbOrder struct {
	Discount float32 `json:":od"`
	Amount float32 `json:":oamt"`
	PaymentMethod string `json:":opm"`
	Rating int32 `json:":or"`
	OrderDuration int32 `json:":odtn"`
	Cuisine string `json:":oc"`
	Time int32 `json:":otm"`
	Verified bool `json:":ov"`
	Customer dbCustomer `json:":octmr"`
	Restaurant dbRestaurant `json:":ortrnt"`
	Items []dbItem `json:":oitms"`
}

func itemsMap(items []models.Item) []dbItem {
	dbitems := make([]dbItem, 0)
	for _, item := range items {
		dbitems = append(dbitems, dbItem{
			IName: item.Name,
			Cuisine: item.Cuisine,
			Discount: item.Discount,
			Amount: item.Amount,
		})
	}
	return dbitems
}

func restaurantMap(restaurant models.Restaurant) dbRestaurant {
	return dbRestaurant{
		RName: restaurant.Name,
		Address: dbAddress{
			Line1: restaurant.Address.Line1,
			Line2: restaurant.Address.Line2,
			City:  restaurant.Address.City,
			State: restaurant.Address.State,
		},
		Items: itemsMap(restaurant.Items),
	}
}

func restaurantNoItemsMap(restaurant models.Restaurant) dbRestaurant {
	return dbRestaurant{
		RName: restaurant.Name,
		Address: dbAddress{
			Line1: restaurant.Address.Line1,
			Line2: restaurant.Address.Line2,
			City:  restaurant.Address.City,
			State: restaurant.Address.State,
		},
	}
}

func addressMap(address models.Address) dbAddress {
	return dbAddress{
		Line1: address.Line1,
		Line2: address.Line2,
		City: address.City,
		State: address.State,
	}
}

func customerMap(customer models.Customer) dbCustomer {
	return dbCustomer{
		CName: customer.Name,
		Address: addressMap(customer.Address),
	}
}

func orderMap(order models.Order) dbOrder {
	return dbOrder{
		Discount: order.Discount,
		Amount: order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating: order.Rating,
		OrderDuration: order.Duration,
		Cuisine: order.Cuisine,
		Time: order.Time,
		Verified: order.Verified,
		Customer: customerMap(order.Customer),
		Restaurant: restaurantNoItemsMap(order.Restaurant),
		Items: itemsMap(order.Items),
	}

}