package auth

import (
	"testing"
	"time"
)

func TestJWT_GenerateAndValidate(t *testing.T) {
	secret := "test-secret-key"
	userID := "user-123"
	email := "test@example.com"

	t.Run("Valid token succeeds validation", func(t *testing.T) {
		token, err := GenerateToken(userID, email, secret, 1*time.Hour)
		if err != nil {
			t.Fatalf("failed to generate token: %v", err)
		}

		claims, err := ValidateToken(token, secret)
		if err != nil {
			t.Fatalf("failed to validate token: %v", err)
		}

		if claims.UserID != userID {
			t.Errorf("expected UserID %s, got %s", userID, claims.UserID)
		}
		if claims.Email != email {
			t.Errorf("expected Email %s, got %s", email, claims.Email)
		}
	})

	t.Run("Invalid secret fails validation", func(t *testing.T) {
		token, err := GenerateToken(userID, email, secret, 1*time.Hour)
		if err != nil {
			t.Fatalf("failed to generate token: %v", err)
		}

		_, err = ValidateToken(token, "wrong-secret")
		if err == nil {
			t.Error("expected error for wrong secret, got nil")
		}
	})
}
