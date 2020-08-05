package http

import (
	stdctx "context"
	"github.com/jyotishp/go-orders/pkg/db"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

type RestaurantsServer struct {
}

const restaurantsTableName = "Restaurants"
const itemsTableName = "Items"

// GetRestaurant returns restaurant of given Id to the RestaurantsServer.
func (r RestaurantsServer) GetRestaurant(ctx stdctx.Context, id *pb.RestaurantId) (*pb.Restaurant, error) {
	restaurant, err := db.GetRestaurant(restaurantsTableName, id.RestaurantId)
	if err != nil {
		return &pb.Restaurant{}, err
	}
	return restaurantToPb(restaurant), nil
}

// GetRestaurantName returns list of restaurant of given name to the RestaurantsServer.
func (r RestaurantsServer) GetRestaurantName(ctx stdctx.Context, name *pb.RestaurantName) (*pb.RestaurantList, error) {
	restaurantList, err := db.GetRestaurantSGI(restaurantsTableName, name.RestaurantName)
	if err != nil {
		return &pb.RestaurantList{}, err
	}
	return restaurantListToPb(restaurantList), nil
}

// PostRestaurant inserts and returns restaurant of given params to the RestaurantsServer.
func (r RestaurantsServer) PostRestaurant(ctx stdctx.Context, restaurant *pb.CreateRestaurant) (*pb.Restaurant, error) {
	newRestaurant, err := db.InsertRestaurant(restaurantsTableName, pbToCreateRestaurant(restaurant))
	return  restaurantToPb(newRestaurant), err
}

// PutRestaurant updates and returns restaurant of given Id to the RestaurantsServer.
func (r RestaurantsServer) PutRestaurant(ctx stdctx.Context, restaurant *pb.UpdateRestaurant) (*pb.Restaurant, error) {
	newRestaurant, err := db.UpdateRestaurant(restaurantsTableName, pbToUpdateRestaurant(restaurant), true)
	return  restaurantToPb(newRestaurant), err
}

// DeleteRestaurant deletes restaurant of given Id.
func (r RestaurantsServer) DeleteRestaurant(ctx stdctx.Context, id *pb.RestaurantId) (*pb.Empty, error) {
	err := db.DeleteRestaurant(restaurantsTableName, id.RestaurantId)
	return &pb.Empty{}, err
}

// ListItems returns list of items of given restaurantId to the RestaurantsServer.
func (r RestaurantsServer) ListItems(ctx stdctx.Context, filter *pb.ItemsFilter) (*pb.ItemList, error) {
	items, err := db.GetAllItems(restaurantsTableName, filter)
	return itemListToPb(items), err
}

// GetItem returns item of given restaurantId and itemId to the RestaurantsServer.
func (r RestaurantsServer) GetItem(ctx stdctx.Context, id *pb.ItemId) (*pb.Item, error) {
	item, err := db.GetItem(itemsTableName, id.RestaurantId, id.ItemId)
	if err != nil {
		return &pb.Item{}, err
	}
	return itemToPb(item), nil
}

// PostItem inserts and returns item of given restaurantId and item params to the RestaurantsServer.
func (r RestaurantsServer) PostItem(ctx stdctx.Context, item *pb.CreateItem) (*pb.Item, error) {
	itemIp, err := db.InsertItem(itemsTableName, item.RestaurantId, pbToCreateItem(item.Item), true)
	if err != nil {
		return &pb.Item{}, err
	}
	return itemToPb(itemIp), nil
}

// PutItem returns item of given restaurantId and itemId to the RestaurantsServer.
func (r RestaurantsServer) PutItem(ctx stdctx.Context, item *pb.UpdateItem) (*pb.Item, error) {
	ip := pbToCreateItem(item.Item)
	ip.Id = item.ItemId
	itemIp, err := db.UpdateItem("Items", item.RestaurantId, ip)
	if err != nil {
		return &pb.Item{}, err
	}
	return itemToPb(itemIp), nil
}

// DeleteItem deletes item of given restaurantId and itemId.
func (r RestaurantsServer) DeleteItem(ctx stdctx.Context, id *pb.ItemId) (*pb.Empty, error) {
	err := db.DeleteItem(itemsTableName, id.RestaurantId, id.ItemId)
	return &pb.Empty{}, err
}

