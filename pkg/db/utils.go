package db

import (
	"fmt"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

//func itemsMap(items []Item) []Item {
//	Items := make([]Item, 0)
//	for _, item := range items {
//		Items = append(Items, Item{
//			Id: item.Id,
//			Name: item.Name,
//			Cuisine: item.Cuisine,
//			Discount: item.Discount,
//			Amount: item.Amount,
//		})
//	}
//	return Items
//}

//func restaurantMap(restaurant Restaurant) Restaurant {
//	return Restaurant{
//		Name: restaurant.Name,
//		Address: addressMap(restaurant.Address),
//		Items: itemsMap(restaurant.Items),
//	}
//}
//
//func restaurantNoItemsMap(restaurant RestaurantNoItems) RestaurantNoItems {
//	return RestaurantNoItems{
//		Name: restaurant.Name,
//		Address: addressMap(restaurant.Address),
//	}
//}
//
//func addressMap(address Address) Address {
//	return Address{
//		Line1: address.Line1,
//		Line2: address.Line2,
//		City: address.City,
//		State: address.State,
//	}
//}
//
//func customerMap(customer Customer) Customer {
//	return Customer{
//		Name: customer.Name,
//		Address: addressMap(customer.Address),
//	}
//}
//
//func orderCustomerMap(customer Customer) Customer {
//	return Customer{
//		Id: customer.Id,
//		Name: customer.Name,
//		Address: addressMap(customer.Address),
//	}
//}

//func orderMap(order Order) Order {
//	return Order{
//		Discount: order.Discount,
//		Amount: order.Amount,
//		PaymentMethod: order.PaymentMethod,
//		Rating: order.Rating,
//		Duration: order.Duration,
//		Cuisine: order.Cuisine,
//		Time: order.Time,
//		Verified: order.Verified,
//		Customer: orderCustomerMap(order.Customer),
//		Restaurant: restaurantNoItemsMap(order.Restaurant),
//		Items: itemsMap(order.Items),
//	}
//}

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
//
//func buildOrder(order OrderIp) Order {
//	return Order{
//		Id: order.Id,
//		Discount: order.Discount,
//		Amount: order.Amount,
//		PaymentMethod: order.PaymentMethod,
//		Rating: order.Rating,
//		Duration: order.Duration,
//		Cuisine: order.Cuisine,
//		Time: order.Time,
//		Verified: order.Verified,
//		Customer: extractCustomer(order.CustomerId),
//		Restaurant: extractRestaurant(order.RestaurantId),
//		Items: extractItems(order.RestaurantId, order.Items),
//	}
//}
//
//func buildRestaurant(restaurant Restaurant) Restaurant {
//	return Restaurant{
//		Id: restaurant.Id,
//		Name: restaurant.Name,
//		Address: restaurant.Address,
//	}
//}

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

func pbToNewCustomerMap(customer *pb.CreateCustomer) Customer {
	return Customer{
		Name: customer.Name,
		Address: pbToAddress(customer.Address),
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

func pbToCreateOrder(order *pb.CreateOrder, orderId int32) Order  {
	return Order{
		Id: orderId,
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

func pbToUpdateOrder(order *pb.UpdateOrder) OrderIp {
	return OrderIp{
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

func pbToCreateRestaurant(restaurant *pb.CreateRestaurant) Restaurant {
	return Restaurant{
		Name: restaurant.Name,
		Address: pbToAddress(restaurant.Address),
	}
}

func restaurantToUpdatePb(restaurant Restaurant) *pb.UpdateRestaurant {
	return &pb.UpdateRestaurant{
		RestaurantId: restaurant.Id,
		Restaurant: &pb.CreateRestaurant{
			Name: restaurant.Name,
			Address: addressToPb(restaurant.Address),
		},
	}
}

func pbToCreateItem(item *pb.CreateItemParams) Item {
	return Item{
		Name: item.Name,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func pbToCreateItems(items []*pb.CreateItemParams) []Item {
	op := make([]Item, 0)
	for _, item := range items {
		op = append(op, pbToCreateItem(item))
	}
	return op
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