package http

import (
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

type RestaurantsServer struct {
	pb.UnimplementedRestaurantsServer
}

//const restaurantsTableName = "Restaurants"
//const itemsTableName = "Items"
//
//func (r RestaurantsServer) GetRestaurant(ctx stdctx.Context, id *pb.RestaurantId) (*pb.Restaurant, error) {
//	restaurant, err := db.GetRestaurant(restaurantsTableName, id.RestaurantId)
//	if err != nil {
//		return &pb.Restaurant{}, err
//	}
//	return restaurantToPb(restaurant), nil
//}
//
//func (r RestaurantsServer) GetRestaurantName(ctx stdctx.Context, name *pb.RestaurantName) (*pb.RestaurantList, error) {
//	restaurantList, err := db.GetRestaurantSGI(restaurantsTableName, name.RestaurantName)
//	if err != nil {
//		return &pb.RestaurantList{}, err
//	}
//	return restaurantListToPb(restaurantList), nil
//}
//
//func (r RestaurantsServer) PostRestaurant(ctx stdctx.Context, restaurant *pb.CreateRestaurant) (*pb.Restaurant, error) {
//	newRestaurant, err := db.InsertRestaurant(restaurantsTableName, pbToCreateRestaurant(restaurant))
//	return  restaurantToPb(newRestaurant), err
//}
//
//func (r RestaurantsServer) PutRestaurant(ctx stdctx.Context, restaurant *pb.UpdateRestaurant) (*pb.Restaurant, error) {
//	newRestaurant, err := db.UpdateRestaurant(restaurantsTableName, pbToUpdateRestaurant(restaurant), true)
//	return  restaurantToPb(newRestaurant), err
//}
//
//func (r RestaurantsServer) DeleteRestaurant(ctx stdctx.Context, id *pb.RestaurantId) (*pb.Empty, error) {
//	err := db.DeleteRestaurant(restaurantsTableName, id.RestaurantId)
//	return &pb.Empty{}, err
//}
//
//func (r RestaurantsServer) ListItems(ctx stdctx.Context, filter *pb.ItemsFilter) (*pb.ItemList, error) {
//	items, err := db.GetAllItems(restaurantsTableName, filter)
//	return itemListToPb(items), err
//}
//
//func (r RestaurantsServer) GetItem(ctx stdctx.Context, id *pb.ItemId) (*pb.Item, error) {
//	item, err := db.GetItem(itemsTableName, id.RestaurantId, id.ItemId)
//	if err != nil {
//		return &pb.Item{}, err
//	}
//	return itemToPb(item), nil
//}
//
//func (r RestaurantsServer) PostItem(ctx stdctx.Context, item *pb.CreateItem) (*pb.Item, error) {
//	itemIp, err := db.InsertItem(itemsTableName, item.RestaurantId, pbToCreateItem(item.Item), true)
//	if err != nil {
//		return &pb.Item{}, err
//	}
//	return itemToPb(itemIp), nil
//}
//
//func (r RestaurantsServer) PutItem(ctx stdctx.Context, item *pb.UpdateItem) (*pb.Item, error) {
//	ip := pbToCreateItem(item.Item)
//	ip.Id = item.ItemId
//	itemIp, err := db.UpdateItem("Items", item.RestaurantId, ip)
//	if err != nil {
//		return &pb.Item{}, err
//	}
//	return itemToPb(itemIp), nil
//}
//
//func (r RestaurantsServer) DeleteItem(ctx stdctx.Context, id *pb.ItemId) (*pb.Empty, error) {
//	err := db.DeleteItem(itemsTableName, id.RestaurantId, id.ItemId)
//	return &pb.Empty{}, err
//}

