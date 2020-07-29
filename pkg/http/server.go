package http

import (
    "context"
    pb "github.com/jyotishp/go-orders/pkg/proto"
)

type Server struct {
    pb.UnimplementedAnalysisServer
}

func (s *Server) GetOrder(context.Context, *pb.GetReq) (*pb.Order, error) {
    return &pb.Order{}, nil
}

func (s *Server) GetCustomer(context.Context, *pb.GetReq) (*pb.Customer, error) {
    return &pb.Customer{}, nil
}

func (s *Server) GetRestaurant(context.Context, *pb.GetReq) (*pb.Restaurant, error) {
    return &pb.Restaurant{}, nil
}

func (s *Server) TopRestaurants(context.Context, *pb.Length) (*pb.RestaurantList, error) {
    return &pb.RestaurantList{}, nil
}

func (s *Server) WorstRestaurants(context.Context, *pb.Length) (*pb.RestaurantList, error) {
    return &pb.RestaurantList{}, nil
}

func (s *Server) TopStateCuisines(context.Context, *pb.Length) (*pb.StateCuisines, error) {
    return &pb.StateCuisines{}, nil
}

func (s *Server) WorstStateCuisines(context.Context, *pb.Length) (*pb.StateCuisines, error) {
    return &pb.StateCuisines{}, nil
}

func (s *Server) CusineCustomerMap(context.Context, *pb.Length) (*pb.CuisineCustomers, error) {
    return &pb.CuisineCustomers{}, nil
}