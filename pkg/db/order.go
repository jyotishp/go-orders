package db

import (
	pb "github.com/jyotishp/go-orders/pkg/proto"
    "github.com/tamerh/jsparser"
)

type OrderHandler struct {
    reader *jsparser.JsonParser
}

func NewOrderHandler(outdir string) *OrderHandler {
    handler := &OrderHandler{}
    handler.reader = JsonHandle(outdir + "/orders.json", "orders")
    return handler
}

func (h *OrderHandler) GetOrder(id int) *pb.Order {
    return &pb.Order{}
}