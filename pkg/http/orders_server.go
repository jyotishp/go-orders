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
    return orderToPb(order), nil
}

func (s OrdersServer) PostOrder(ctx stdctx.Context, order *pb.CreateOrder) (*pb.Order, error) {
    ipOrder := pbToCreateOrder(order)

    newOrder, err := db.InsertOrder(ordersTableName, ipOrder)
    if err != nil {
        return &pb.Order{}, err
    }

    return orderToPb(newOrder), nil
}

func (s OrdersServer) PutOrder(ctx stdctx.Context, order *pb.UpdateOrder) (*pb.Order, error) {
    ipOrder := pbToCreateOrder(order.Order)
    ipOrder.Id = order.OrderId

    newOrder, err := db.UpdateOrder(ordersTableName, ipOrder)
    if err != nil {
        return &pb.Order{}, err
    }

    return orderToPb(newOrder), nil
}

func (s OrdersServer) DeleteOrder(ctx stdctx.Context, id *pb.OrderId) (*pb.Empty, error) {
    err := db.DeleteKey(ordersTableName, id.OrderId)
    return &pb.Empty{}, err
}
