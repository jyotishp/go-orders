package db

import (
    pb "github.com/jyotishp/go-orders/pkg/proto"
    "github.com/tamerh/jsparser"
)

type CustomerHandler struct {
    reader *jsparser.JsonParser
}

func NewCustomerHandler(outdir string) *CustomerHandler {
    handler := &CustomerHandler{}
    handler.reader = JsonHandle(outdir + "/customers.json", "orders")
    return handler
}

func GetOrder(id int) *pb.Customer {
    return &pb.Customer{}
}