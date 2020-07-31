package http

import (
    stdctx "context"
    "github.com/jyotishp/go-orders/pkg/db"
    "github.com/jyotishp/go-orders/pkg/models"
    pb "github.com/jyotishp/go-orders/pkg/proto"
)

const ordersTableName = "Orders"

type OrdersServer struct {
}

func (s OrdersServer) GetOrder(ctx stdctx.Context, id *pb.OrderId) (*pb.Order, error) {
    order, err := db.GetOrder(ordersTableName, id.OrderId)
    if err != nil {
        return nil, err
    }
    return orderToPb(order), nil
}

func (s OrdersServer) PostOrder(ctx stdctx.Context, order *pb.CreateOrder) (*pb.Order, error) {
    ipOrder := models.Order{
        Restaurant: models.Restaurant{
            Id: order.RestaurantId,
        },
        Customer: models.Customer{
            Id: order.CustomerId,
        },
        Discount: order.Discount,
        Amount: order.Amount,
        PaymentMethod: order.PaymentMethod,
        Rating: order.Rating,
        Duration: order.Duration,
        Cuisine: order.Cuisine,
        Time: order.Time,
        Items: pbToItems(order.Items.Items),
    }

    newOrder, err := db.InsertOrder(ordersTableName, ipOrder)
    if err != nil {
        return nil, err
    }

    return orderToPb(newOrder), nil
}

func (s OrdersServer) PutOrder(ctx stdctx.Context, order *pb.UpdateOrder) (*pb.Order, error) {
    panic("implement me")
}

func (s OrdersServer) DeleteOrder(ctx stdctx.Context, id *pb.OrderId) (*pb.Empty, error) {
    panic("implement me")
}
