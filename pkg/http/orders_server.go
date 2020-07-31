package http

import (
    stdctx "context"
    "github.com/jyotishp/go-orders/pkg/db"
    pb "github.com/jyotishp/go-orders/pkg/proto"
)

const ordersTableName = "Orders"

type OrdersServer struct {
}

func (s OrdersServer) GetOrder(ctx stdctx.Context, id *pb.OrderId) (*pb.Order, error) {
    order, err := db.GetOrder(ordersTableName, id.OrderId)
    if err != nil {
        return &pb.Order{}, err
    }
    return createOrder(order), nil
}

func (s OrdersServer) PostOrder(ctx stdctx.Context, order *pb.CreateOrder) (*pb.Order, error) {
    panic("implement me")
}

func (s OrdersServer) PutOrder(ctx stdctx.Context, order *pb.UpdateOrder) (*pb.Order, error) {
    panic("implement me")
}

func (s OrdersServer) DeleteOrder(ctx stdctx.Context, id *pb.OrderId) (*pb.Empty, error) {
    panic("implement me")
}
