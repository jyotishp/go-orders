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

const tableName = "Customers"

func (s *CustomerServer) GetCustomer(ctx stdctx.Context, id *pb.CustomerId) (*pb.Customer, error) {
	customer, err := db.GetCustomer(tableName, id.CustomerId)
	if err != nil {
		return &pb.Customer{}, err
	}
	return &pb.Customer{
		Name: customer.Name,
		Id: customer.Id,
		Address: &pb.Address{
			Line1: customer.Address.Line1,
			Line2: customer.Address.Line2,
			City: customer.Address.City,
			State: customer.Address.State,
		},
	}, nil
}

func (s *CustomerServer) PostCustomer(ctx stdctx.Context, customer *pb.CreateCustomer) (*pb.Customer, error) {
	createCustomer := models.Customer{
		Name: customer.Name,
		Address: models.Address{
			Line1: customer.Address.Line1,
			Line2: customer.Address.Line2,
			City: customer.Address.City,
			State: customer.Address.State,
		},
	}
	newCustomer, err := db.CreateCustomer(tableName, createCustomer)
	if err != nil {
		return &pb.Customer{}, nil
	}
	return &pb.Customer{
		Id: newCustomer.Id,
		Name: newCustomer.Name,
		Address: &pb.Address{
			Line1: newCustomer.Address.Line1,
			Line2: newCustomer.Address.Line2,
			City: newCustomer.Address.City,
			State: newCustomer.Address.State,
		},
	}, nil
}

func (s *CustomerServer) PutCustomer(ctx stdctx.Context, customer *pb.UpdateCustomer) (*pb.Customer, error) {
	return &pb.Customer{}, nil
}

func (s *CustomerServer) DeleteCustomer(ctx stdctx.Context, id *pb.CustomerId) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *CustomerServer) ListCustomers(ctx context.Context, empty *pb.Empty) (*pb.CustomerList, error) {
	return &pb.CustomerList{}, nil
}
