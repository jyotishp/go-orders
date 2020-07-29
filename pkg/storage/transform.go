package storage

import (
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"log"
	"strconv"
)

func ParseInt(txt string) int32 {
	ret, err := strconv.ParseInt(txt, 10, 32)
	if err != nil {
		log.Fatalln("Failed to convert value to int")
	}
	return int32(ret)
}

func ParseFloat(txt string) float32 {
	ret, err := strconv.ParseFloat(txt, 32)
	if err != nil {
		log.Fatalln("Failed to convert value to float")
	}
	return float32(ret)
}

func NewOrder(data []string) *pb.Order {
	return &pb.Order{
		Id:          ParseInt(data[0]),
		Discount:    ParseFloat(data[1]),
		Amount:      ParseFloat(data[2]),
		PaymentMode: data[3],
		Rating:      ParseInt(data[4]),
		Duration:    ParseInt(data[5]),
		Cuisine:     data[6],
		Time:        ParseInt(data[7]),
	}
}

func NewCustomer(data []string) *pb.Customer {
	return &pb.Customer{
		Id:    ParseInt(data[11]),
		Name:  data[12],
		State: data[10],
	}
}

func NewRestaurant(data []string) *pb.Restaurant {
	return &pb.Restaurant{
		Id:    ParseInt(data[8]),
		Name:  data[9],
		State: data[10],
	}
}
