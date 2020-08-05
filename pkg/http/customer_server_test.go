package http

import (
	"context"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"testing"
)

//func init() {
//	lis = bufconn.Listen(bufSize)
//	s := grpc.NewServer()
//	pb.RegisterCustomersServer(s, &CustomerServer{})
//
//	go func() {
//		if err := s.Serve(lis); err != nil {
//			log.Fatalf("Server exited with error: %v", err)
//		}
//	}()
//}

var cust_id int32

var (
	port = ":50051"
)



func Server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCustomersServer(s, &CustomerServer{})
	pb.RegisterOrdersServer(s,&OrdersServer{})
	pb.RegisterRestaurantsServer(s,&RestaurantsServer{})
	pb.RegisterUtilsServer(s, &UtilsServer{})
	pb.RegisterAnalysisServer(s, &AnalysisServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestMain(m *testing.M) {
	go Server()
	os.Exit(m.Run())
}

func TestPostCustomer(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersClient(conn)
	req := &pb.CreateCustomer{
		Name:    "mnq",
		Address: &pb.Address{},
	}
	res, err := client.PostCustomer(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PostCustomer : %v ", err)
	}else {
		cust_id = res.Id
	}
}

func TestGetCustomer(t *testing.T) {
	id := cust_id
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersClient(conn)
	req := &pb.CustomerId{CustomerId: int32(id)}
	_, err = client.GetCustomer(context.Background(),req )
	if err != nil {
		t.Fatalf("Error While calling GetCustomers : %v ", err)
	}
}


func TestListCustomer(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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


func TestPutCustomer(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersClient(conn)
	req := &pb.UpdateCustomer{
		CustomerId: cust_id,
		Customer:&pb.CreateCustomer{
			Name: "pyz",
			Address: &pb.Address{},
		},
	}
	_, err = client.PutCustomer(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PutCustomer : %v ", err)
	}
}


func TestDeleteCustomer(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersClient(conn)
	req := &pb.CustomerId{
		CustomerId: cust_id,
	}
	_, err = client.DeleteCustomer(context.Background(), req )
	if err != nil {
		t.Fatalf("Error While calling DeleteCustomers : %v ", err)
	}
}

