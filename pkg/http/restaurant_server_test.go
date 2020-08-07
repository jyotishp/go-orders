package http_test

import (
	"context"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"google.golang.org/grpc"
	"testing"
)

var rest_id int32
var rest_name string
var item_id int32

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
	res, err := client.PostRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PostRestaurant : %v ", err)
	} else{
		rest_id = res.Id
		rest_name = res.Name
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
		RestaurantId: rest_id,
		Item:&pb.CreateItemParams{
			Name:     "Pizza",
			Amount:   500,
			Discount: 10,
			Cuisine:  "Indian",
		},
	}
	res, err := client.PostItem(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PostItem : %v ", err)
	}else {
		item_id = res.Id
	}
}

func TestGetRestaurant(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRestaurantsClient(conn)
	req := &pb.RestaurantId{
		RestaurantId: rest_id,
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
		RestaurantName: rest_name,
	}
	_, err = client.GetRestaurantName(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling GetRestaurantName : %v ", err)
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
		RestaurantId: rest_id,
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
		ItemId:item_id,
		RestaurantId: rest_id,
	}
	_, err = client.GetItem(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling GetItem : %v ", err)
	}else{
		//fmt.Println(res)
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
		RestaurantId: rest_id,
		Restaurant:   &pb.UpdateRestaurantParams{
			Name: "Mughlai",
			Address: &pb.Address{},
			Items: nil,
		},
	}
	_, err = client.PutRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PutRestaurant : %v ", err)
	}else{
		rest_name = "Mughlai"
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
		ItemId:item_id,
		RestaurantId: rest_id,
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
			ItemId:       item_id,
			RestaurantId: rest_id,
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


	req := &pb.RestaurantId{RestaurantId: rest_id}
	_, err = client.DeleteRestaurant(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling DeleteRestaurant : %v ", err)
	}
}


