# Backend Architecture Documentation

This document provides an in-depth explanation of the backend architecture, implementation details, service communication patterns, database design, authentication mechanisms, and operational procedures.

## 📋 Table of Contents

1. [Overview](#overview)
2. [Architecture Patterns](#architecture-patterns)
3. [Service Details](#service-details)
4. [Communication Layer](#communication-layer)
5. [Data Layer](#data-layer)
6. [Authentication & Authorization](#authentication--authorization)
7. [Environment Configuration](#environment-configuration)
8. [Development Workflow](#development-workflow)
9. [Production Considerations](#production-considerations)
10. [Troubleshooting Guide](#troubleshooting-guide)

## Overview

### Technology Stack

| Component | Technology | Version | Purpose |
|-----------|-----------|---------|---------|
| **Runtime** | Go | 1.24 | Primary language |
| **Communication** | gRPC + Protocol Buffers | Latest | Inter-service RPC |
| **Database** | PostgreSQL | 13 | Primary data store |
| **Messaging** | Apache Kafka | TBD | Event streaming (planned) |
| **Authentication** | JWT | golang-jwt/jwt v5 | Token-based auth |
| **Configuration** | Viper | v1.21.0 | Config management |
| **Logging** | Zap | v1.27.0 | Structured logging |
| **Containerization** | Docker & Compose | Latest | Service orchestration |

### Project Structure

```
GoApp/
├── cmd/                      # Service entry points (main.go files)
│   ├── auth_service/        # Authentication microservice
│   ├── chat_service/        # Chat management microservice
│   ├── realtime_service/    # WebSocket real-time service
│   ├── status_service/      # User status/stories service
│   ├── message_worker/      # Kafka message consumer
│   └── api_gateway/         # HTTP to gRPC gateway
│
├── internal/                 # Private application logic
│   ├── auth/                # Auth domain implementation
│   ├── chat/                # Chat domain implementation
│   ├── realtime/            # Real-time messaging logic
│   └── worker/              # Background job processing
│
├── pkg/                      # Shared/public libraries
│   ├── config/              # Configuration management
│   ├── database/            # PostgreSQL connection pooling
│   ├── domain/              # Domain models and types
│   ├── eventbus/            # Kafka client wrapper
│   ├── jwt/                 # JWT token management
│   └── logger/              # Structured logging utilities
│
├── proto/                    # Protocol Buffer definitions
│   ├── auth.proto           # Auth service contract
│   ├── chat.proto           # Chat service contract
│   ├── realtime.proto       # Realtime service contract
│   └── *.pb.go              # Generated Go code
│
└── migrations/               # Database schema versions
    ├── 0000_enable_uuid.sql
    ├── 0001_initial_schema.up.sql
    ├── 0002_user_devices_revocation.up.sql
    ├── 0003_chat_schema.up.sql
    └── 0004_add_group_support.up.sql
```

### Key Design Principles

1. **Clean Architecture**: Separation of concerns with clear layer boundaries
2. **Domain-Driven Design**: Business logic organized by domain
3. **Microservices**: Independent, deployable services
4. **API-First**: Protocol Buffers define service contracts
5. **Stateless Services**: Horizontal scalability
6. **Database per Service**: Data isolation (future consideration)

## Architecture Patterns

### Clean Architecture Implementation

The project follows Uncle Bob's Clean Architecture with these layers:

```
┌─────────────────────────────────────────────────┐
│              Handler Layer (gRPC)               │  ← External interface
│  - Handles gRPC requests                        │
│  - Input validation                             │
│  - Response marshaling                          │
└────────────────┬────────────────────────────────┘
                 │ depends on ↓
┌────────────────▼────────────────────────────────┐
│              Service Layer                      │  ← Business logic
│  - Domain logic                                 │
│  - Use case orchestration                       │
│  - Transaction management                       │
└────────────────┬────────────────────────────────┘
                 │ depends on ↓
┌────────────────▼────────────────────────────────┐
│            Repository Layer                     │  ← Data abstraction
│  - Interfaces for data access                   │
│  - Domain model definitions                     │
└────────────────┬────────────────────────────────┘
                 │ implemented by ↓
┌────────────────▼────────────────────────────────┐
│              Store Layer                        │  ← Data access
│  - PostgreSQL implementations                   │
│  - Query execution                              │
│  - Data mapping                                 │
└─────────────────────────────────────────────────┘
```

**Dependency Rule**: Inner layers don't depend on outer layers. Dependencies point inward.

### Example: Auth Service Layers

**Handler** (`internal/auth/handler/grpc.go`):
```go
type AuthHandler struct {
    authService service.AuthService
    logger      *zap.Logger
}

func (h *AuthHandler) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.AuthResponse, error) {
    // 1. Validate input
    // 2. Call service layer
    user, tokens, err := h.authService.VerifyOTP(ctx, req.PhoneNumber, req.OtpCode)
    // 3. Map to response
    return &pb.AuthResponse{...}, nil
}
```

**Service** (`internal/auth/service/service.go`):
```go
type authService struct {
    userRepo   repository.UserRepository
    deviceRepo repository.DeviceRepository
    tokenMgr   *jwt.TokenManager
    twilioClient *twilio.Client
}

func (s *authService) VerifyOTP(ctx context.Context, phone, code string) (*domain.User, *Tokens, error) {
    // 1. Verify OTP with Twilio
    // 2. Get or create user via repository
    // 3. Generate JWT tokens
    // 4. Store device session
    // 5. Return user and tokens
}
```

**Repository** (`internal/auth/repository/repository.go`):
```go
type UserRepository interface {
    GetByPhoneNumber(ctx context.Context, phoneNumber string) (*domain.User, error)
    Create(ctx context.Context, user *domain.User) error
    Update(ctx context.Context, user *domain.User) error
}
```

**Store** (`internal/auth/store/postgres.go`):
```go
type postgresUserRepository struct {
    db *sql.DB
}

func (r *postgresUserRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*domain.User, error) {
    query := `SELECT id, phone_number, display_name, ... FROM users WHERE phone_number = $1`
    // Execute query and map to domain.User
}
```

### Microservices Architecture

Each service is independently deployable with its own:
- Entry point (`cmd/service_name/main.go`)
- Domain logic (`internal/service_name/`)
- gRPC server configuration
- Database connection pool
- Configuration management

**Inter-Service Communication**:
- Synchronous: gRPC (for request-response)
- Asynchronous: Kafka (for events, planned)

**Service Discovery** (Future):
- Currently: Direct connection via host:port
- Planned: Service mesh or discovery service

## Service Details

### Auth Service (Active)

**Location**: `cmd/auth_service/main.go`

**Port**: 50051 (gRPC)

**Responsibilities**:
1. Phone number verification via Twilio OTP
2. User registration and profile management
3. JWT token generation (access + refresh)
4. Token validation and refresh
5. Device session management
6. Session revocation (single/all devices)

**External Dependencies**:
- PostgreSQL: User and device data
- Twilio Verify API: OTP delivery
- None for other services (self-contained)

**Internal Components**:

```
internal/auth/
├── handler/
│   └── grpc.go              # gRPC endpoint implementations
├── service/
│   ├── service.go           # Business logic
│   └── service_test.go      # Unit tests
├── repository/
│   ├── repository.go        # Data interfaces
│   └── device_repository.go # Device-specific interfaces
├── store/
│   ├── postgres.go          # User data access
│   └── user_device_postgres.go # Device data access
└── middleware/
    └── auth_interceptor.go  # JWT validation middleware
```

**gRPC Methods** (defined in `proto/auth.proto`):

1. **SendOTP**: Initiates phone verification
2. **VerifyOTP**: Validates OTP and returns tokens
3. **ValidateToken**: Checks access token validity
4. **RefreshToken**: Rotates refresh token
5. **RevokeCurrentDevice**: Logs out specific device
6. **LogoutAllDevices**: Logs out all user devices

**Configuration**:
```env
AUTH_SERVICE_GRPC_PORT=50051
AUTH_DEV_MODE=true  # Bypass Twilio for development
TWILIO_ACCOUNT_SID=...
TWILIO_AUTH_TOKEN=...
TWILIO_VERIFY_SERVICE_SID=...
```

### Chat Service (In Development)

**Location**: `cmd/chat_service/main.go`

**Planned Port**: 50052 (gRPC)

**Responsibilities**:
1. Create and manage chat rooms (direct + group)
2. Add/remove chat room members
3. Send and receive messages
4. Message history and pagination
5. Read receipts and delivery status
6. Typing indicators

**Database Tables**:
- `chat_rooms`: Room metadata
- `chat_room_members`: Membership records
- `messages`: Message content and metadata

**Planned gRPC Methods**:
- `CreateChatRoom`
- `GetChatRoom`
- `ListChatRooms`
- `SendMessage`
- `GetMessages`
- `MarkAsRead`

### Realtime Service (In Development)

**Location**: `cmd/realtime_service/main.go`

**Planned Ports**:
- 50053 (gRPC for service communication)
- 8080 (WebSocket for clients)

**Responsibilities**:
1. Manage WebSocket connections
2. Real-time message broadcasting
3. Online/offline presence
4. Typing indicators
5. Connection pooling

**Components**:
```
internal/realtime/
├── hub.go           # WebSocket connection hub
├── global.go        # Global hub instance
└── handler/
    └── grpc.go      # gRPC endpoints for other services
```

**Architecture**:
```
Client WebSocket ←→ Hub (goroutine) ←→ gRPC (from Chat Service)
                         ↓
                    PostgreSQL (presence)
```

### Status Service (In Development)

**Location**: `cmd/status_service/main.go`

**Planned Port**: 50054 (gRPC)

**Responsibilities**:
1. Create and manage user stories
2. 24-hour auto-expiry
3. View tracking
4. Media attachment support
5. Privacy controls

### Message Worker (In Development)

**Location**: `cmd/message_worker/main.go`

**Responsibilities**:
1. Consume messages from Kafka
2. Persist messages to database
3. Trigger push notifications
4. Process media uploads
5. Handle failed deliveries

**Kafka Topics** (Planned):
- `messages.sent`: New messages
- `messages.delivered`: Delivery confirmations
- `messages.read`: Read receipts

### API Gateway (In Development)

**Location**: `cmd/api_gateway/main.go`

**Planned Port**: 8000 (HTTP/REST)

**Responsibilities**:
1. HTTP to gRPC translation
2. Request routing to services
3. Authentication verification
4. Rate limiting
5. API versioning
6. CORS handling

**Architecture**:
```
Client (HTTP/REST) → API Gateway → gRPC Services
                         ↓
                    JWT Validation
                    Rate Limiting
                    Request Logging
```

## Communication Layer

### gRPC & Protocol Buffers

**Why gRPC?**
- Type-safe service contracts
- High performance (HTTP/2, binary)
- Bidirectional streaming support
- Code generation for multiple languages
- Built-in authentication and load balancing

**Proto File Structure**:

```protobuf
// proto/auth.proto
syntax = "proto3";

package auth;
option go_package = "github.com/dykethecreator/GoApp/proto";

service AuthService {
  rpc SendOTP(SendOTPRequest) returns (SendOTPResponse);
  rpc VerifyOTP(VerifyOTPRequest) returns (AuthResponse);
  // ... more methods
}

message SendOTPRequest {
  string phone_number = 1;
}

message SendOTPResponse {
  string message = 1;
}
```

**Code Generation**:

```powershell
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/auth.proto
```

Generates:
- `auth.pb.go`: Message types
- `auth_grpc.pb.go`: Service interfaces and client/server code

**gRPC Server Setup**:

```go
func main() {
    lis, _ := net.Listen("tcp", ":50051")
    
    grpcServer := grpc.NewServer(
        grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor(tokenMgr)),
    )
    
    pb.RegisterAuthServiceServer(grpcServer, authHandler)
    
    grpcServer.Serve(lis)
}
```

**gRPC Client Usage** (from other services):

```go
conn, _ := grpc.Dial("auth-service:50051", grpc.WithInsecure())
defer conn.Close()

client := pb.NewAuthServiceClient(conn)
resp, _ := client.ValidateToken(ctx, &pb.ValidateTokenRequest{
    AccessToken: token,
})
```

### Kafka Integration (Planned)

**Event-Driven Communication**:

Producer (Chat Service):
```go
producer := eventbus.NewKafkaProducer()
producer.Send("messages.sent", &domain.Message{...})
```

Consumer (Message Worker):
```go
consumer := eventbus.NewKafkaConsumer("messages.sent")
for msg := range consumer.Messages() {
    // Process message
}
```

## Data Layer

### Database Design

**PostgreSQL Version**: 13

**Connection Management**:

Location: `pkg/database/postgres.go`

```go
func NewPostgresDB(databaseURL string) (*sql.DB, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }
    
    // Connection pooling
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)
    
    return db, db.Ping()
}
```

**Migration Management**:

Migrations are SQL files in `migrations/` directory:

```
0000_enable_uuid.sql              # Enable UUID extension
0001_initial_schema.up.sql        # Base schema
0002_user_devices_revocation.up.sql  # Device sessions
0003_chat_schema.up.sql           # Chat tables
0004_add_group_support.up.sql    # Group features
```

**Applying Migrations**:

Automated:
```powershell
.\scripts\setup-db.ps1
```

Manual:
```powershell
Get-Content migrations\0001_initial_schema.up.sql | `
  docker compose exec -T postgres psql -U user -d whatsapp_clone_dev
```

### Schema Details

#### Users Table

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    display_name VARCHAR(255),
    profile_picture_url TEXT,
    bio TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_users_phone ON users(phone_number);
```

**Purpose**: Store user profiles

**Relationships**:
- One-to-many with `user_devices`
- One-to-many with `messages`
- Many-to-many with `chat_rooms` via `chat_room_members`

#### User Devices Table

```sql
CREATE TABLE user_devices (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    refresh_token_hash VARCHAR(64) NOT NULL,
    device_name VARCHAR(255),
    device_type VARCHAR(50),
    push_notification_token TEXT,
    last_login_at TIMESTAMP DEFAULT NOW(),
    revoked_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    
    CONSTRAINT unique_user_refresh_hash UNIQUE (user_id, refresh_token_hash)
);

CREATE INDEX idx_user_devices_user_id ON user_devices(user_id);
CREATE INDEX idx_user_devices_revocation ON user_devices(user_id, revoked_at);
```

**Purpose**: Track device sessions and refresh tokens

**Key Fields**:
- `refresh_token_hash`: SHA-256 hash of refresh token
- `revoked_at`: NULL for active sessions, timestamp for revoked

**Queries**:

Find active device:
```sql
SELECT * FROM user_devices 
WHERE user_id = $1 
  AND refresh_token_hash = $2 
  AND revoked_at IS NULL;
```

Revoke all user devices:
```sql
UPDATE user_devices 
SET revoked_at = NOW() 
WHERE user_id = $1 AND revoked_at IS NULL;
```

#### Chat Rooms Table

```sql
CREATE TABLE chat_rooms (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    is_group BOOLEAN DEFAULT false,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_chat_rooms_created_by ON chat_rooms(created_by);
```

**Purpose**: Store chat room metadata

**Types**:
- Direct (1-on-1): `is_group = false`, `name` optional
- Group: `is_group = true`, `name` required

#### Messages Table

```sql
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    chat_room_id UUID NOT NULL REFERENCES chat_rooms(id) ON DELETE CASCADE,
    sender_id UUID REFERENCES users(id) ON DELETE SET NULL,
    content TEXT NOT NULL,
    message_type VARCHAR(20) DEFAULT 'text',
    status VARCHAR(20) DEFAULT 'sent',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_messages_chat_room ON messages(chat_room_id, created_at DESC);
CREATE INDEX idx_messages_sender ON messages(sender_id);
```

**Purpose**: Store message content and metadata

**Message Types**: `text`, `image`, `video`, `audio`, `document`

**Status Flow**: `sent` → `delivered` → `read`

### Data Access Patterns

**Repository Pattern**:

Interface defines contract:
```go
type UserRepository interface {
    GetByID(ctx context.Context, id string) (*domain.User, error)
    GetByPhoneNumber(ctx context.Context, phone string) (*domain.User, error)
    Create(ctx context.Context, user *domain.User) error
    Update(ctx context.Context, user *domain.User) error
}
```

PostgreSQL implementation:
```go
type postgresUserRepository struct {
    db *sql.DB
}

func (r *postgresUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
    query := `SELECT id, phone_number, display_name, profile_picture_url, bio, created_at, updated_at 
              FROM users WHERE id = $1`
    
    var user domain.User
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID, &user.PhoneNumber, &user.DisplayName,
        &user.ProfilePictureURL, &user.Bio, &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err == sql.ErrNoRows {
        return nil, ErrUserNotFound
    }
    
    return &user, err
}
```

## Authentication & Authorization

### JWT Token Architecture

**Token Manager** (`pkg/jwt/token.go`):

```go
type TokenManager struct {
    secretKey          []byte
    accessTokenExpiry  time.Duration  // 15 minutes
    refreshTokenExpiry time.Duration  // 7 days
}

type CustomClaims struct {
    Type string `json:"type"` // "access" or "refresh"
    jwt.RegisteredClaims
}
```

**Token Generation**:

Access Token:
```go
func (tm *TokenManager) GenerateAccessToken(userID string) (string, error) {
    claims := CustomClaims{
        Type: "access",
        RegisteredClaims: jwt.RegisteredClaims{
            Subject:   userID,
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.accessTokenExpiry)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "my-auth-service",
            Audience:  []string{"my-app-client"},
            ID:        uuid.New().String(),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(tm.secretKey)
}
```

Refresh Token:
```go
func (tm *TokenManager) GenerateRefreshToken(userID string) (string, error) {
    claims := CustomClaims{
        Type: "refresh",
        RegisteredClaims: jwt.RegisteredClaims{
            Subject:   userID,
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.refreshTokenExpiry)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "my-auth-service",
            Audience:  []string{"my-auth-service"},
            ID:        uuid.New().String(),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(tm.secretKey)
}
```

**Token Validation**:

```go
func (tm *TokenManager) ValidateToken(tokenString string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        // Verify signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return tm.secretKey, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, errors.New("invalid token")
}
```

### Refresh Token Rotation

**Security Feature**: Prevents token replay attacks by invalidating refresh tokens after use.

**Flow**:

1. Client sends refresh token to `RefreshToken` endpoint
2. Server validates token:
   ```go
   claims, err := tokenMgr.ValidateToken(refreshToken)
   if err != nil || claims.Type != "refresh" {
       return ErrInvalidToken
   }
   ```

3. Server computes token hash and checks database:
   ```go
   hash := sha256.Sum256([]byte(refreshToken))
   hashStr := hex.EncodeToString(hash[:])
   
   device, err := deviceRepo.GetActiveDeviceByHash(ctx, claims.Subject, hashStr)
   if err != nil {
       return ErrDeviceNotFound
   }
   ```

4. Server revokes current session:
   ```go
   deviceRepo.RevokeDevice(ctx, device.ID)
   ```

5. Server generates new token pair:
   ```go
   newAccessToken, _ := tokenMgr.GenerateAccessToken(claims.Subject)
   newRefreshToken, _ := tokenMgr.GenerateRefreshToken(claims.Subject)
   ```

6. Server stores new refresh token hash:
   ```go
   newHash := sha256.Sum256([]byte(newRefreshToken))
   newDevice := &domain.UserDevice{
       UserID:           claims.Subject,
       RefreshTokenHash: hex.EncodeToString(newHash[:]),
       DeviceName:       device.DeviceName,
       DeviceType:       device.DeviceType,
       LastLoginAt:      time.Now(),
   }
   deviceRepo.Create(ctx, newDevice)
   ```

7. Server returns new tokens to client

**Security Benefits**:
- Old refresh token becomes invalid immediately
- Token theft detection (if old token is reused)
- Limits token lifetime even if stolen
- Maintains user sessions across rotations

### Device Session Management

**Device Registration** (on VerifyOTP):

```go
func (s *authService) VerifyOTP(ctx context.Context, phone, code string) (*domain.User, *Tokens, error) {
    // ... verify OTP ...
    
    // Generate tokens
    accessToken, _ := s.tokenMgr.GenerateAccessToken(user.ID)
    refreshToken, _ := s.tokenMgr.GenerateRefreshToken(user.ID)
    
    // Hash and store refresh token
    hash := sha256.Sum256([]byte(refreshToken))
    device := &domain.UserDevice{
        UserID:           user.ID,
        RefreshTokenHash: hex.EncodeToString(hash[:]),
        DeviceName:       "Unknown Device",
        DeviceType:       "mobile",
        LastLoginAt:      time.Now(),
    }
    
    s.deviceRepo.Create(ctx, device)
    
    return user, &Tokens{accessToken, refreshToken}, nil
}
```

**Active Session Query**:

```sql
SELECT id, user_id, device_name, device_type, last_login_at
FROM user_devices
WHERE user_id = $1 
  AND refresh_token_hash = $2 
  AND revoked_at IS NULL;
```

**Session Revocation**:

Single device (by refresh token):
```go
func (s *authService) RevokeCurrentDevice(ctx context.Context, refreshToken string) error {
    claims, _ := s.tokenMgr.ValidateToken(refreshToken)
    hash := sha256.Sum256([]byte(refreshToken))
    hashStr := hex.EncodeToString(hash[:])
    
    return s.deviceRepo.RevokeByHash(ctx, claims.Subject, hashStr)
}
```

All devices (by access token):
```go
func (s *authService) LogoutAllDevices(ctx context.Context, accessToken string) error {
    claims, _ := s.tokenMgr.ValidateToken(accessToken)
    return s.deviceRepo.RevokeAllUserDevices(ctx, claims.Subject)
}
```

### gRPC Authentication Interceptor

**Location**: `internal/auth/middleware/auth_interceptor.go`

**Purpose**: Validates JWT access tokens in gRPC requests

**Implementation**:

```go
func UnaryAuthInterceptor(tokenManager *jwt.TokenManager) grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
        // Exempt public methods
        publicMethods := map[string]bool{
            "/auth.AuthService/SendOTP":    true,
            "/auth.AuthService/VerifyOTP":  true,
        }
        
        if publicMethods[info.FullMethod] {
            return handler(ctx, req)
        }
        
        // Extract token from metadata
        md, ok := metadata.FromIncomingContext(ctx)
        if !ok {
            return nil, status.Error(codes.Unauthenticated, "missing metadata")
        }
        
        values := md.Get("authorization")
        if len(values) == 0 {
            return nil, status.Error(codes.Unauthenticated, "missing authorization header")
        }
        
        token := strings.TrimPrefix(values[0], "Bearer ")
        
        // Validate token
        claims, err := tokenManager.ValidateToken(token)
        if err != nil {
            return nil, status.Error(codes.Unauthenticated, "invalid token")
        }
        
        if claims.Type != "access" {
            return nil, status.Error(codes.Unauthenticated, "invalid token type")
        }
        
        // Inject user ID into context
        ctx = context.WithValue(ctx, "user_id", claims.Subject)
        
        return handler(ctx, req)
    }
}
```

**Usage in Services**:

```go
func main() {
    tokenMgr, _ := jwt.NewTokenManager(
        []byte(os.Getenv("JWT_SECRET")),
        15*time.Minute,
        7*24*time.Hour,
    )
    
    grpcServer := grpc.NewServer(
        grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor(tokenMgr)),
    )
    
    // Register services...
}
```

**Accessing User ID in Handlers**:

```go
func (h *Handler) ProtectedMethod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    userID, ok := ctx.Value("user_id").(string)
    if !ok {
        return nil, status.Error(codes.Internal, "user ID not found in context")
    }
    
    // Use userID...
}
```

### Twilio OTP Integration

**Configuration**:

```env
TWILIO_ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
TWILIO_AUTH_TOKEN=your_auth_token
TWILIO_VERIFY_SERVICE_SID=VAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
AUTH_DEV_MODE=false  # Set to true to bypass Twilio
```

**Send OTP**:

```go
func (s *authService) SendOTP(ctx context.Context, phoneNumber string) error {
    if s.devMode {
        s.logger.Info("DEV MODE: OTP is 123456")
        return nil
    }
    
    params := &verify.CreateVerificationParams{}
    params.SetTo(phoneNumber)
    params.SetChannel("sms")
    
    _, err := s.twilioClient.VerifyV2.CreateVerification(s.verifyServiceSID, params)
    return err
}
```

**Verify OTP**:

```go
func (s *authService) VerifyOTPCode(ctx context.Context, phoneNumber, code string) (bool, error) {
    if s.devMode {
        return code == "123456", nil
    }
    
    params := &verify.CreateVerificationCheckParams{}
    params.SetTo(phoneNumber)
    params.SetCode(code)
    
    resp, err := s.twilioClient.VerifyV2.CreateVerificationCheck(s.verifyServiceSID, params)
    if err != nil {
        return false, err
    }
    
    return *resp.Status == "approved", nil
}
```

## Environment Configuration

### Configuration Files

**Environment-Based Loading**:

The application automatically loads configuration based on runtime context:

| Environment | File | Trigger |
|-------------|------|---------|
| Local Development | `.env.local` | Default |
| Docker Container | `.env.docker` | `RUNNING_IN_DOCKER=true` |
| Base Configuration | `.env` | Always loaded last |

**Configuration Manager** (`pkg/config/config.go`):

```go
func LoadConfig() (*Config, error) {
    // Determine which .env file to load
    envFile := ".env.local"
    if os.Getenv("RUNNING_IN_DOCKER") == "true" {
        envFile = ".env.docker"
    }
    
    // Load environment-specific file
    if err := godotenv.Load(envFile); err != nil {
        log.Printf("Warning: %s not found, using environment variables", envFile)
    }
    
    // Load base .env for overrides
    _ = godotenv.Load()
    
    // Initialize Viper
    viper.AutomaticEnv()
    
    return &Config{
        DatabaseURL:          viper.GetString("DATABASE_URL"),
        JWTSecret:            viper.GetString("JWT_SECRET"),
        AuthServiceGRPCPort:  viper.GetString("AUTH_SERVICE_GRPC_PORT"),
        TwilioAccountSID:     viper.GetString("TWILIO_ACCOUNT_SID"),
        TwilioAuthToken:      viper.GetString("TWILIO_AUTH_TOKEN"),
        TwilioVerifyServiceSID: viper.GetString("TWILIO_VERIFY_SERVICE_SID"),
        AuthDevMode:          viper.GetBool("AUTH_DEV_MODE"),
    }, nil
}
```

### Environment Variables Reference

#### Database Configuration

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `DATABASE_URL` | string | - | PostgreSQL connection string |

Example:
```env
DATABASE_URL=postgres://user:password@localhost:5433/whatsapp_clone_dev?sslmode=disable
```

#### Authentication Configuration

| Variable | Type | Default | Description |
|----------|------|---------|-------------|
| `JWT_SECRET` | string | - | HMAC signing key (min 32 chars) |
| `AUTH_DEV_MODE` | bool | false | Bypass Twilio (OTP = "123456") |
| `AUTH_SERVICE_GRPC_PORT` | string | 50051 | gRPC server port |

#### Twilio Configuration (Production)

## Database & migrations

Migrations live under `migrations/`:

- `0001_initial_schema.*.sql`: base schema including `users` and related tables
- `0002_user_devices_revocation.*.sql`: adds `user_devices` fields for refresh token hash storage and revocation (`revoked_at`), unique constraints and indexes

Tables (selected):
- `users`: stores user profiles
- `user_devices`: stores device sessions keyed by the SHA-256 hash of refresh tokens; contains `revoked_at` for session revocation

Notes:
- Ensure the UUID extension exists (e.g., `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`) if referenced by schema.
- Run migrations against the exact DB the service uses (match `DATABASE_URL`).

## Authentication model

The system issues a pair of JWTs per login: an access token (short-lived) and a refresh token (longer-lived). The refresh token is device-scoped and stored as a SHA-256 hash in `user_devices` for revocation.

JWT claims (see `pkg/jwt/token.go`):
- Custom claim: `type` ∈ { `access`, `refresh` }
- Registered claims used:
  - `sub` (Subject) = user ID
  - `jti` (JWT ID) = unique token id
  - `iss` (Issuer) = `my-auth-service`
  - `aud` (Audience) = `my-app-client` for access, `my-auth-service` for refresh
  - `iat` (Issued At), `exp` (Expires At)

Token durations (defaults in code):
- Access tokens: 15 minutes
- Refresh tokens: 7 days

Refresh rotation:
- On `RefreshToken`, the provided refresh token is validated and matched in `user_devices`.
- The matching device session is revoked (`revoked_at` set).
- New access and refresh tokens are issued; the new refresh token hash is stored in `user_devices`.
- Result: previous refresh token cannot be reused (one-time refresh semantics).

Session revocation:
- Revoke single device by refresh token.
- Revoke all devices by access token (extracts `sub` and revokes all active sessions for the user).

## Auth Service gRPC API (summary)

Service: `auth.AuthService` (see `proto/auth.proto`). Requests use JSON mapping in tools like Postman.

- SendOTP
  - Request: `{ "phone_number": "+90..." }`
  - Response: `{ "message": "..." }`

- VerifyOTP
  - Request: `{ "phone_number": "+90...", "otp_code": "123456" }`
  - Response: `{ "user": { ... }, "access_token": "...", "refresh_token": "..." }`

- ValidateToken
  - Request: `{ "access_token": "..." }`
  - Response: `{ "is_valid": true|false, "user_id": "..." }`
  - Notes: returns `is_valid=false` on invalid token instead of an error.

- RefreshToken (with rotation)
  - Request: `{ "refresh_token": "..." }`
  - Response: `{ "access_token": "...", "refresh_token": "..." }`

- RevokeCurrentDevice
  - Request: `{ "refresh_token": "..." }`
  - Response: `{ "success": true }`

- LogoutAllDevices
  - Request: `{ "access_token": "..." }`
  - Response: `{ "success": true }`

## Access token interceptor

File: `internal/auth/middleware/auth_interceptor.go`.

- A gRPC unary interceptor validates the `authorization: Bearer <access-token>` header and injects `user_id` into the request context.
- In the auth service it is wired with an exemption for `auth.AuthService` methods (they remain public).
- Other services can enable it by constructing a `TokenManager` with the same `JWT_SECRET` and using `grpc.UnaryInterceptor(UnaryAuthInterceptor(tm))` when creating the server.

## Local development quickstart (Auth Service)

Prerequisites: Docker (for Postgres), Go installed.

1) Start PostgreSQL via Docker Compose (optional example):
   - `docker-compose up -d postgres`

2) Apply migrations to your target DB. Ensure the UUID extension exists when required. Example using the Postgres container + psql:
   - Create extension: `docker exec -i <pg_id> psql -h 127.0.0.1 -U user -d whatsapp_clone_dev -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"`
   - Apply schema: pipe `migrations/0001_initial_schema.up.sql` and `0002_user_devices_revocation.up.sql` to psql.

3) Run the auth service locally:
   - `go run ./cmd/auth_service`

4) Test with Postman gRPC:
   - Server: `localhost:50051` (plaintext)
   - Import `proto/auth.proto`, pick `auth.AuthService`
   - Call `SendOTP`, `VerifyOTP`, then `RefreshToken`, `ValidateToken`, and revocation endpoints

## Troubleshooting

- Postman says: "Message violates its Protobuf type definition"
  - Ensure request field names match the proto (e.g., `access_token`, not `accessToken`). Re-import `proto/auth.proto` if in doubt.
- DB errors (relation does not exist): verify migrations applied to the same DB as in `DATABASE_URL`.
- `function uuid_generate_v4() does not exist`: run `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` on the target DB, or adjust schema to use `pgcrypto` `gen_random_uuid()`.
- JWT invalid/expired: server intentionally returns generic `Unauthenticated` for security; verify `JWT_SECRET` and token types (use refresh only at refresh endpoint).

## Roadmap (short)

- Implement API Gateway routing and per-service gRPC clients
- Expand Chat/Status/Realtime service APIs and wire interceptor
- Add unit tests for Auth flows and DB stores
- Add metrics/tracing and structured logging
