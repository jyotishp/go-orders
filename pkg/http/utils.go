package http

import (
	"github.com/jyotishp/go-orders/pkg/models"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

func customerToPb(customer models.Customer) *pb.Customer  {
	return &pb.Customer{
		Name: customer.CustomerName,
		Id:   customer.Id,
		Address: addressToPb(customer.Address),
	}
}

func restaurantToPb(restaurant models.Restaurant) *pb.Restaurant {
	return &pb.Restaurant{
		Name: restaurant.RestaurantName,
		Id: restaurant.Id,
		Address: addressToPb(restaurant.Address),
		Items: itemsToPb(restaurant.Items),
	}
}

func createRestaurantNoItem(restaurant models.Restaurant) *pb.Restaurant {
	return &pb.Restaurant{
		Name: restaurant.RestaurantName,
		Id: restaurant.Id,
		Address: addressToPb(restaurant.Address),
	}
}

func addressToPb(address models.Address) *pb.Address {
	return &pb.Address{
		Line1: address.Line1,
		Line2: address.Line2,
		City:  address.City,
		State: address.State,
	}
}

func orderToPb(order models.Order) *pb.Order  {
	return &pb.Order{
		Id: order.Id,
		Discount: order.Discount,
		Amount: order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating: order.Rating,
		Duration: order.OrderDuration,
		Cuisine: order.Cuisine,
		Time: order.OrderTime,
		Verified: order.Verified,
		Customer: customerToPb(order.Customer),
		Restaurant: createRestaurantNoItem(order.Restaurant),
		Items: itemsToPb(order.OrderItems),
	}
}

func itemToPb(item models.Item) *pb.Item {
	return &pb.Item{
		Id: item.Id,
		Name: item.ItemName,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func itemsToPb(items []models.Item) []*pb.Item {
	itemList := make([]*pb.Item, 0)
	for _, item := range items {
		itemList = append(itemList, itemToPb(item))
	}
	return itemList
}

func pbToAddress(address *pb.Address) models.Address  {
	return models.Address{
		Line1: address.Line1,
		Line2: address.Line2,
		City: address.City,
		State: address.State,
	}
}

func pbToItem(item *pb.Item)  models.Item {
	return models.Item{
		Id: item.Id,
		ItemName: item.Name,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func pbToItems(items []*pb.Item) []models.Item {
	ip := make([]models.Item, 0)
	for _, item := range items {
		ip = append(ip, pbToItem(item))
	}
	return ip
}

func pbToCreateOrder(order *pb.CreateOrder) models.Order  {
	return models.Order{
		Restaurant: models.Restaurant{
			Id: order.RestaurantId,
		},
		Customer: models.Customer{
			Id: order.CustomerId,
		},
		Discount: order.Discount,
		Amount: order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating: order.Rating,
		OrderDuration: order.Duration,
		Cuisine: order.Cuisine,
		OrderTime: order.Time,
		OrderItems: pbToItems(order.Items),
	}
}