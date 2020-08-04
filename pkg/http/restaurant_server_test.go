package http

import (
	"context"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"google.golang.org/grpc"
	"testing"
)


//func init() {
//	lis = bufconn.Listen(bufSize)
//	s := grpc.NewServer()
//	pb.RegisterRestaurantsServer(s,&RestaurantsServer{})
//
//	go func() {
//		if err := s.Serve(lis); err != nil {
//			log.Fatalf("Server exited with error: %v", err)
//		}
//	}()
//}

func TestGetRestaurant(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.RestaurantName{
		RestaurantName: "BBQ",
	}
	_, err = client.GetRestaurantName(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling GetRestaurantName : %v ", err)
	}
}

func TestPostRestaurant(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.CreateRestaurant{
		Name:    "xyz",
		Address: &pb.Address{},
		Items:   nil,
	}
	_, err = client.PostRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PostRestaurant : %v ", err)
	}
}

func TestPutRestaurant(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.UpdateRestaurant{
		RestaurantId: 437812006,
		Restaurant:   &pb.CreateRestaurant{
			Name: "ghhg",
			Address: &pb.Address{},
			Items: nil,
		},
	}
	_, err = client.PutRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PutRestaurant : %v ", err)
	}
}



func TestListItems(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.ItemId{
		ItemId:437947876,
		RestaurantId: 437812006,
	}
	_, err = client.GetItem(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling GetItem : %v ", err)
	}
}


func TestPostItem(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.CreateItem{
		RestaurantId: -876961234,
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
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.UpdateItem{
		ItemId:1292107078,
		RestaurantId: 1291847688,
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
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.ItemId{
		ItemId:       37872416,
		RestaurantId: 37680746,
	}
	_, err = client.DeleteItem(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling DeleteItem : %v ", err)
	}
}

func TestDeleteRestaurant(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.RestaurantId{RestaurantId: -876961234}
	_, err = client.DeleteRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling DeleteRestaurant : %v ", err)
	}
}

