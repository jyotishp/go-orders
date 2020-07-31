package http

import (
    "context"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    pb "github.com/jyotishp/go-orders/pkg/proto"
    "google.golang.org/grpc"
    "log"
    "net"
    "net/http"
    "time"
)

func StartGRPC() {
    lis, err := net.Listen("tcp", "localhost:5566")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()

    authServer := NewAuthServer("admin", "admin", "secret", 3*time.Minute)
    pb.RegisterAuthenticationServer(grpcServer, authServer)

    pb.RegisterOrdersServer(grpcServer, &OrdersServer{})
    pb.RegisterAnalysisServer(grpcServer, &AnalysisServer{})

    log.Println("gRPC server ready...")
    grpcServer.Serve(lis)
}

func RegisterRpcClient(
    ctx *context.Context,
    mux *runtime.ServeMux,
    conn *grpc.ClientConn,
    client func(grpc.ClientConnInterface) interface{},
    registrar func(*context.Context, *runtime.ServeMux, interface{}) error,
    ) {
    c := client(conn)
    err := registrar(ctx, mux, c)
    if err != nil {
        log.Fatalf("unable to register client: %v", err)
    }
}

func ClientErr(err error) {
    if err != nil {
        log.Fatalf("failed to register client: %v", err)
    }
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

    authClient := pb.NewAuthenticationClient(conn)
    err = pb.RegisterAuthenticationHandlerClient(ctx, rmux, authClient)
    ClientErr(err)

    ordersClient := pb.NewOrdersClient(conn)
    err = pb.RegisterOrdersHandlerClient(ctx, rmux, ordersClient)
    ClientErr(err)

    analysisClient := pb.NewAnalysisClient(conn)
    err = pb.RegisterAnalysisHandlerClient(ctx, rmux, analysisClient)
    ClientErr(err)

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
    http.ServeFile(w, r, "swagger-ui/app.swagger.json")
}
