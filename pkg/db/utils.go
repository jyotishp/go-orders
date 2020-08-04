package db

import (
	"fmt"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

func createItem(restaurantId int32, item Item) ItemIp {
	return ItemIp{
		RestaurantId: restaurantId,
		ItemId: item.Id,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func updateItemMap(item ItemIp) ItemUp {
	return ItemUp{
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func extractCustomer(id int32) Customer {
	op, _ := GetCustomer("Customers", id)
	return op
}

func extractRestaurant(id int32) RestaurantNoItems {
	op, _ := GetRestaurant("Restaurants", id)
	return RestaurantNoItems{
		Name: op.Name,
		Id: op.Id,
		Address: op.Address,
	}
}

func extractItems(restaurantId int32, itemIds []int32) []Item {
	items := make([]Item, 0)
	for _, id := range itemIds {
		op, err := GetItem("Items", restaurantId, id)
		fmt.Println("ITEM -> ", op)
		if err != nil {
			fmt.Println("ERROR")
			panic(err.Error())
		}
		items = append(items, op)
	}
	return items
}

func insertItems(restaurantId int32, items []Item, updateRestaurants bool) []Item {
	op := make([]Item, 0)
	for _, item := range items {
		opItem, err := InsertItem("Items", restaurantId, item, updateRestaurants)
		if err != nil {
			fmt.Println("REACH4")
			printError(err)
			return op
		}
		op = append(op, opItem)
	}
	return op
}

func pbToAddress(address *pb.Address) Address  {
	return Address{
		Line1: address.Line1,
		Line2: address.Line2,
		City: address.City,
		State: address.State,
	}
}

func removeCustId(customer Customer) CustomerNoId {
	return CustomerNoId{
		Name: customer.Name,
		Address: customer.Address,
	}
}

func addressToPb(address Address) *pb.Address {
	return &pb.Address{
		Line1: address.Line1,
		Line2: address.Line2,
		City:  address.City,
		State: address.State,
	}
}

func toNormalOrder(order OrderIp) Order {
	return Order{
		Id:            order.Id,
		Discount:      order.Discount,
		Amount:        order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating:        order.Rating,
		Duration:      order.Duration,
		Cuisine:       order.Cuisine,
		Time:          order.Time,
		Verified:      order.Verified,
		Customer:      order.Customer,
		Restaurant:    order.Restaurant,
		Items:         order.Items,
	}
}

func toNormalRestaurant(restaurant RestaurantNoId) Restaurant {
	return Restaurant{
		Name:    restaurant.Name,
		Address: restaurant.Address,
		Items:   restaurant.Items,
	}
}

func removeOrderId(order OrderIp) OrderNoId  {
	return OrderNoId{
		Discount:      order.Discount,
		Amount:        order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating:        order.Rating,
		Duration:      order.Duration,
		Cuisine:       order.Cuisine,
		Time:          order.Time,
		Verified:      order.Verified,
		Customer:      order.Customer,
		Restaurant:    order.Restaurant,
		Items:         order.Items,
	}
}

func removeRestId(restaurant Restaurant) RestaurantNoId {
	return RestaurantNoId{
		Name:    restaurant.Name,
		Address: restaurant.Address,
		Items:   restaurant.Items,
	}
}