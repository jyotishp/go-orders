package http

import (
    "github.com/jyotishp/go-orders/pkg/db"
    pb "github.com/jyotishp/go-orders/pkg/proto"
    stdctx "context"
)

type AnalysisServer struct {
}

// TopRestaurants returns the top Restaurants to the AnalysisServer.
func (a AnalysisServer) TopRestaurants(ctx stdctx.Context, quantity *pb.Quantity) (*pb.RestaurantList, error) {
    restaurantList, err := db.GetTopRestaurants("Restaurants", quantity.Size)
    if err != nil {
        return &pb.RestaurantList{}, err
    }
    return restaurantListToPb(restaurantList), nil
}
// WorstRestaurants returns the worst Restaurants to the AnalysisServer.
func (a AnalysisServer) WorstRestaurants(ctx stdctx.Context, quantity *pb.Quantity) (*pb.RestaurantList, error) {
    restaurantList, err := db.GetWorstRestaurants("Restaurants", quantity.Size)
    if err != nil {
        return &pb.RestaurantList{}, err
    }
    return restaurantListToPb(restaurantList), nil
}
