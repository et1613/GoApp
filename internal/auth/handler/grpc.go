package handler

import (
	"context"
	"errors"
	"log"

	"github.com/dykethecreator/GoApp/internal/auth/service"
	appjwt "github.com/dykethecreator/GoApp/pkg/jwt"
	"github.com/dykethecreator/GoApp/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	proto.UnimplementedAuthServiceServer
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(s *grpc.Server) {
	proto.RegisterAuthServiceServer(s, h)
}

func (h *AuthHandler) SendOTP(ctx context.Context, req *proto.SendOTPRequest) (*proto.SendOTPResponse, error) {
	log.Printf("Received SendOTP request for phone number: %s", req.PhoneNumber)

	status, err := h.service.SendOTP(req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &proto.SendOTPResponse{
		Message: "OTP sent successfully, status: " + status,
	}, nil
}

func (h *AuthHandler) VerifyOTP(ctx context.Context, req *proto.VerifyOTPRequest) (*proto.VerifyOTPResponse, error) {
	log.Printf("Received VerifyOTP request for phone number: %s, device_id: %s", req.PhoneNumber, req.DeviceId)

	user, accessToken, refreshToken, err := h.service.VerifyOTP(ctx, req.PhoneNumber, req.OtpCode, req.DeviceId)
	if err != nil {
		// Check for specific error types if needed, otherwise return a general internal error.
		// The service layer already logs the details.
		return nil, status.Errorf(codes.Internal, "failed to verify OTP: %v", err)
	}

	// Return user info along with tokens
	log.Printf("VerifyOTP success for %s (user_id=%s, access len=%d, refresh len=%d)", req.PhoneNumber, user.ID, len(accessToken), len(refreshToken))
	return &proto.VerifyOTPResponse{
		User: &proto.User{
			Id:                user.ID.String(),
			PhoneNumber:       user.PhoneNumber,
			DisplayName:       user.DisplayName,
			ProfilePictureUrl: user.ProfilePictureURL,
			AboutText:         user.AboutText,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Helper to convert *string to string
func stringPtrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (h *AuthHandler) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	log.Printf("ValidateToken called. Token length: %d, first 50 chars: %s", len(req.AccessToken), truncate(req.AccessToken, 50))
	valid, userID := h.service.ValidateAccessToken(req.AccessToken)
	log.Printf("ValidateToken result: valid=%v, user_id=%s (len=%d)", valid, userID, len(userID))
	// Do not treat invalid token as an RPC error; return is_valid=false
	return &proto.ValidateTokenResponse{
		IsValid: valid,
		UserId:  userID,
	}, nil
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func (h *AuthHandler) RefreshToken(ctx context.Context, req *proto.RefreshTokenRequest) (*proto.RefreshTokenResponse, error) {
	newAccess, newRefresh, err := h.service.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		// Map known token errors to Unauthenticated; others to Internal
		if errors.Is(err, appjwt.ErrInvalidToken) {
			return nil, status.Error(codes.Unauthenticated, "invalid or expired token")
		}
		return nil, status.Errorf(codes.Internal, "failed to refresh token: %v", err)
	}
	return &proto.RefreshTokenResponse{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}, nil
}

// RevokeCurrentDevice revokes the current device session identified by the provided refresh token.
func (h *AuthHandler) RevokeCurrentDevice(ctx context.Context, req *proto.RevokeCurrentDeviceRequest) (*proto.RevokeResponse, error) {
	if err := h.service.RevokeByRefreshToken(ctx, req.RefreshToken); err != nil {
		if errors.Is(err, appjwt.ErrInvalidToken) {
			return nil, status.Error(codes.Unauthenticated, "invalid or expired token")
		}
		return nil, status.Errorf(codes.Internal, "failed to revoke device: %v", err)
	}
	return &proto.RevokeResponse{Success: true}, nil
}

// LogoutAllDevices revokes all active device sessions for the user identified by the access token.
func (h *AuthHandler) LogoutAllDevices(ctx context.Context, req *proto.LogoutAllDevicesRequest) (*proto.RevokeResponse, error) {
	if err := h.service.RevokeAllForAccessToken(ctx, req.AccessToken); err != nil {
		if errors.Is(err, appjwt.ErrInvalidToken) {
			return nil, status.Error(codes.Unauthenticated, "invalid or expired token")
		}
		return nil, status.Errorf(codes.Internal, "failed to logout all devices: %v", err)
	}
	return &proto.RevokeResponse{Success: true}, nil
}
