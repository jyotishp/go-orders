package http

import pb "github.com/jyotishp/go-orders/pkg/proto"

type OrdersServer struct {
    pb.UnimplementedOrdersServer
}