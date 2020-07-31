package models

import (
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

type Model interface {}

func CreateAddressCreateCustomer(customer *pb.CreateCustomer) Address  {
	return Address{
		Line1: customer.Address.Line1,
		Line2: customer.Address.Line2,
		City: customer.Address.City,
		State: customer.Address.State,
	}
}