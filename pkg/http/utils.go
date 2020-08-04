package http

import (
	"fmt"
	"github.com/jyotishp/go-orders/pkg/db"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

func customerToPb(customer db.Customer) *pb.Customer  {
	return &pb.Customer{
		Name: customer.Name,
		Id:   customer.Id,
		Address: addressToPb(customer.Address),
	}
}

func restaurantToPb(restaurant db.Restaurant) *pb.Restaurant {
	return &pb.Restaurant{
		Name: restaurant.Name,
		Id: restaurant.Id,
		Address: addressToPb(restaurant.Address),
		Items: itemsToPb(restaurant.Items),
	}
}

func restaurantListToPb(restaurantList []db.Restaurant) *pb.RestaurantList {
	op := make([]*pb.Restaurant, 0)
	for _, restaurant := range restaurantList {
		op = append(op, restaurantToPb(restaurant))
	}
	return &pb.RestaurantList{
		Restaurants: op,
	}
}

func customerListToPb(customerList []db.Customer) *pb.CustomerList {
	op := make([]*pb.Customer, 0)
	for _, customer := range customerList {
		op = append(op, customerToPb(customer))
	}
	return &pb.CustomerList{
		Customer: op,
	}
}

func createRestaurantNoItem(restaurant db.RestaurantNoItems) *pb.RestaurantNoItems {
	return &pb.RestaurantNoItems{
		Name: restaurant.Name,
		Id: restaurant.Id,
		Address: addressToPb(restaurant.Address),
	}
}

func addressToPb(address db.Address) *pb.Address {
	return &pb.Address{
		Line1: address.Line1,
		Line2: address.Line2,
		City:  address.City,
		State: address.State,
	}
}

func orderToPb(order db.Order) *pb.Order  {
	return &pb.Order{
		Id: order.Id,
		Discount: order.Discount,
		Amount: order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating: order.Rating,
		Duration: order.Duration,
		Cuisine: order.Cuisine,
		Time: order.Time,
		Verified: order.Verified,
		Customer: customerToPb(order.Customer),
		Restaurant: createRestaurantNoItem(order.Restaurant),
		Items: itemsToPb(order.Items),
	}
}

func itemToPb(item db.Item) *pb.Item {
	return &pb.Item{
		Id: item.Id,
		Name: item.Name,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func itemsToPb(items []db.Item) []*pb.Item {
	itemList := make([]*pb.Item, 0)
	for _, item := range items {
		itemList = append(itemList, itemToPb(item))
	}
	return itemList
}


func itemListToPb(items []db.Item) *pb.ItemList {
	return &pb.ItemList{
		Items: itemsToPb(items),
	}
}

func pbToCreateItem(item *pb.CreateItemParams) db.Item {
	return db.Item{
		Name: item.Name,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func pbToCreateItems(items []*pb.CreateItemParams) []db.Item {
	op := make([]db.Item, 0)
	for _, item := range items {
		op = append(op, pbToCreateItem(item))
	}
	return op
}

func pbToAddress(address *pb.Address) db.Address  {
	return db.Address{
		Line1: address.Line1,
		Line2: address.Line2,
		City: address.City,
		State: address.State,
	}
}

func pbToCreateCustomer(customer *pb.CreateCustomer) db.Customer {
	return db.Customer{
		Name: customer.Name,
		Address: pbToAddress(customer.Address),
	}
}

func pbToUpdateCustomer(customer *pb.UpdateCustomer) db.Customer {
	return db.Customer{
		Id: customer.CustomerId,
		Name: customer.Customer.Name,
		Address: pbToAddress(customer.Customer.Address),
	}
}

func extractCustomer(id int32) db.Customer {
	op, _ := db.GetCustomer("Customers", id)
	return op
}

func extractRestaurant(id int32) db.RestaurantNoItems {
	op, _ := db.GetRestaurant("Restaurants", id)
	return db.RestaurantNoItems{
		Name: op.Name,
		Id: op.Id,
		Address: op.Address,
	}
}

func extractItems(restaurantId int32, itemIds []int32) []db.Item {
	items := make([]db.Item, 0)
	for _, id := range itemIds {
		op, err := db.GetItem("Items", restaurantId, id)
		fmt.Println("ITEM -> ", op)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, op)
	}
	return items
}

func pbToCreateOrder(order *pb.CreateOrder) db.Order  {
	return db.Order{
		Restaurant: extractRestaurant(order.RestaurantId),
		Customer: extractCustomer(order.CustomerId),
		Discount: order.Discount,
		Amount: order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating: order.Rating,
		Duration: order.Duration,
		Cuisine: order.Cuisine,
		Time: order.Time,
		Items: extractItems(order.RestaurantId, order.Items),
	}
}

func pbToUpdateOrder(order *pb.UpdateOrder) db.OrderIp {
	return db.OrderIp{
		Id: order.OrderId,
		Discount: order.Order.Discount,
		Amount: order.Order.Amount,
		PaymentMethod: order.Order.PaymentMethod,
		Rating: order.Order.Rating,
		Duration: order.Order.Duration,
		Cuisine: order.Order.Cuisine,
		Time: order.Order.Time,
		Verified: order.Order.Verified,
		Customer: extractCustomer(order.Order.CustomerId),
		Restaurant: extractRestaurant(order.Order.RestaurantId),
		Items: extractItems(order.Order.RestaurantId, order.Order.Items),
	}
}

func pbToCreateRestaurant(restaurant *pb.CreateRestaurant) db.Restaurant {
	return db.Restaurant{
		Name: restaurant.Name,
		Address: pbToAddress(restaurant.Address),
		Items: pbToCreateItems(restaurant.Items),
	}
}

func pbToUpdateRestaurant(restaurant *pb.UpdateRestaurant) db.Restaurant {
	return db.Restaurant{
		Id:      restaurant.RestaurantId,
		Name:    restaurant.Restaurant.Name,
		Address: pbToAddress(restaurant.Restaurant.Address),
		Items:   pbToCreateItems(restaurant.Restaurant.Items),
	}
}