package db

import (
	"fmt"
	"github.com/jyotishp/go-orders/pkg/models"
)

type dbItemIp struct {
	RestaurantId int32
	Name string
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
	Id int32 `json:"Id"`
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
			Id: item.Id,
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

func restaurantNoItemsMap(restaurant models.RestaurantNoItems) dbOrderRestaurant {
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
		Restaurant: restaurantNoItemsMap(order.Restaurant),
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

func extractCustomer(id int32) models.Customer {
	op, _ := GetCustomer("Customers", id)
	return op
}

func extractRestaurant(id int32) models.RestaurantNoItems {
	op, _ := GetRestaurant("Restaurants", id)
	return models.RestaurantNoItems{
		Name: op.Name,
		Id: op.Id,
		Address: op.Address,
	}
}

func extractItems(restaurantId int32, itemIds []int32) []models.Item {
	items := make([]models.Item, 0)
	for _, id := range itemIds {
		op, err := GetItem("Items", restaurantId, id)
		fmt.Println("ITEM -> ", op)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, op)
	}
	return items
}

func buildOrder(order models.OrderIp) models.Order {
	return models.Order{
		Id: order.Id,
		Discount: order.Discount,
		Amount: order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating: order.Rating,
		Duration: order.Duration,
		Cuisine: order.Cuisine,
		Time: order.Time,
		Verified: order.Verified,
		Customer: extractCustomer(order.CustomerId),
		Restaurant: extractRestaurant(order.RestaurantId),
		Items: extractItems(order.RestaurantId, order.Items),
	}
}

func buildRestaurant(restaurant models.Restaurant) models.Restaurant {
	return models.Restaurant{
		Id: restaurant.Id,
		Name: restaurant.Name,
		Address: restaurant.Address,
	}
}

func insertItems(restaurantId int32, items []models.Item, updateRestaurants bool) []models.Item {
	op := make([]models.Item, 0)
	for _, item := range items {
		opItem, err := InsertItem("Items", restaurantId, item, updateRestaurants)
		if err != nil {
			printError(err)
			return op
		}
		op = append(op, opItem)
	}
	return op
}