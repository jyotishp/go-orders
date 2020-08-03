package http

import (
	"context"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"testing"
	"log"
)

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterCustomersServer(s, &CustomerServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}


func TestListCustomer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersClient(conn)
	req := &pb.Empty{}
	_, err = client.ListCustomers(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling ListCustomers : %v ", err)
	}
}

func TestGetCustomer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersClient(conn)
	req := &pb.CustomerId{
		CustomerId:1,
	}
	_, err = client.GetCustomer(context.Background(),req )
	if err != nil {
		t.Fatalf("Error While calling GetCustomers : %v ", err)
	}
}



func TestPostCustomer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersClient(conn)
	req := &pb.CreateCustomer{
		Name:    "Swiggy",
		Address: nil,
	}
	_, err = client.PostCustomer(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PostCustomer : %v ", err)
	}
}


func TestPutCustomer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersClient(conn)
	req := &pb.UpdateCustomer{
		CustomerId: 1,
		Customer:   nil,
	}
	_, err = client.PutCustomer(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PutCustomer : %v ", err)
	}
}


func TestDeleteCustomer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersClient(conn)
	req := &pb.CustomerId{
		CustomerId: 1,
	}
	_, err = client.DeleteCustomer(context.Background(), req )
	if err != nil {
		t.Fatalf("Error While calling DeleteCustomers : %v ", err)
	}
}

