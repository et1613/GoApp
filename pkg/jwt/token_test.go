package jwt

import (
	"testing"
	"time"
)

func mustManager(t *testing.T) *TokenManager {
	t.Helper()
	tm, err := NewTokenManager("0123456789abcdefghijklmnopqrstuvwxyz!@#%$^&*()", 15*time.Minute, 24*time.Hour)
	if err != nil {
		t.Fatalf("NewTokenManager error: %v", err)
	}
	return tm
}

func TestGenerateAndValidateTokens(t *testing.T) {
	tm := mustManager(t)

	access, refresh, err := tm.GenerateTokens("user-123")
	if err != nil {
		t.Fatalf("GenerateTokens error: %v", err)
	}
	if access == "" || refresh == "" {
		t.Fatalf("expected non-empty tokens")
	}

	ac, err := tm.ValidateToken(access)
	if err != nil {
		t.Fatalf("ValidateToken(access) error: %v", err)
	}
	if ac.Type != TokenTypeAccess {
		t.Fatalf("expected access token type, got %v", ac.Type)
	}
	if ac.Subject != "user-123" {
		t.Fatalf("expected sub=user-123, got %s", ac.Subject)
	}

	rc, err := tm.ValidateToken(refresh)
	if err != nil {
		t.Fatalf("ValidateToken(refresh) error: %v", err)
	}
	if rc.Type != TokenTypeRefresh {
		t.Fatalf("expected refresh token type, got %v", rc.Type)
	}
	if rc.Subject != "user-123" {
		t.Fatalf("expected sub=user-123, got %s", rc.Subject)
	}
}

func TestValidateToken_Invalid(t *testing.T) {
	tm := mustManager(t)

	if _, err := tm.ValidateToken("not-a-jwt"); err == nil {
		t.Fatalf("expected error for invalid token")
	}
}

func TestValidateToken_WrongSecret(t *testing.T) {
	tm1 := mustManager(t)
	tm2, err := NewTokenManager("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#%$^&*()", 15*time.Minute, 24*time.Hour)
	if err != nil {
		t.Fatalf("NewTokenManager(2) error: %v", err)
	}

	access, _, err := tm1.GenerateTokens("user-xyz")
	if err != nil {
		t.Fatalf("GenerateTokens error: %v", err)
	}

	if _, err := tm2.ValidateToken(access); err == nil {
		t.Fatalf("expected signature validation error with wrong secret")
	}
}
