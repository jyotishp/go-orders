package http

import (
    pb "github.com/jyotishp/go-orders/pkg/proto"
)

type Server struct {
    pb.UnimplementedAnalysisServer
}
