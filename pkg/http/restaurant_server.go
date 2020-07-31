package http

import (
	stdctx "context"
	"github.com/jyotishp/go-orders/pkg/db"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

type RestaurantsServer struct {
}

const restaurantsTableName = "Restaurants"

func (r RestaurantsServer) GetRestaurant(ctx stdctx.Context, id *pb.RestaurantId) (*pb.Restaurant, error) {
	restaurant, err := db.GetRestaurant(restaurantsTableName, id.RestaurantId)
	if err != nil {
		return nil, err
	}
	return restaurantToPb(restaurant), nil
}

func (r RestaurantsServer) GetRestaurantName(ctx stdctx.Context, name *pb.RestaurantName) (*pb.RestaurantList, error) {
	panic("implement me")
}

func (r RestaurantsServer) PostRestaurant(ctx stdctx.Context, restaurant *pb.CreateRestaurant) (*pb.Restaurant, error) {
	panic("implement me")
}

func (r RestaurantsServer) PutRestaurant(ctx stdctx.Context, restaurant *pb.UpdateRestaurant) (*pb.Restaurant, error) {
	panic("implement me")
}

func (r RestaurantsServer) DeleteRestaurant(ctx stdctx.Context, id *pb.RestaurantId) (*pb.Empty, error) {
	panic("implement me")
}

func (r RestaurantsServer) ListItems(ctx stdctx.Context, filter *pb.ItemsFilter) (*pb.ItemList, error) {
	panic("implement me")
}

func (r RestaurantsServer) GetItem(ctx stdctx.Context, id *pb.ItemId) (*pb.Item, error) {
	panic("implement me")
}

func (r RestaurantsServer) PostItem(ctx stdctx.Context, item *pb.CreateItem) (*pb.Item, error) {
	panic("implement me")
}

func (r RestaurantsServer) PutItem(ctx stdctx.Context, item *pb.UpdateItem) (*pb.Item, error) {
	panic("implement me")
}

func (r RestaurantsServer) DeleteItem(ctx stdctx.Context, id *pb.ItemId) (*pb.Empty, error) {
	panic("implement me")
}

