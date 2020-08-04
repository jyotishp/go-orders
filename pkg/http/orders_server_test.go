package http

import (
	"context"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"testing"
)


func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterOrdersServer(s, &OrdersServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func TestGetOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrdersClient(conn)
	req := &pb.OrderId{
		OrderId:int32(-10390236),
	}
	_, err = client.GetOrder(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling GetOrder : %v ", err)
	}
}

func TestPostOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrdersClient(conn)
	req := &pb.CreateOrder{
		RestaurantId:  0,
		CustomerId:    0,
		Discount:      0,
		Amount:        0,
		PaymentMethod: "",
		Rating:        0,
		Duration:      0,
		Cuisine:       "",
		Time:          0,
		Items:         nil,
	}
	_, err = client.PostOrder(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PostOrder : %v ", err)
	}
}


func TestPutOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrdersClient(conn)
	req := &pb.UpdateOrder{
		OrderId: 0,
		Order:   &pb.CreateOrder{
			RestaurantId:  0,
			CustomerId:    0,
			Discount:      0,
			Amount:        0,
			PaymentMethod: "",
			Rating:        0,
			Duration:      0,
			Cuisine:       "",
			Verified:      false,
			Time:          0,
			Items:         nil,
		},
	}
	_, err = client.PutOrder(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PutOrder : %v ", err)
	}
}


func TestDeleteOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrdersClient(conn)
	req := &pb.OrderId{
		OrderId: 0,
	}
	_, err = client.DeleteOrder(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling DeleteOrder : %v ", err)
	}
}

