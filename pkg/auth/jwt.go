package auth

import (
    "errors"
    "time"
    "github.com/dgrijalva/jwt-go"
)

// UserClaims stores the username and TTL of the token
type UserClaims struct {
    jwt.StandardClaims
    Username string
}

// JWT Authenticator
type JwtAuthenticator struct {
    secret string
    ttl time.Duration
}

// Creates a new instance of JWT authenticator
func NewJwtAuthenticator(secret string, ttl time.Duration) *JwtAuthenticator {
    return &JwtAuthenticator{
        secret: secret,
        ttl:    ttl,
    }
}

// Generates the token with UserClaims
func (a *JwtAuthenticator) CreateToken(username string) (string, error) {
    claims := &UserClaims{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(a.ttl).Unix(),
        },
        Username:       username,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
    return token.SignedString([]byte(a.secret))
}

// Validate JWT token
func (a *JwtAuthenticator) ValidateToken(accessToken string) (*UserClaims, error) {
    token, err := jwt.ParseWithClaims(
        accessToken,
        &UserClaims{},
        func(token *jwt.Token) (interface{}, error) {
            _, ok := token.Method.(*jwt.SigningMethodHMAC)
            if !ok {
                return nil, errors.New("invalid signing method")
            }

            return []byte(a.secret), nil
        },
    )
    if err != nil {
         return nil, errors.New("invalid token")
    }

    claims, ok := token.Claims.(*UserClaims)
    if !ok {
        return nil, errors.New("invalid claims")
    }
    return claims, nil
}