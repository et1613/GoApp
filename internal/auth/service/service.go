package service

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"crypto/sha256"
	"encoding/hex"

	"github.com/dykethecreator/GoApp/internal/auth/repository"
	"github.com/dykethecreator/GoApp/pkg/domain"
	"github.com/dykethecreator/GoApp/pkg/jwt"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type AuthService struct {
	twilioClient     *twilio.RestClient
	verifyServiceSID string
	userRepo         repository.UserRepository
	deviceRepo       repository.DeviceRepository
	tokenManager     *jwt.TokenManager
	devMode          bool
}

func NewAuthService(userRepo repository.UserRepository, deviceRepo repository.DeviceRepository) *AuthService {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	verifyServiceSID := os.Getenv("TWILIO_VERIFY_SERVICE_SID")
	devMode := os.Getenv("AUTH_DEV_MODE") == "true" || os.Getenv("AUTH_DEV_MODE") == "1"

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	tokenManager, err := jwt.NewTokenManager(jwtSecret, 15*time.Minute, 7*24*time.Hour)
	if err != nil {
		log.Fatalf("Failed to create token manager: %v", err)
	}

	var client *twilio.RestClient
	if !devMode {
		if accountSid == "" || authToken == "" || verifyServiceSID == "" {
			log.Fatal("Twilio environment variables not set (set AUTH_DEV_MODE=true to bypass in development)")
		}
		client = twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		})
	} else {
		log.Printf("[DEV MODE] Twilio OTP is bypassed. Use OTP code 123456 for verification.")
	}

	return &AuthService{
		twilioClient:     client,
		verifyServiceSID: verifyServiceSID,
		userRepo:         userRepo,
		deviceRepo:       deviceRepo,
		tokenManager:     tokenManager,
		devMode:          devMode,
	}
}

func (s *AuthService) SendOTP(phoneNumber string) (string, error) {
	if s.devMode {
		log.Printf("[DEV MODE] SendOTP bypassed for %s. Returning 'sent'.", phoneNumber)
		return "sent", nil
	}
	params := &verify.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := s.twilioClient.VerifyV2.CreateVerification(s.verifyServiceSID, params)
	if err != nil {
		log.Printf("Failed to send OTP via Twilio: %v\n", err)
		return "", err
	}

	log.Printf("OTP sent successfully. Status: %s, SID: %s\n", *resp.Status, *resp.Sid)
	return *resp.Status, nil
}

func (s *AuthService) VerifyOTP(ctx context.Context, phoneNumber, code, deviceID string) (*domain.User, string, string, error) {
	// 1. Verify code (Twilio or Dev)
	if s.devMode {
		if code != "123456" {
			return nil, "", "", errors.New("invalid OTP code in dev mode; expected 123456")
		}
		log.Printf("[DEV MODE] OTP accepted for %s", phoneNumber)
	} else {
		checkParams := &verify.CreateVerificationCheckParams{}
		checkParams.SetTo(phoneNumber)
		checkParams.SetCode(code)

		resp, err := s.twilioClient.VerifyV2.CreateVerificationCheck(s.verifyServiceSID, checkParams)
		if err != nil {
			log.Printf("Failed to verify OTP via Twilio: %v\n", err)
			return nil, "", "", err
		}

		if resp.Status == nil || *resp.Status != "approved" {
			log.Printf("OTP verification failed for %s. Status: %s\n", phoneNumber, *resp.Status)
			return nil, "", "", errors.New("OTP verification failed or code is incorrect")
		}

		log.Printf("OTP verification successful for %s\n", phoneNumber)
	}

	// 2. Check if user exists in the database
	user, err := s.userRepo.FindByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		log.Printf("Error finding user by phone number: %v", err)
		return nil, "", "", err
	}

	// 3. If user does not exist, create a new one
	if user == nil {
		log.Printf("User with phone number %s not found. Creating a new user.", phoneNumber)
		newUser := &domain.User{
			PhoneNumber: phoneNumber,
		}
		user, err = s.userRepo.CreateUser(ctx, newUser)
		if err != nil {
			log.Printf("Error creating new user: %v", err)
			return nil, "", "", err
		}
		log.Printf("New user created with ID: %s", user.ID)
	} else {
		log.Printf("Found existing user with ID: %s", user.ID)
	}

	// 4. Generate tokens for the user
	accessToken, refreshToken, err := s.tokenManager.GenerateTokens(user.ID.String())
	if err != nil {
		log.Printf("Error generating tokens for user %s: %v", user.ID, err)
		return nil, "", "", err
	}

	// Debug: Validate and log token types to confirm mapping
	if claimsA, errA := s.tokenManager.ValidateToken(accessToken); errA == nil {
		log.Printf("Access token claims: type=%s sub=%s", claimsA.Type, claimsA.Subject)
	} else {
		log.Printf("Access token validate error: %v", errA)
	}
	if claimsR, errR := s.tokenManager.ValidateToken(refreshToken); errR == nil {
		log.Printf("Refresh token claims: type=%s sub=%s", claimsR.Type, claimsR.Subject)
	} else {
		log.Printf("Refresh token validate error: %v", errR)
	}

	log.Printf("Generated tokens for user %s (access len=%d, refresh len=%d)", user.ID, len(accessToken), len(refreshToken))

	// 5. Persist refresh token hash for revocation checks
	// Use device_id from request for better tracking
	if s.deviceRepo != nil {
		hash := hashRefreshToken(refreshToken)
		deviceName := deviceID
		if deviceName == "" {
			deviceName = "unknown"
		}
		dev := &domain.UserDevice{
			UserID:           user.ID,
			RefreshTokenHash: hash,
			DeviceName:       deviceName,
			DeviceType:       "mobile", // Could be enhanced with actual device type
			LastLoginAt:      time.Now(),
		}
		if err := s.deviceRepo.UpsertDevice(ctx, dev); err != nil {
			log.Printf("Warning: failed to upsert user device for user %s: %v", user.ID, err)
		} else {
			log.Printf("Device %s registered for user %s", deviceName, user.ID)
		}
	}

	return user, accessToken, refreshToken, nil
}

