package http

import (
	"github.com/jyotishp/go-orders/pkg/models"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

func createCustomer(customer models.Customer) *pb.Customer  {
	return &pb.Customer{
		Name: customer.Name,
		Id:   customer.Id,
		Address: createAddress(customer.Address),
	}
}

func createRestaurant(restaurant models.Restaurant) *pb.Restaurant {
	return &pb.Restaurant{
		Name: restaurant.Name,
		Id: restaurant.Id,
		Address: createAddress(restaurant.Address),
		Items: createItems(restaurant.Items),
	}
}

func createRestaurantNoItem(restaurant models.Restaurant) *pb.Restaurant {
	return &pb.Restaurant{
		Name: restaurant.Name,
		Id: restaurant.Id,
		Address: createAddress(restaurant.Address),
	}
}

func createAddress(address models.Address) *pb.Address {
	return &pb.Address{
		Line1: address.Line1,
		Line2: address.Line2,
		City:  address.City,
		State: address.State,
	}
}

func createOrder(order models.Order) *pb.Order  {
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
		Customer: createCustomer(order.Customer),
		Restaurant: createRestaurantNoItem(order.Restaurant),
		Items: createItems(order.Items),
	}
}

func createItem(item models.Item) *pb.Item {
	return &pb.Item{
		Id: item.Id,
		Name: item.Name,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func createItems(items []models.Item) *pb.ItemList {
	itemList := &pb.ItemList{}
	for _, item := range items {
		itemList.Items = append(itemList.Items, createItem(item))
	}
	return itemList
}
