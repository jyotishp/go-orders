package http

import (
    stdctx "context"
    "github.com/jyotishp/go-orders/pkg/db"
    pb "github.com/jyotishp/go-orders/pkg/proto"
)

const ordersTableName = "Orders"

type OrdersServer struct {
}

// GetOrder returns an order of given Id to the OrdersServer.
func (s OrdersServer) GetOrder(ctx stdctx.Context, id *pb.OrderId) (*pb.Order, error) {
    order, err := db.GetOrder(ordersTableName, id.OrderId)
    return orderToPb(order), err
}

// PostOrder inserts and returns an order of given params to the OrdersServer.
func (s OrdersServer) PostOrder(ctx stdctx.Context, order *pb.CreateOrder) (*pb.Order, error) {
    newOrder, err := db.InsertOrder(ordersTableName, pbToCreateOrder(order))
    return orderToPb(newOrder), err
}

// PutOrder updates returns an order of given Id to the OrdersServer.
func (s OrdersServer) PutOrder(ctx stdctx.Context, order *pb.UpdateOrder) (*pb.Order, error) {
    newOrder, err := db.UpdateOrder(ordersTableName, pbToUpdateOrder(order))
    return orderToPb(newOrder), err
}

// DeleteOrder deletes an order of given Id.
func (s OrdersServer) DeleteOrder(ctx stdctx.Context, id *pb.OrderId) (*pb.Empty, error) {
    err := db.DeleteKey(ordersTableName, id.OrderId)
    return &pb.Empty{}, err
}
