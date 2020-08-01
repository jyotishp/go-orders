package interceptors

import (
    "context"
    "github.com/jyotishp/go-orders/pkg/auth"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
    "log"
    "strings"
    "time"
)

// Intercepts the requests for authentication
type AuthInterceptor struct {
    Authenticator *auth.JwtAuthenticator
}

// Create instance of AuthMiddleware
func NewAuthInterceptor(secret string, ttl time.Duration) *AuthInterceptor {
    return &AuthInterceptor{
        Authenticator: auth.NewJwtAuthenticator(secret, ttl),
    }
}

// Unary Server Interceptor
func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
        ) (interface{}, error) {
        err := a.Authenticate(ctx, info.FullMethod)
        if err != nil {
            return nil, err
        }
        return handler(ctx, req)
    }
}

// Authenticate requests
func (a *AuthInterceptor) Authenticate(ctx context.Context, method string) error {
    if method == "/protos.Authentication/Login" {
        return nil
    }

    meta, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return status.Errorf(codes.Unauthenticated, "invalid metadata")
    }
    values := meta.Get("authorization")
    if len(values) == 0 {
        return status.Errorf(codes.Unauthenticated, "invalid authorization token")
    }

    accessKey := strings.Split(values[0], " ")
    if len(accessKey) != 2 {
        return status.Errorf(codes.Unauthenticated, "token should be of the form 'Bearer <token>'")
    }
    claims, err := a.Authenticator.ValidateToken(accessKey[1])
    if err != nil {
        return status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
    }
    log.Printf("authenticated user: %v", claims.Username)
    return nil
}