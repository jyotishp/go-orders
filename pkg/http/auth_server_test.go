package http_test

import (
	"context"
	pb "github.com/jyotishp/go-orders/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	"time"
)

const (
	username = "test_admin"
	password = "test_password"
	secret   = "test_secret"
	ttl      = time.Hour
	bufSize  = 1024 * 1024
)

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterAuthenticationServer(s, NewAuthServer(username, password, secret, ttl))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestNewAuthServer(t *testing.T) {
	server := NewAuthServer(username, password, secret, ttl)

	if server.AdminUsername != username {
		t.Errorf("username mismatch")
	}
	if server.AdminPassword != password {
		t.Errorf("password mismatch")
	}
}

func TestAuthServer_Login(t *testing.T) {
	conn := OpenConnection(t)
	defer conn.Close()

	client := pb.NewAuthenticationClient(conn)

	// Success
	req := &pb.LoginCred{
		Username: username,
		Password: password,
	}
	_, err := client.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to login: %v ", err)
	}

	// Fail
	req.Password = "random_string"
	_, err = client.Login(context.Background(), req)
	if err == nil {
		t.Fatalf("failed to invalidate login: %v ", err)
	}
}

func OpenConnection(t *testing.T) *grpc.ClientConn {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithContextDialer(bufDialer),
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	return conn
}
