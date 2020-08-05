package db

import (
	"fmt"
)

// createItem returns Item Params for the given item.
func createItem(restaurantId int32, item Item) ItemIp {
	return ItemIp{
		RestaurantId: restaurantId,
		Name: item.Name,
		ItemId: item.Id,
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

// updateItemMap returns the Item Params for updating items.
func updateItemMap(item ItemIp) ItemUp {
	return ItemUp{
		Cuisine: item.Cuisine,
		Discount: item.Discount,
		Amount: item.Amount,
	}
}

// insertItems accepts an item list and inserts all of these into the Items table.
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

// removeCustId returns the Customer insert params.
func removeCustId(customer Customer) CustomerNoId {
	return CustomerNoId{
		Name: customer.Name,
		Address: customer.Address,
	}
}

// toNormalOrder adds order ID to the order.
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
// toNormalRestaurant adds Id to the Restaurant.
func toNormalRestaurant(restaurant RestaurantNoId) Restaurant {
	return Restaurant{
		Name:    restaurant.Name,
		Address: restaurant.Address,
		Items:   restaurant.Items,
	}
}

// removeOrderId removes order Id.
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

// removeRestId removes restaurant Id.
func removeRestId(restaurant Restaurant) RestaurantNoId {
	return RestaurantNoId{
		Name:    restaurant.Name,
		Address: restaurant.Address,
		Items:   restaurant.Items,
	}
}