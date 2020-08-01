package http

import (
	stdctx "context"
	"github.com/jyotishp/go-orders/pkg/db"
	pb "github.com/jyotishp/go-orders/pkg/proto"
)

type UtilsServer struct {
}

func (s UtilsServer) CreateTable(ctx stdctx.Context, table *pb.Table) (*pb.Empty, error) {
	db.CreateTable(table.TableName)
	return &pb.Empty{}, nil
}

func (s UtilsServer) DeleteTable(ctx stdctx.Context, table *pb.Table) (*pb.Empty, error) {
	err := db.DeleteTable(table.TableName)
	return &pb.Empty{}, err
}
