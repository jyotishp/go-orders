package middlewares

import (
    "context"
    "github.com/jyotishp/go-orders/pkg/auth"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
    "log"
)

// Intercepts the requests for authentication
type AuthMiddleware struct {
    Authenticator *auth.JwtAuthenticator
}

// Unary Server Interceptor
func (a *AuthMiddleware) Unary() grpc.UnaryServerInterceptor {
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
        return handler(ctx, req), nil
    }
}

// Authenticate requests
func (a *AuthMiddleware) Authenticate(ctx context.Context, method string) error {
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

    accessKey := values[0]
    claims, err := a.Authenticator.ValidateToken(accessKey)
    if err != nil {
        return status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
    }
    log.Printf("authenticated user: %v", claims.Username)
    return nil
}