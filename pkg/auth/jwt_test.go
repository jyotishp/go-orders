package auth

import (
	"testing"
	"time"
)

func TestNewJwtAuthenticator(t *testing.T) {
	secret := "test_secret"
	ttl := time.Minute
	authentictor := NewJwtAuthenticator(secret, ttl)

	if authentictor.secret != secret {
		t.Errorf("JWT secret is different from the one set")
	}

	if authentictor.ttl.String() != ttl.String() {
		t.Errorf("TTL is different from the one set")
	}
}

func TestCreateToken(t *testing.T) {
	jwt := NewJwtAuthenticator("test_secret", time.Minute)
	_, err := jwt.CreateToken("test_user")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestValidateToken(t *testing.T) {
	username := "test_user"
	jwt := NewJwtAuthenticator("test_secret", time.Second)
	token, _ := jwt.CreateToken(username)

	// Success
	claims, err := jwt.ValidateToken(token)
	if err != nil {
		t.Errorf("could not validate token: %v", err)
	}
	if claims.Username != username {
		t.Errorf("username mismatch in claims")
	}

	// Invalid token
	_, err = jwt.ValidateToken("random_string")
	if err == nil {
		t.Errorf("accepted random string instead of token")
	}

	// Expire token
	t.Logf("Waiting for the token to expire")
	time.Sleep(2 * time.Second)
	claims, err = jwt.ValidateToken(token)
	if err == nil {
		t.Errorf("Token did not expire: %v", err)
	}
}
