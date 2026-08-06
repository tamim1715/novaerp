package auth

import (
	"testing"
	"time"
)

func TestRS256JWT_GenerateAndValidate(t *testing.T) {
	km, err := NewKeyManager("", "")
	if err != nil {
		t.Fatalf("failed to initialize KeyManager: %v", err)
	}

	userID := "3f393bbd-9790-4045-853e-26a2c732ee06"
	email := "test@novaerp.com"
	role := "admin"

	token, err := GenerateAccessToken(userID, email, role, km.PrivateKey, 15*time.Minute)
	if err != nil {
		t.Fatalf("failed to generate access token: %v", err)
	}

	claims, err := ValidateAccessToken(token, km.PublicKey)
	if err != nil {
		t.Fatalf("failed to validate access token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("expected UserID %s, got %s", userID, claims.UserID)
	}
	if claims.Email != email {
		t.Errorf("expected Email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("expected Role %s, got %s", role, claims.Role)
	}
}

func TestGenerateRefreshTokenString(t *testing.T) {
	token1, err := GenerateRefreshTokenString()
	if err != nil {
		t.Fatalf("failed to generate refresh token: %v", err)
	}
	token2, err := GenerateRefreshTokenString()
	if err != nil {
		t.Fatalf("failed to generate refresh token: %v", err)
	}

	if len(token1) != 64 {
		t.Errorf("expected hex string length 64, got %d", len(token1))
	}
	if token1 == token2 {
		t.Error("expected two unique random refresh tokens")
	}
}
