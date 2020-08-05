package http

//
//func TestCreateTable(t *testing.T) {
//	const address = "localhost:50051"
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("failed to dial: %v", err)
//	}
//	defer conn.Close()
//	client := pb.NewUtilsClient(conn)
//	req := &pb.Table{TableName: "Testing_table1"}
//	_, err = client.CreateTable(context.Background(),req)
//	if err != nil {
//		t.Fatalf("Error While calling CreateTable : %v ", err)
//	}
//}
//
//
//func TestDeleteTable(t *testing.T) {
//	const address = "localhost:50051"
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("failed to dial: %v", err)
//	}
//	defer conn.Close()
//	client := pb.NewUtilsClient(conn)
//	req := &pb.Table{TableName: "Testing_table1"}
//	_, err = client.DeleteTable(context.Background(),req)
//	if err != nil {
//		t.Fatalf("Error While calling DeleteTable : %v ", err)
//	}
//}
//
//

