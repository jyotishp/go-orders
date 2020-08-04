package http

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jyotishp/go-orders/pkg/interceptors"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

func StartGRPC(jwtSecret string, jwtTtl time.Duration) {
	lis, err := net.Listen("tcp", "0.0.0.0:5566")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//grpcServer := grpc.NewServer()

	authMiddleware := interceptors.NewAuthInterceptor(jwtSecret, jwtTtl)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(authMiddleware.Unary(), grpc_prometheus.UnaryServerInterceptor),
	))

	authServer := NewAuthServer("admin", "admin", jwtSecret, jwtTtl)
	pb.RegisterAuthenticationServer(grpcServer, authServer)

	pb.RegisterCustomersServer(grpcServer, &CustomerServer{})
	pb.RegisterUtilsServer(grpcServer, &UtilsServer{})
	pb.RegisterOrdersServer(grpcServer, &OrdersServer{})
	pb.RegisterAnalysisServer(grpcServer, &AnalysisServer{})
	pb.RegisterRestaurantsServer(grpcServer, &RestaurantsServer{})

	grpc_prometheus.Register(grpcServer)
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

	conn, err := grpc.Dial(
		"localhost:5566",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	rmux := runtime.NewServeMux()

	customersClient := pb.NewCustomersClient(conn)
	err = pb.RegisterCustomersHandlerClient(ctx, rmux, customersClient)
	ClientErr(err)

	utilsClient := pb.NewUtilsClient(conn)
	err = pb.RegisterUtilsHandlerClient(ctx, rmux, utilsClient)
	ClientErr(err)

	authClient := pb.NewAuthenticationClient(conn)
	err = pb.RegisterAuthenticationHandlerClient(ctx, rmux, authClient)
	ClientErr(err)

	ordersClient := pb.NewOrdersClient(conn)
	err = pb.RegisterOrdersHandlerClient(ctx, rmux, ordersClient)
	ClientErr(err)

	analysisClient := pb.NewAnalysisClient(conn)
	err = pb.RegisterAnalysisHandlerClient(ctx, rmux, analysisClient)
	ClientErr(err)

	restaurantsClient := pb.NewRestaurantsClient(conn)
	err = pb.RegisterRestaurantsHandlerClient(ctx, rmux, restaurantsClient)
	ClientErr(err)

	log.Println("Registered with gRPC...")

	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/swagger.json", serveSwagger)
	fs := http.FileServer(http.Dir("swagger-ui"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))

	err = http.ListenAndServe("0.0.0.0:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger-ui/app.swagger.json")
}
