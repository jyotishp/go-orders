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
//	pb.RegisterOrdersServer(s,&OrdersServer{})
//
//	go func() {
//		if err := s.Serve(lis); err != nil {
//			log.Fatalf("Server exited with error: %v", err)
//		}
//	}()
//}

var order_id int32

func TestPostOrder(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
	res, err := client.PostOrder(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling PostOrder : %v ", err)
	} else{
		order_id = res.Id
	}
}



func TestGetOrder(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrdersClient(conn)
	req := &pb.OrderId{
		OrderId:int32(order_id),
	}
	_, err = client.GetOrder(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling GetOrder : %v ", err)
	}
}




func TestPutOrder(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrdersClient(conn)
	req := &pb.UpdateOrder{
		OrderId: order_id,
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
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrdersClient(conn)
	req := &pb.OrderId{
		OrderId: order_id,
	}
	_, err = client.DeleteOrder(context.Background(),req)
	if err != nil {
		t.Fatalf("Error While calling DeleteOrder : %v ", err)
	}
}

