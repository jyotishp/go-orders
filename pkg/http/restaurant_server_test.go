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
	pb.RegisterRestaurantsServer(s,&RestaurantsServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func TestGetRestaurant(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.RestaurantId{
		RestaurantId: 437812006,
	}
	_, err = client.GetRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling GetRestaurant : %v ", err)
	}
}

func TestGetRestaurantName(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.RestaurantName{
		RestaurantName: "Motimahal",
	}
	_, err = client.GetRestaurantName(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling GetRestaurantName : %v ", err)
	}
}

func TestPostRestaurant(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.CreateRestaurant{
		Name:    "xyz",
		Address: nil,
		Items:   nil,
	}
	_, err = client.PostRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PostRestaurant : %v ", err)
	}
}

func TestPutRestaurant(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.UpdateRestaurant{
		RestaurantId: 1291847688,
		Restaurant:   nil,
	}
	_, err = client.PutRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PutRestaurant : %v ", err)
	}
}

func TestDeleteRestaurant(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.RestaurantId{RestaurantId: 0}
	_, err = client.DeleteRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling DeleteRestaurant : %v ", err)
	}
}

func TestListItems(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.ItemsFilter{
		RestaurantId: 437812006,
		Min:          float32(1),
		Max:          float32(500),
	}
	_, err = client.ListItems(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling ListItems : %v ", err)
	}
}

func TestGetItem(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.ItemId{
		ItemId:-965301954,
		RestaurantId: -965507624,
	}
	_, err = client.GetItem(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling GetItem : %v ", err)
	}
}


func TestPostItem(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.CreateItem{
		RestaurantId: 1291847688,
		Item:&pb.CreateItemParams{
			Name:     "Pizza",
			Amount:   500,
			Discount: 10,
			Cuisine:  "Indian",
		},
	}
	_, err = client.PostItem(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PostItem : %v ", err)
	}
}

func TestPutItem(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.UpdateItem{
		ItemId:-965301954,
		RestaurantId: -965507624,
		Item:&pb.CreateItemParams{
			Name:     "Pizza",
			Amount:   500,
			Discount: 10,
			Cuisine:  "Indian",
		},
	}
	_, err = client.PutItem(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PutItem : %v ", err)
	}
}


func TestDeleteItem(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(),)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.ItemId{
		ItemId:       1,
		RestaurantId: 1,
	}
	_, err = client.DeleteItem(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling DeleteItem : %v ", err)
	}
}

