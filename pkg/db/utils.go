package db

import (
	"github.com/jyotishp/go-orders/pkg/models"
)

type dbItemIp struct {
	RestaurantId int32
	ItemId int32
	Cuisine string
	Discount float32
	Amount float32
}

type dbItemUp struct {
	Cuisine string `json:":ic"`
	Discount float32 `json:":idc"`
	Amount float32 `json:":iamt"`
}

type dbCustomer struct {
	Name string `json:":cn"`
	Address dbAddress `json:":cadr"`
}

type dbAddress struct {
	Line1 string `json:"Line1"`
	Line2 string `json:"Line2"`
	City string `json:"City"`
	State string `json:"State"`
}

type dbItem struct {
	Name string `json:"Name"`
	Cuisine string `json:"Cuisine"`
	Discount float32 `json:"Discount"`
	Amount float32 `json:"Amount"`
}

type dbOrderCustomer struct {
	Name string `json:"Name"`
	Address dbAddress `json:"Address"`
}

type dbRestaurant struct {
	Name string `json:":rn"`
	Address dbAddress `json:":radr"`
	Items []dbItem `json:":ritms"`
}

type dbOrderRestaurant struct {
	Name string `json:"Name`
	Address dbAddress `json:"Address"`
}

type dbOrder struct {
	Discount float32 `json:":od"`
	Amount float32 `json:":oamt"`
	PaymentMethod string `json:":opm"`
	Rating int32 `json:":or"`
	Duration int32 `json:":odtn"`
	Cuisine string `json:":oc"`
	Time int32 `json:":otm"`
	Verified bool `json:":ov"`
	Customer dbOrderCustomer `json:":octmr"`
	Restaurant dbOrderRestaurant `json:":ortrnt"`
	Items []dbItem `json:":oitms"`
}

func itemsMap(items []models.Item) []dbItem {
	dbitems := make([]dbItem, 0)
	for _, item := range items {
		dbitems = append(dbitems, dbItem{
			Name: item.Name,
			Cuisine: item.Cuisine,
			Discount: item.Discount,
			Amount: item.Amount,
		})
	}
	return dbitems
}

func restaurantMap(restaurant models.Restaurant) dbRestaurant {
	return dbRestaurant{
		Name: restaurant.Name,
		Address: addressMap(restaurant.Address),
		Items: itemsMap(restaurant.Items),
	}
}

func orderRestaurantNoItemsMap(restaurant models.Restaurant) dbOrderRestaurant {
	return dbOrderRestaurant{
		Name: restaurant.Name,
		Address: addressMap(restaurant.Address),
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
		Name: customer.Name,
		Address: addressMap(customer.Address),
	}
}

func orderCustomerMap(customer models.Customer) dbOrderCustomer {
	return dbOrderCustomer{
		Name: customer.Name,
		Address: addressMap(customer.Address),
	}
}

func orderMap(order models.Order) dbOrder {
	return dbOrder{
		Discount: order.Discount,
		Amount: order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating: order.Rating,
		Duration: order.Duration,
		Cuisine: order.Cuisine,
		Time: order.Time,
		Verified: order.Verified,
		Customer: orderCustomerMap(order.Customer),
		Restaurant: orderRestaurantNoItemsMap(order.Restaurant),
		Items: itemsMap(order.Items),
	}
}

func createItem(restaurantId int32, item models.Item) dbItemIp {
	return dbItemIp{
		RestaurantId: restaurantId,
		ItemId: item.Id,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func updateItemMap(item dbItemIp) dbItemUp {
	return dbItemUp{
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func updateItemRest(restaurantId int32, item models.Item)   {
	
}