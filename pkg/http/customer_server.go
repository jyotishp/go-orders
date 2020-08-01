package http

import (
	stdctx "context"
	"github.com/jyotishp/go-orders/pkg/db"
	"github.com/jyotishp/go-orders/pkg/models"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"golang.org/x/net/context"
)

type CustomerServer struct {
}

const customerTableName = "Customers"

func (s *CustomerServer) GetCustomer(ctx stdctx.Context, id *pb.CustomerId) (*pb.Customer, error) {
	customer, err := db.GetCustomer(customerTableName, id.CustomerId)
	if err != nil {
		return &pb.Customer{}, err
	}
	return customerToPb(customer), nil
}

func (s *CustomerServer) PostCustomer(ctx stdctx.Context, customer *pb.CreateCustomer) (*pb.Customer, error) {
	ipCustomer := models.Customer{
		Name: customer.Name,
		Address: pbToAddress(customer.Address),
	}
	newCustomer, err := db.InsertCustomer(customerTableName, ipCustomer)
	if err != nil {
		return &pb.Customer{}, err
	}
	return customerToPb(newCustomer), nil
}

func (s *CustomerServer) PutCustomer(ctx stdctx.Context, customer *pb.UpdateCustomer) (*pb.Customer, error) {
	ipCustomer := models.Customer{
		Id: customer.CustomerId,
		Name: customer.Customer.Name,
		Address: pbToAddress(customer.Customer.Address),
	}

	newCustomer, err := db.UpdateCustomer(customerTableName, ipCustomer)
	if err != nil {
		return &pb.Customer{}, err
	}
	return customerToPb(newCustomer), nil

}

func (s *CustomerServer) DeleteCustomer(ctx stdctx.Context, id *pb.CustomerId) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *CustomerServer) ListCustomers(ctx context.Context, empty *pb.Empty) (*pb.CustomerList, error) {
	return &pb.CustomerList{}, nil
}
