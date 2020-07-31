package http

import (
    "context"
    "github.com/jyotishp/go-orders/pkg/auth"
    pb "github.com/jyotishp/go-orders/pkg/proto"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "time"
)

// Authentication server
type AuthServer struct {
    pb.UnimplementedAuthenticationServer
    AdminUsername string
    AdminPassword string
    Authenticator *auth.JwtAuthenticator
}

// Creates an instance of authentication server
func NewAuthServer(username, password, secret string, tokenTtl time.Duration) *AuthServer {
    return &AuthServer{
        AdminUsername:                     username,
        AdminPassword:                     password,
        Authenticator: auth.NewJwtAuthenticator(secret, tokenTtl),
    }
}

// Generate JWT token for a user given username and password
func (a *AuthServer) Login(ctx context.Context, creds *pb.LoginCred) (*pb.Token, error) {
    if creds.Username == a.AdminUsername && creds.Password == a.AdminPassword {
        token, err := a.Authenticator.CreateToken(creds.Username)
        if err != nil {
            return nil, status.Errorf(codes.Unknown, "failed to create JWT token")
        }
        return &pb.Token{Token: token}, nil
    }
    return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
}
