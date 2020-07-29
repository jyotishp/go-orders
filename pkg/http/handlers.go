package http

import (
    "context"
    "google.golang.org/grpc"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    pb "github.com/jyotishp/go-orders/pkg/proto"
    "log"
    "net"
    "net/http"
)

func StartGRPC() {
    lis, err := net.Listen("tcp", "localhost:5566")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterAnalysisServer(grpcServer, &Server{})
    log.Println("gRPC server ready...")
    grpcServer.Serve(lis)
}

func StartHTTP() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    conn, err := grpc.Dial("localhost:5566", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to dial: %v", err)
    }
    defer conn.Close()

    rmux := runtime.NewServeMux()
    client := pb.NewAnalysisClient(conn)
    err = pb.RegisterAnalysisHandlerClient(ctx, rmux, client)
    if err != nil {
       log.Fatal(err)
    }
    log.Println("Registered with gRPC...")

    mux := http.NewServeMux()
    mux.Handle("/", rmux)
    mux.HandleFunc("/swagger.json", serveSwagger)
    fs := http.FileServer(http.Dir("swagger-ui"))
    mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))

    err = http.ListenAndServe("localhost:8080", mux)
    if err != nil {
        log.Fatal(err)
    }
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "swagger-ui/analysis.swagger.json")
}