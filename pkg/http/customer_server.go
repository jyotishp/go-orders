package http

import (
	stdctx "context"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/jyotishp/go-orders/pkg/db"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"golang.org/x/net/context"
)

type CustomersServer struct {
	pb.UnimplementedCustomersServer
	db dynamodbiface.DynamoDBAPI
}

func NewCustomersServer(svc dynamodbiface.DynamoDBAPI) *CustomersServer {
	return &CustomersServer{
		db: svc,
	}
}

func (s *CustomersServer) GetCustomer(ctx stdctx.Context, id *pb.CustomerId) (*pb.Customer, error) {
	return db.GetCustomer(s.db, id)
}

func (s *CustomersServer) PostCustomer(ctx stdctx.Context, customer *pb.CreateCustomer) (*pb.Customer, error) {
	return db.AddCustomer(s.db, customer)
}

func (s *CustomersServer) PutCustomer(ctx stdctx.Context, customer *pb.UpdateCustomer) (*pb.Customer, error) {
	return db.UpdateCustomer(s.db, customer)
}

func (s *CustomersServer) DeleteCustomer(ctx stdctx.Context, id *pb.CustomerId) (*pb.Empty, error) {
	return db.DeleteCustomer(s.db, id)
}

func (s *CustomersServer) ListCustomers(ctx context.Context, empty *pb.Empty) (*pb.CustomerList, error) {
	return db.ListCustomers(s.db)
}
