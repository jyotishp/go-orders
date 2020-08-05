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
