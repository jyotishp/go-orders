package db

import (
    pb "github.com/jyotishp/go-orders/pkg/proto"
    "github.com/tamerh/jsparser"
)

type RestaurantHandler struct {
    reader *jsparser.JsonParser
}

func NewRestaurantHandler(outdir string) *RestaurantHandler {
    handler := &RestaurantHandler{}
    handler.reader = JsonHandle(outdir + "/restaurants.json", "orders")
    return handler
}

func (h *RestaurantHandler) GetRestaurant(id int) *pb.Restaurant {
    return &pb.Restaurant{}
}