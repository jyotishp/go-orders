package http

import (
	"github.com/jyotishp/go-orders/pkg/models"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

func customerToPb(customer models.Customer) *pb.Customer  {
	return &pb.Customer{
		Name: customer.Name,
		Id:   customer.Id,
		Address: addressToPb(customer.Address),
	}
}

func restaurantToPb(restaurant models.Restaurant) *pb.Restaurant {
	return &pb.Restaurant{
		Name: restaurant.Name,
		Id: restaurant.Id,
		Address: addressToPb(restaurant.Address),
		Items: itemsToPb(restaurant.Items),
	}
}

func restaurantListToPb(restaurantList []models.Restaurant) *pb.RestaurantList {
	op := make([]*pb.Restaurant, 0)
	for _, restaurant := range restaurantList {
		op = append(op, restaurantToPb(restaurant))
	}
	return &pb.RestaurantList{
		Restaurants: op,
	}
}

func createRestaurantNoItem(restaurant models.RestaurantNoItems) *pb.RestaurantNoItems {
	return &pb.RestaurantNoItems{
		Name: restaurant.Name,
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
		Duration: order.Duration,
		Cuisine: order.Cuisine,
		Time: order.Time,
		Verified: order.Verified,
		Customer: customerToPb(order.Customer),
		Restaurant: createRestaurantNoItem(order.Restaurant),
		Items: itemsToPb(order.Items),
	}
}

func itemToPb(item models.Item) *pb.Item {
	return &pb.Item{
		Id: item.Id,
		Name: item.Name,
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


func itemListToPb(items []models.Item) *pb.ItemList {
	return &pb.ItemList{
		Items: itemsToPb(items),
	}
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
		Name: item.Name,
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

func pbToCreateOrder(order *pb.CreateOrder) models.OrderIp  {
	return models.OrderIp{
		RestaurantId: order.RestaurantId,
		CustomerId: order.CustomerId,
		Discount: order.Discount,
		Amount: order.Amount,
		PaymentMethod: order.PaymentMethod,
		Rating: order.Rating,
		Duration: order.Duration,
		Cuisine: order.Cuisine,
		Time: order.Time,
		Items: order.Items,
	}
}

func pbToFilter(filter *pb.ItemsFilter) models.ItemFilter {
	return models.ItemFilter{
		RestaurantId: filter.RestaurantId,
		Min: filter.Min,
		Max: filter.Max,
	}
}

func pbToCreateItem(item *pb.CreateItemParams) models.Item {
	return models.Item{
		Name: item.Name,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

func pbToCreateItems(items []*pb.CreateItemParams) []models.Item {
	op := make([]models.Item, 0)
	for _, item := range items {
		op = append(op, pbToCreateItem(item))
	}
	return op
}
