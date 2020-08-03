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
    return orderToPb(order), err
}

func (s OrdersServer) PostOrder(ctx stdctx.Context, order *pb.CreateOrder) (*pb.Order, error) {
    newOrder, err := db.InsertOrder(ordersTableName, pbToCreateOrder(order))
    return orderToPb(newOrder), err
}

func (s OrdersServer) PutOrder(ctx stdctx.Context, order *pb.UpdateOrder) (*pb.Order, error) {
    newOrder, err := db.UpdateOrder(ordersTableName, pbToUpdateOrder(order))
    return orderToPb(newOrder), err
}

func (s OrdersServer) DeleteOrder(ctx stdctx.Context, id *pb.OrderId) (*pb.Empty, error) {
    err := db.DeleteKey(ordersTableName, id.OrderId)
    return &pb.Empty{}, err
}
