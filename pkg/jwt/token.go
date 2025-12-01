package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid" // <-- NEW: Added for jti (JWT ID)
)

// ErrInvalidToken is a standard error for token validation failures.
var ErrInvalidToken = errors.New("invalid or expired token")

// TokenManager manages JWT creation and validation operations.
type TokenManager struct {
	secretKey            []byte
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

// --- NEW: Token Type Constants ---
// Prevents magic strings like "access", "refresh".
type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

// --- UPDATED: CustomClaims ---
// UserID is now carried within the 'sub' (Subject) standard.
type CustomClaims struct {
	Type TokenType `json:"type"` // 'access' or 'refresh'
	jwt.RegisteredClaims
}

// --- UPDATED: NewTokenManager ---
// Creates a new TokenManager instance.
func NewTokenManager(secret string, accessDuration, refreshDuration time.Duration) (*TokenManager, error) {
	// HS256 (SHA-256) expects a 256-bit (32 byte) key.
	// Enforcing minimum length for security.
	if len(secret) < 32 {
		return nil, errors.New("JWT secret key must be at least 32 bytes for security")
	}

	if accessDuration <= 0 || refreshDuration <= 0 {
		return nil, errors.New("token durations must be positive values")
	}

	return &TokenManager{
		secretKey:            []byte(secret),
		accessTokenDuration:  accessDuration,
		refreshTokenDuration: refreshDuration,
	}, nil
}

// generateToken, verilen claim'ler ile yeni bir token imzalar.
func (tm *TokenManager) generateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tm.secretKey)
}

// --- UPDATED: GenerateTokens ---
// Generates a new access and refresh token pair for a user ID
// containing standard claims (sub, jti, iat, etc.).
func (tm *TokenManager) GenerateTokens(userID string) (string, string, error) {
	now := time.Now()

	// Access Token Claims
	accessClaims := CustomClaims{
		Type: TokenTypeAccess,
		RegisteredClaims: jwt.RegisteredClaims{
			// Using 'sub' (Subject) standard for UserID
			Subject: userID,
			// 'jti' (JWT ID) makes the token unique, used for revocation
			ID: uuid.NewString(),
			// 'iss' (Issuer) who created the token
			Issuer: "my-auth-service",
			// 'aud' (Audience) who/which service the token is for
			Audience: jwt.ClaimStrings{"my-app-client"},
			// 'iat' (Issued At) when it was created
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(tm.accessTokenDuration)),
		},
	}
	accessTokenString, err := tm.generateToken(accessClaims)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	// Refresh Token Claims
	refreshClaims := CustomClaims{
		Type: TokenTypeRefresh,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userID,
			ID:      uuid.NewString(), // Unique ID for refresh token as well
			Issuer:  "my-auth-service",
			// Refresh token's audience is *only* the auth service
			Audience:  jwt.ClaimStrings{"my-auth-service"},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(tm.refreshTokenDuration)),
		},
	}
	refreshTokenString, err := tm.generateToken(refreshClaims)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return accessTokenString, refreshTokenString, nil
}

// --- UPDATED: ValidateToken ---
// Validates a token string and returns its claims.
// Makes error handling more specific.
func (tm *TokenManager) ValidateToken(tokenString string) (*CustomClaims, error) {

	// Parse the token according to our CustomClaims structure
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method (alg)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing algorithm: %v", token.Header["alg"])
		}
		// Return our secret key
		return tm.secretKey, nil
	})

	// Error Handling:
	if err != nil {
		// Check if the error is due to expiration
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrInvalidToken // Expired
		}

		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, ErrInvalidToken // Invalid signature
		}

		// For all other JWT-related errors (e.g., malformed format, not yet valid, etc.)
		// return our standard error. Not revealing details to the client is most secure.
		return nil, ErrInvalidToken
	}

	// Get the token and claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