// RefreshToken validates a refresh token and issues a new pair of access and refresh tokens.
func (s *AuthService) RefreshToken(ctx context.Context, refreshTokenString string) (string, string, error) {
	// 1. Validate the refresh token
	claims, err := s.tokenManager.ValidateToken(refreshTokenString)
	if err != nil {
		return "", "", err // Returns ErrInvalidToken
	}

	// 2. Ensure it's a refresh token
	if claims.Type != jwt.TokenTypeRefresh {
		return "", "", errors.New("provided token is not a refresh token")
	}

	// 3. Verify refresh token hash exists and is not revoked; capture device for rotation
	var currentDev *domain.UserDevice
	if s.deviceRepo != nil {
		hash := hashRefreshToken(refreshTokenString)
		dev, derr := s.deviceRepo.FindActiveByUserAndHash(ctx, claims.Subject, hash)
		if derr != nil {
			return "", "", derr
		}
		if dev == nil {
			return "", "", jwt.ErrInvalidToken
		}
		currentDev = dev
	}

	// 4. Check if the user exists
	userID := claims.Subject
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		log.Printf("Error finding user by ID during token refresh: %v", err)
		return "", "", err
	}
	if user == nil {
		return "", "", errors.New("user not found for the given token")
	}

	// 5. Issue new access and refresh tokens (rotation)
	newAccessToken, newRefreshToken, err := s.tokenManager.GenerateTokens(userID)
	if err != nil {
		log.Printf("Error generating new tokens for user %s: %v", userID, err)
		return "", "", err
	}

	// 6. Persist the new refresh token hash and revoke the old one
	if s.deviceRepo != nil {
		// Revoke old device session to prevent reuse (rotate)
		if currentDev != nil {
			if rerr := s.deviceRepo.RevokeByID(ctx, currentDev.ID.String()); rerr != nil {
				log.Printf("Warning: failed to revoke old device %s for user %s: %v", currentDev.ID, userID, rerr)
			}
		}
		// Create a new device session record, copy metadata if available
		newDev := &domain.UserDevice{
			UserID:           user.ID,
			RefreshTokenHash: hashRefreshToken(newRefreshToken),
			DeviceName:       "unknown",
			DeviceType:       "unknown",
			LastLoginAt:      time.Now(),
		}
		if currentDev != nil {
			newDev.DeviceName = currentDev.DeviceName
			newDev.DeviceType = currentDev.DeviceType
			newDev.PushNotificationToken = currentDev.PushNotificationToken
		}
		if uerr := s.deviceRepo.UpsertDevice(ctx, newDev); uerr != nil {
			log.Printf("Warning: failed to upsert new device for user %s: %v", userID, uerr)
		}
	}

	return newAccessToken, newRefreshToken, nil
}

// ValidateAccessToken validates an access token and returns whether it's valid and the associated user ID.
// It does not return an error for invalid tokens; instead, it returns (false, ""). Errors are only for unexpected conditions.
func (s *AuthService) ValidateAccessToken(accessToken string) (bool, string) {
	claims, err := s.tokenManager.ValidateToken(accessToken)
	if err != nil {
		return false, ""
	}

	// Ensure it's an access token
	if claims.Type != jwt.TokenTypeAccess {
		return false, ""
	}

	return true, claims.Subject
}

// hashRefreshToken returns a hex-encoded SHA-256 hash of the refresh token string.
func hashRefreshToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

// RevokeByRefreshToken revokes the specific device session identified by the provided refresh token.
// If the token is invalid or not found, it returns ErrInvalidToken for security (no enumeration).
func (s *AuthService) RevokeByRefreshToken(ctx context.Context, refreshToken string) error {
	if s.deviceRepo == nil {
		return errors.New("device repository not configured")
	}
	claims, err := s.tokenManager.ValidateToken(refreshToken)
	if err != nil {
		return jwt.ErrInvalidToken
	}
	if claims.Type != jwt.TokenTypeRefresh {
		return jwt.ErrInvalidToken
	}
	hash := hashRefreshToken(refreshToken)
	dev, err := s.deviceRepo.FindActiveByUserAndHash(ctx, claims.Subject, hash)
	if err != nil {
		return err
	}
	if dev == nil {
		return jwt.ErrInvalidToken
	}
	return s.deviceRepo.RevokeByID(ctx, dev.ID.String())
}

// RevokeAllForAccessToken revokes all active device sessions for the user extracted from the access token.
// Returns ErrInvalidToken when the token is invalid.
func (s *AuthService) RevokeAllForAccessToken(ctx context.Context, accessToken string) error {
	if s.deviceRepo == nil {
		return errors.New("device repository not configured")
	}
	claims, err := s.tokenManager.ValidateToken(accessToken)
	if err != nil {
		return jwt.ErrInvalidToken
	}
	if claims.Type != jwt.TokenTypeAccess {
		return jwt.ErrInvalidToken
	}
	return s.deviceRepo.RevokeAllForUser(ctx, claims.Subject)
}
