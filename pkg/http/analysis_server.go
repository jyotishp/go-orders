package http

import (
	pb "github.com/jyotishp/go-orders/pkg/proto"
    stdctx "context"
)

type AnalysisServer struct {
}

func (a AnalysisServer) TopRestaurants(ctx stdctx.Context, quantity *pb.Quantity) (*pb.RestaurantList, error) {
    panic("implement me")
}

func (a AnalysisServer) WorstRestaurants(ctx stdctx.Context, quantity *pb.Quantity) (*pb.RestaurantList, error) {
    panic("implement me")
}

func (a AnalysisServer) TopStateCuisines(ctx stdctx.Context, quantity *pb.Quantity) (*pb.StateCuisines, error) {
    panic("implement me")
}

func (a AnalysisServer) WorstStateCuisines(ctx stdctx.Context, quantity *pb.Quantity) (*pb.StateCuisines, error) {
    panic("implement me")
}

func (a AnalysisServer) CuisineCustomerMap(ctx stdctx.Context, cuisine *pb.Cuisine) (*pb.CuisineCustomers, error) {
    panic("implement me")
}
