package http

import (
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"golang.org/x/net/context"
)

type CustomerServer struct {
	pb.UnimplementedCustomersServer
}

func (s *CustomerServer) ListCustomers(ctx context.Context, empty *pb.Empty) (*pb.CustomerList, error) {
	return &pb.CustomerList{}, nil
}
