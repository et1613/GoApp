package service

import (
	"context"
	"testing"
	"time"

	"github.com/dykethecreator/GoApp/pkg/domain"
	appjwt "github.com/dykethecreator/GoApp/pkg/jwt"
	"github.com/google/uuid"
)

// --- Fakes ---

type fakeUserRepo struct{}

func (f *fakeUserRepo) FindByPhoneNumber(ctx context.Context, phone string) (*domain.User, error) {
	return nil, nil
}
func (f *fakeUserRepo) CreateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	return u, nil
}
func (f *fakeUserRepo) FindByID(ctx context.Context, userID string) (*domain.User, error) {
	id, _ := uuid.Parse(userID)
	return &domain.User{ID: id, PhoneNumber: "+900000000000"}, nil
}

type deviceRecord struct {
	dev *domain.UserDevice
}

type fakeDeviceRepo struct {
	// key: userID|hash
	store map[string]deviceRecord
}

func newFakeDeviceRepo() *fakeDeviceRepo { return &fakeDeviceRepo{store: map[string]deviceRecord{}} }

func (f *fakeDeviceRepo) key(uid, hash string) string { return uid + "|" + hash }

func (f *fakeDeviceRepo) UpsertDevice(ctx context.Context, dev *domain.UserDevice) error {
	if dev.ID == uuid.Nil {
		dev.ID = uuid.New()
	}
	k := f.key(dev.UserID.String(), dev.RefreshTokenHash)
	f.store[k] = deviceRecord{dev: dev}
	return nil
}

func (f *fakeDeviceRepo) FindActiveByUserAndHash(ctx context.Context, userID string, hash string) (*domain.UserDevice, error) {
	k := f.key(userID, hash)
	rec, ok := f.store[k]
	if !ok {
		return nil, nil
	}
	if rec.dev.RevokedAt != nil {
		return nil, nil
	}
	return rec.dev, nil
}

func (f *fakeDeviceRepo) RevokeByID(ctx context.Context, id string) error {
	// linear scan fine for test
	for k, rec := range f.store {
		if rec.dev.ID.String() == id && rec.dev.RevokedAt == nil {
			now := time.Now()
			rec.dev.RevokedAt = &now
			f.store[k] = rec
			break
		}
	}
	return nil
}

func (f *fakeDeviceRepo) RevokeAllForUser(ctx context.Context, userID string) error {
	for k, rec := range f.store {
		if rec.dev.UserID.String() == userID && rec.dev.RevokedAt == nil {
			now := time.Now()
			rec.dev.RevokedAt = &now
			f.store[k] = rec
		}
	}
	return nil
}

// --- Tests ---

func TestRefreshToken_RotationAndRevocation(t *testing.T) {
	tm, err := appjwt.NewTokenManager("0123456789abcdefghijklmnopqrstuvwxyz!@#%$^&*()", 15*time.Minute, 24*time.Hour)
	if err != nil {
		t.Fatalf("token manager: %v", err)
	}

	// fabricate service instance without calling NewAuthService (to avoid Twilio deps)
	s := &AuthService{
		userRepo:     &fakeUserRepo{},
		deviceRepo:   newFakeDeviceRepo(),
		tokenManager: tm,
	}

	userID := uuid.NewString()
	_, refresh, err := tm.GenerateTokens(userID)
	if err != nil {
		t.Fatalf("GenerateTokens: %v", err)
	}

	// simulate VerifyOTP persistence of device
	hash := hashRefreshToken(refresh)
	dev := &domain.UserDevice{UserID: uuid.MustParse(userID), RefreshTokenHash: hash, DeviceName: "test", DeviceType: "test", LastLoginAt: time.Now()}
	if err := s.deviceRepo.UpsertDevice(context.Background(), dev); err != nil {
		t.Fatalf("UpsertDevice: %v", err)
	}

	// perform refresh -> expect rotation
	newAccess, newRefresh, err := s.RefreshToken(context.Background(), refresh)
	if err != nil {
		t.Fatalf("RefreshToken error: %v", err)
	}
	if newAccess == "" || newRefresh == "" {
		t.Fatalf("expected new tokens returned")
	}
	if newRefresh == refresh {
		t.Fatalf("expected refresh token rotation")
	}

	// old refresh should no longer be valid/active in repo
	if d, _ := s.deviceRepo.FindActiveByUserAndHash(context.Background(), userID, hash); d != nil {
		t.Fatalf("expected old device revoked, but still active")
	}

	// new refresh should be persisted and active
	newHash := hashRefreshToken(newRefresh)
	if d, _ := s.deviceRepo.FindActiveByUserAndHash(context.Background(), userID, newHash); d == nil {
		t.Fatalf("expected new device session to be active")
	}
}
