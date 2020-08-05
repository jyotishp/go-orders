package http_test

import (
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"testing"
)


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
	_, err = client.TopRestaurants(context.Background(), req)
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
	_, err = client.TopRestaurants(context.Background(), req)
	if err != nil {
		t.Fatalf("Error in Testing Top Restaurants : #{err}")
	}
}


