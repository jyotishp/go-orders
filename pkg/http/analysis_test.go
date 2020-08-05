package http

import (
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"testing"
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
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestMain(m *testing.M) {
	go Server()
	os.Exit(m.Run())
}

func TestAnalysisServer_TopRestaurants(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAnalysisClient(conn)
	req := &pb.Quantity{
		Size: 5,
	}
	_, err := client.TopRestaurants(context.Background(), req)
	if err != nil {
		t.Fatalf("Error in Testing Top Restaurants : #{err}")
	}
}

func TestAnalysisServer_WorstRestaurants(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAnalysisClient(conn)
	req := &pb.Quantity{
		Size: 5,
	}
	_, err := client.TopRestaurants(context.Background(), req)
	if err != nil {
		t.Fatalf("Error in Testing Top Restaurants : #{err}")
	}
}


