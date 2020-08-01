package http

import (
	stdctx "context"
	"github.com/jyotishp/go-orders/pkg/db"
	"github.com/jyotishp/go-orders/pkg/models"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

type RestaurantsServer struct {
}

const restaurantsTableName = "Restaurants"
const itemsTableName = "Items"

func (r RestaurantsServer) GetRestaurant(ctx stdctx.Context, id *pb.RestaurantId) (*pb.Restaurant, error) {
	restaurant, err := db.GetRestaurant(restaurantsTableName, id.RestaurantId)
	if err != nil {
		return &pb.Restaurant{}, err
	}
	return restaurantToPb(restaurant), nil
}

func (r RestaurantsServer) GetRestaurantName(ctx stdctx.Context, name *pb.RestaurantName) (*pb.RestaurantList, error) {
	panic("implement me")
}

func (r RestaurantsServer) PostRestaurant(ctx stdctx.Context, restaurant *pb.CreateRestaurant) (*pb.Restaurant, error) {
	ipRestaurant := models.Restaurant{
		Name: restaurant.Name,
		Address: pbToAddress(restaurant.Address),
		Items: pbToItems(restaurant.Items),
	}
	newRestaurant, err := db.InsertRestaurant(restaurantsTableName, ipRestaurant)
	if err != nil {
		return &pb.Restaurant{}, err
	}

	return  restaurantToPb(newRestaurant), nil
}

func (r RestaurantsServer) PutRestaurant(ctx stdctx.Context, restaurant *pb.UpdateRestaurant) (*pb.Restaurant, error) {
	ipRestaurant := models.Restaurant{
		Name: restaurant.Restaurant.Name,
		Address: pbToAddress(restaurant.Restaurant.Address),
		Items: pbToItems(restaurant.Restaurant.Items),
	}
	ipRestaurant.Id = restaurant.RestaurantId
	newRestaurant, err := db.UpdateRestaurant(restaurantsTableName, ipRestaurant)
	if err != nil {
		return &pb.Restaurant{}, err
	}

	return  restaurantToPb(newRestaurant), nil
}

func (r RestaurantsServer) DeleteRestaurant(ctx stdctx.Context, id *pb.RestaurantId) (*pb.Empty, error) {
	err := db.DeleteKey(restaurantsTableName, id.RestaurantId)
	return &pb.Empty{}, err
}

func (r RestaurantsServer) ListItems(ctx stdctx.Context, filter *pb.ItemsFilter) (*pb.ItemList, error) {
	items, err := db.GetAllItems(restaurantsTableName, pbToFilter(filter))
	if err != nil {
		return &pb.ItemList{}, nil
	}
	return itemListToPb(items), nil
}

func (r RestaurantsServer) GetItem(ctx stdctx.Context, id *pb.ItemId) (*pb.Item, error) {
	item, err := db.GetItem(itemsTableName, id.RestaurantId, id.ItemId)
	if err != nil {
		return &pb.Item{}, err
	}
	return itemToPb(item), nil
}

func (r RestaurantsServer) PostItem(ctx stdctx.Context, item *pb.CreateItem) (*pb.Item, error) {
	itemIp, err := db.InsertItem(itemsTableName, item.RestaurantId, models.Item{
		Name: item.Item.Name,
		Cuisine: item.Item.Cuisine,
		Discount: item.Item.Discount,
		Amount: item.Item.Amount,
	})
	if err != nil {
		return &pb.Item{}, err
	}
	return itemToPb(itemIp), nil
}

func (r RestaurantsServer) PutItem(ctx stdctx.Context, item *pb.UpdateItem) (*pb.Item, error) {
	itemIp, err := db.UpdateItem("Items", item.RestaurantId, models.Item{
		Id: item.ItemId,
		Name: item.Item.Name,
		Cuisine: item.Item.Cuisine,
		Discount: item.Item.Discount,
		Amount: item.Item.Amount,
	})
	if err != nil {
		return &pb.Item{}, err
	}
	return itemToPb(itemIp), nil
}

func (r RestaurantsServer) DeleteItem(ctx stdctx.Context, id *pb.ItemId) (*pb.Empty, error) {
	panic("implement me")
}

