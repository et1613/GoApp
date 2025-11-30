# WhatsApp Clone Backend# WhatsApp Clone Backend# WhatsApp Clone Backend



[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)

[![gRPC](https://img.shields.io/badge/gRPC-Protocol%20Buffers-green.svg)](https://grpc.io)

[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue.svg)](https://www.postgresql.org)[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)This project is a Go-based backend for a WhatsApp-like application, built with a microservices architecture.



A WhatsApp-like application backend built with Go and microservices architecture. Features JWT-based authentication, real-time messaging, and gRPC communication protocol.[![gRPC](https://img.shields.io/badge/gRPC-Protocol%20Buffers-green.svg)](https://grpc.io)



## 🏗️ Project Structure[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue.svg)](https://www.postgresql.org)## Project Structure



The project follows Clean Architecture and microservices pattern:



```Go tabanlı, mikroservis mimarisi ile geliştirilmiş bir WhatsApp benzeri uygulama backend'idir. JWT tabanlı kimlik doğrulama, gerçek zamanlı mesajlaşma ve gRPC iletişim protokolü içerir.The project follows a clean architecture and microservices pattern.

GoApp/

├── cmd/                    # Service entry points

│   ├── auth_service/      # Authentication service

│   ├── chat_service/      # Chat management service## 🏗️ Proje Yapısı- `/cmd`: Entry points for each service (`main.go`).

│   ├── realtime_service/  # WebSocket real-time communication

│   ├── status_service/    # User status (stories) service- `/internal`: Private application and business logic for each service.

│   ├── message_worker/    # Message queue processor

│   └── api_gateway/       # API Gateway (in development)Proje Clean Architecture ve mikroservis desenini takip eder:- `/pkg`: Shared libraries and domain types used across services.

├── internal/              # Private application and business logic

│   ├── auth/             # Auth domain logic- `/proto`: gRPC protocol definitions for inter-service communication.

│   │   ├── handler/      # gRPC handlers

│   │   ├── service/      # Business logic layer```- `/migrations`: Database schema migrations.

│   │   ├── repository/   # Data access interfaces

│   │   ├── store/        # PostgreSQL implementationsGoApp/- `/configs`: Configuration files for different environments.

│   │   └── middleware/   # JWT interceptor

│   ├── chat/             # Chat domain logic├── cmd/                    # Servis giriş noktaları

│   ├── realtime/         # WebSocket hub and handlers

│   └── worker/           # Kafka consumer│   ├── auth_service/      # Kimlik doğrulama servisi## Getting Started

├── pkg/                   # Shared libraries

│   ├── config/           # Configuration management (Viper)│   ├── chat_service/      # Chat yönetim servisi

│   ├── database/         # PostgreSQL connection management

│   ├── domain/           # Domain models (User, Message, etc.)│   ├── realtime_service/  # WebSocket gerçek zamanlı iletişim### Prerequisites

│   ├── jwt/              # JWT token management

│   ├── logger/           # Zap logger│   ├── status_service/    # Kullanıcı durumu (hikaye) servisi

│   └── eventbus/         # Kafka client

├── proto/                 # gRPC Protocol Buffer definitions│   ├── message_worker/    # Mesaj kuyruğu işleyici- Docker and Docker Compose

├── migrations/            # PostgreSQL schema migrations

├── scripts/              # Helper PowerShell scripts│   └── api_gateway/       # API Gateway (geliştirme aşamasında)- Go 1.18 or higher

└── docker-compose.yml    # Docker Compose configuration

```├── internal/              # Özel uygulama ve iş mantığı



## 🚀 Quick Start│   ├── auth/             # Auth domain logic### Running the application



### Prerequisites│   │   ├── handler/      # gRPC handlers



- **Docker & Docker Compose** (for PostgreSQL)│   │   ├── service/      # İş mantığı katmanı1.  **Start the infrastructure:**

- **Go 1.24 or higher**

- **PowerShell** (for running scripts on Windows)│   │   ├── repository/   # Veri erişim arayüzleri    ```bash

- **Postman** (for gRPC testing, optional)

│   │   ├── store/        # PostgreSQL implementasyonları    docker-compose up -d

### 1. Environment Variables Setup

│   │   └── middleware/   # JWT interceptor    ```

The project automatically loads .env files based on context:

- Local run: `.env.local`│   ├── chat/             # Chat domain logic

- Docker run: `.env.docker` (with RUNNING_IN_DOCKER=true)

│   ├── realtime/         # WebSocket hub ve handlers2.  **Run database migrations:**

Copy the example file to get started:

│   └── worker/           # Kafka consumer    You'll need a migration tool like `golang-migrate/migrate`.

```powershell

Copy-Item .env.example .env.local├── pkg/                   # Paylaşılan kütüphaneler    ```bash

```

│   ├── config/           # Yapılandırma yönetimi (Viper)    migrate -database "postgres://postgres:Fbtex1967.@localhost:5432/whatsapp_clone_dev?sslmode=disable" -path migrations up

**Important variables:**

- `DATABASE_URL`: PostgreSQL connection string│   ├── database/         # PostgreSQL bağlantı yönetimi    ```

- `JWT_SECRET`: Strong secret for JWT signing (minimum 32 bytes)

- `AUTH_DEV_MODE=true`: Twilio bypass (OTP code is always `123456`)│   ├── domain/           # Domain modelleri (User, Message, vb.)

- Twilio (for Production): `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, `TWILIO_VERIFY_SERVICE_SID`

│   ├── jwt/              # JWT token yönetimi3.  **Run the services:**

### 2. Start PostgreSQL

│   ├── logger/           # Zap logger    Navigate to each service's directory and run it.

Start only PostgreSQL with Docker Compose:

│   └── eventbus/         # Kafka client    ```bash

```powershell

docker-compose up -d postgres├── proto/                 # gRPC Protocol Buffer tanımları    go run ./cmd/api_gateway/

```

├── migrations/            # PostgreSQL şema migrasyonları    go run ./cmd/auth_service/

PostgreSQL will run on port `5433` (to avoid conflicts with local PostgreSQL).

├── scripts/              # Yardımcı PowerShell scriptleri    # ... and so on for other services

### 3. Apply Database Migrations

└── docker-compose.yml    # Docker Compose yapılandırması    ```

**Automated method (recommended):**

```

```powershell

Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass## Quickstart: AuthService + Postman gRPC (Windows)

.\scripts\setup-db.ps1

```## 🚀 Hızlı Başlangıç



**Manual method:**End-to-end minimum setup to test OTP and JWT issuance via Postman using gRPC.



```powershell### Ön Gereksinimler

# Enable UUID extension

Get-Content migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev### 1) Environment variables



# Main schema- **Docker & Docker Compose** (PostgreSQL için)

Get-Content migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

- **Go 1.24 veya üstü**This service auto-loads environment files based on context:

# Device management and revocation

Get-Content migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1- **PowerShell** (Windows için script çalıştırma)



# Chat schema- **Postman** (gRPC test için, opsiyonel)- Local runs: loads `.env.local` (we added one with sensible defaults)

Get-Content migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

- Docker runs: `docker-compose` passes `.env.docker` and also sets `RUNNING_IN_DOCKER=true` (the app attempts to load `.env.docker` but works even if the file isn't baked into the image)

# Group support

Get-Content migrations\0004_add_group_support.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1### 1. Ortam Değişkenlerini Ayarlama- Base `.env` is also loaded last if present (for overrides)

```



### 4. Run Auth Service

Proje farklı ortamlar için otomatik .env dosyası yükler:Example files provided:

From the project root directory:

- Local çalıştırma: `.env.local`

```powershell

go run ./cmd/auth_service- Docker içinde: `.env.docker` (RUNNING_IN_DOCKER=true ile)- `.env.example` — connects to Postgres at `localhost:5432`, sets a sample `JWT_SECRET`, and enables `AUTH_DEV_MODE=true` (Twilio bypass: OTP code is `123456`).

```

- `.env.docker.example` — same but using `postgres:5432` for Compose and `RUNNING_IN_DOCKER=true`.

Output: `auth_service listening on :50051 (env=local)`

Başlamak için örnek dosyayı kopyalayın:

### 5. Test with Postman gRPC

Create a copy for your environment:

1. **Create a new gRPC Request in Postman**

2. **Server URL**: `localhost:50051` (plaintext/TLS off)```powershell

3. **Import proto file**: `proto/auth.proto`

4. **Test AuthService methods**:Copy-Item .env.example .env.local```powershell



**SendOTP example:**```Copy-Item .env.example .env

```json

{```

  "phone_number": "+905551234567"

}**Önemli değişkenler:**

```

- `DATABASE_URL`: PostgreSQL bağlantı dizesiEdit `.env` as needed:

**VerifyOTP example:**

```json- `JWT_SECRET`: JWT imzalama için güçlü bir secret (minimum 32 byte)

{

  "phone_number": "+905551234567",- `AUTH_DEV_MODE=true`: Twilio bypass (OTP kodu her zaman `123456`)- Set `JWT_SECRET` to a strong value

  "otp_code": "123456"

}- Twilio (Production için): `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, `TWILIO_VERIFY_SERVICE_SID`- Add real Twilio credentials if you want to use OTP verification against Twilio (otherwise keep `AUTH_DEV_MODE=true`).

```



Successful response returns `access_token` and `refresh_token`.

### 2. PostgreSQL Başlatma### 2) Start PostgreSQL (Docker)

**ValidateToken example:**

```json

{

  "access_token": "<ACCESS_TOKEN>"Docker Compose ile sadece PostgreSQL'i başlatın:Run only Postgres in the background:

}

```



**RefreshToken example:**```powershell```powershell

```json

{docker-compose up -d postgresdocker-compose up -d postgres

  "refresh_token": "<REFRESH_TOKEN>"

}``````

```



**RevokeCurrentDevice example:**

```jsonPostgreSQL `5433` portunda çalışacaktır (yerel PostgreSQL ile çakışmayı önlemek için).### 3) Apply database schema (migrations)

{

  "refresh_token": "<REFRESH_TOKEN>"

}

```### 3. Veritabanı Migrasyonlarını UygulamaOption A: Use the automated script (recommended):



**LogoutAllDevices example:**

```json

{**Otomatik yöntem (önerilen):**```powershell

  "access_token": "<ACCESS_TOKEN>"

}Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass

```

```powershell.\scripts\setup-db.ps1

## 📦 Microservices

Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass```

### Auth Service (Active ✅)

- **Port**: 50051 (gRPC).\scripts\setup-db.ps1

- **Responsibilities**:

  - Phone number verification with Twilio OTP```Option B: Manual migration:

  - JWT (access + refresh token) generation

  - Token validation and refresh (rotation)

  - Device-based session management

  - Single device or all devices logout**Manuel yöntem:**```powershell

- **Technologies**: gRPC, JWT, Twilio, PostgreSQL

- **Middleware**: JWT interceptor (portable to other services)Get-Content -Raw migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -f -



### Chat Service (In Development 🚧)```powershellGet-Content -Raw migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

- Chat room management

- User membership management# UUID extension'ı etkinleştirGet-Content -Raw migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

- Group support

Get-Content migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_devGet-Content -Raw migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

### Realtime Service (In Development 🚧)

- WebSocket connection management```

- Real-time message delivery

- Presence (online status) management# Ana şema



### Status Service (In Development 🚧)Get-Content migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1If you previously ran into an inline index syntax error for `call_logs`, this repository already fixes it by creating the index separately.

- User stories

- Status updates



### Message Worker (In Development 🚧)# Cihaz yönetimi ve revocation### 4) Run Auth Service locally

- Kafka consumer

- Asynchronous message processingGet-Content migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1



### API Gateway (In Development 🚧)From the project root:

- Single entry point for client requests

- Routing to services via gRPC clients# Chat şeması



## 🔐 Authentication ModelGet-Content migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1```powershell



The system generates two JWT tokens per login:go run ./cmd/auth_service



### Token Types# Grup desteği```



| Token Type | Validity | Purpose | Audience |Get-Content migrations\0004_add_group_support.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

|------------|----------|---------|----------|

| **Access Token** | 15 minutes | API access | `my-app-client` |```You should see a log similar to: `auth_service listening on :50051 (env=local)`.

| **Refresh Token** | 7 days | New token generation | `my-auth-service` |



### Token Claims (pkg/jwt/token.go)

### 4. Auth Service'i Çalıştırma### 5) Test via Postman (gRPC)

```go

type CustomClaims struct {

    Type string `json:"type"` // "access" or "refresh"

    jwt.RegisteredClaimsProje kök dizininden:1. In Postman, create a new gRPC Request.

}

```2. Server URL: `localhost:50051` (plaintext/TLS off).



**Registered Claims:**```powershell3. Import `proto/auth.proto`.

- `sub` (Subject): User ID

- `jti` (JWT ID): Unique token IDgo run ./cmd/auth_service4. Select `auth.AuthService` and call:

- `iss` (Issuer): `my-auth-service`

- `aud` (Audience): Based on token type```     - `SendOTP` with body `{ "phone_number": "+9055xxxxxxx" }`

- `iat` (Issued At): Creation time

- `exp` (Expires At): Expiration time     - `VerifyOTP` with body `{ "phone_number": "+9055xxxxxxx", "otp_code": "123456" }`



### Refresh Token RotationÇıktı: `auth_service listening on :50051 (env=local)`



When a refresh token is used:On success, `VerifyOTP` returns `access_token` and `refresh_token`.

1. Token is validated and hash is checked in DB

2. Current session is revoked (`revoked_at` is set)### 5. Postman ile gRPC Test

3. New access + refresh token pair is generated

4. New refresh token hash is stored in DB5a) Token utilities via gRPC

5. **Previous refresh token cannot be reused** (one-time use)

1. **Postman'de yeni gRPC Request oluşturun**

### Session Revocation

2. **Server URL**: `localhost:50051` (plaintext/TLS kapalı)- `ValidateToken` with body `{ "access_token": "<ACCESS>" }` → returns `{ is_valid, user_id }` (no error for invalid; just `is_valid=false`).

**Single device logout:**

```3. **Proto dosyasını import edin**: `proto/auth.proto`- `RefreshToken` with body `{ "refresh_token": "<REFRESH>" }` → returns `{ access_token, refresh_token }` (refresh token rotasyonu etkin).

RevokeCurrentDevice(refresh_token) → user_devices.revoked_at = NOW()

```4. **AuthService metodlarını test edin**:



**All devices logout:**5b) Revoke sessions via gRPC

```

LogoutAllDevices(access_token) → Revoke all user's devices**SendOTP örneği:**

```

```json- `RevokeCurrentDevice` with body `{ "refresh_token": "<REFRESH>" }` → returns `{ success: true }`. Afterwards, the same refresh token can no longer be used.

### JWT Interceptor Middleware

{- `LogoutAllDevices` with body `{ "access_token": "<ACCESS>" }` → returns `{ success: true }`. Afterwards, any existing refresh tokens for that user are invalidated (server checks DB-stored hashes and sees they are revoked).

`internal/auth/middleware/auth_interceptor.go`:

- gRPC unary interceptor  "phone_number": "+905551234567"

- Validates access token from `Authorization: Bearer <token>` header

- Injects `user_id` into context}## Notes: Local vs Docker run

- AuthService methods are exempted by default (public)

```

**Usage in other services:**

- Local app run (recommended for quick testing):

```go

tokenManager, _ := jwt.NewTokenManager(jwtSecret, 15*time.Minute, 7*24*time.Hour)**VerifyOTP örneği:**    - App connects to Dockerized Postgres via `localhost:5432`.

server := grpc.NewServer(

    grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor(tokenManager)),```json    - `.env.local` already contains `DATABASE_URL=...@localhost:5432/...`.

)

```{    - Start with `go run ./cmd/auth_service`.



## 🗄️ Database Schema  "phone_number": "+905551234567",



### Migrations  "otp_code": "123456"- Docker app run (Compose):



| File | Description |}    - App connects via the Compose network using host `postgres`.

|------|-------------|

| `0000_enable_uuid.sql` | Enable UUID extension |```    - `docker-compose up --build` will build and run `auth_service` against `postgres`.

| `0001_initial_schema.up.sql` | Base schema (users, contacts, etc.) |

| `0002_user_devices_revocation.up.sql` | Device management and refresh token revocation |    - `docker-compose` passes `.env.docker` to the container.

| `0003_chat_schema.up.sql` | Chat rooms and messages |

| `0004_add_group_support.up.sql` | Group chat support |Başarılı yanıt `access_token` ve `refresh_token` döner.



### Main Tables## Troubleshooting



**users****ValidateToken örneği:**

- `id` (UUID, PK)

- `phone_number` (UNIQUE)```json- `unknown driver "postgres"`: Make sure the project has `github.com/lib/pq` and the driver is blank-imported (already added in `pkg/database/postgres.go`).

- `display_name`

- `profile_picture_url`{- `relation "..." does not exist`: Ensure you applied migrations to the exact database your service is using.

- `created_at`, `updated_at`

  "access_token": "<ACCESS_TOKEN>"- `function uuid_generate_v4() does not exist`: Run `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` on the target database, or switch to `pgcrypto` + `gen_random_uuid()` in schema.

**user_devices**

- `id` (UUID, PK)}

- `user_id` (FK → users)

- `refresh_token_hash` (SHA-256 hash)```## Services

- `device_name`, `device_type`

- `push_notification_token`

- `last_login_at`

- `revoked_at` (for session revocation)**RefreshToken örneği:**- **API Gateway**: The single entry point for all client requests.

- **UNIQUE constraint**: `(user_id, refresh_token_hash)`

```json- **Auth Service**: Handles user authentication, registration, and session management.

**chat_rooms**

- `id` (UUID, PK){- **Chat Service**: Manages chat rooms and user memberships.

- `name`

- `is_group`  "refresh_token": "<REFRESH_TOKEN>"- **Message Worker**: Asynchronously processes and stores messages from a queue.

- `created_at`, `updated_at`

}- **Realtime Service**: Manages WebSocket connections for real-time communication.

**messages**

- `id` (UUID, PK)```- **Status Service**: Handles user status updates (stories).

- `chat_room_id` (FK → chat_rooms)

- `sender_id` (FK → users)

- `content`

- `message_type` (text, image, video, etc.)**RevokeCurrentDevice örneği:**- Revocation and middleware (overview)

- `status` (sent, delivered, read)

- `created_at````json    - Refresh tokens: We added a migration (`0002_user_devices_revocation.up.sql`) to support storing refresh token hashes and revocation timestamps in `user_devices`.



## 🛠️ Technology Stack{    - Middleware: `internal/auth/middleware/auth_interceptor.go` bir gRPC unary interceptor sağlar; `authorization: Bearer <token>` başlığından access token doğrular, `user_id`’yi context’e ekler.



### Backend  "refresh_token": "<REFRESH_TOKEN>"    - AuthService içinde interceptor varsayılan olarak AuthService RPC’lerini muaf tutar (public uçlar). Diğer servislerde enable etmek için:

- **Language**: Go 1.24

- **gRPC**: Inter-service communication}        - Aynı `JWT_SECRET` ile bir `TokenManager` oluşturun (örn. 15dk access / 7gün refresh).

- **Protocol Buffers**: Data serialization

```        - gRPC server’ı `grpc.UnaryInterceptor(UnaryAuthInterceptor(tm))` ile oluşturun.

### Database

- **PostgreSQL 13**: Main data store        - Gerekirse `info.FullMethod` ile public uçları muaf tutabilirsiniz.

- **UUID**: For primary keys

**LogoutAllDevices örneği:**

### Authentication

- **JWT**: golang-jwt/jwt v5```json    - Applying the revocation migration:

- **Twilio Verify**: OTP verification

{        ```powershell

### Messaging (Planned)

- **Apache Kafka**: Asynchronous message queue  "access_token": "<ACCESS_TOKEN>"        # Local Postgres



### Configuration & Logging}        psql -U postgres -h localhost -d whatsapp_clone_dev -f .\migrations\0002_user_devices_revocation.up.sql

- **Viper**: Configuration management

- **Zap**: Structured logging```        ```

- **godotenv**: Environment variables

    - Next steps to fully wire revocation:

### Containerization

- **Docker**: Container runtime## 📦 Mikroservisler        - On VerifyOTP: hash the refresh token (e.g., SHA-256), store in `user_devices` with `last_login_at=NOW()`.

- **Docker Compose**: Multi-service orchestration

        - On RefreshToken: compute the hash and ensure an active (revoked_at IS NULL) record exists for that user and hash; otherwise return Unauthenticated.

## 📝 Development Notes

### Auth Service (Aktif ✅)        - Optionally rotate refresh token and update storage; return the new refresh token in the response (requires proto change).

### Local vs Docker Running

- **Port**: 50051 (gRPC)

**Local (Recommended - For Quick Testing):**- **Sorumluluklar**:

- Application runs locally, connects to Dockerized PostgreSQL  - Twilio OTP ile telefon numarası doğrulama

- Uses `.env.local`: `DATABASE_URL=...@localhost:5433/...`  - JWT (access + refresh token) üretimi

- Start: `go run ./cmd/auth_service`  - Token validasyonu ve yenileme (rotation)

  - Cihaz bazlı session yönetimi

**Docker (Production-like):**  - Tek cihaz veya tüm cihazlar için logout

- Both application and PostgreSQL run in containers- **Teknolojiler**: gRPC, JWT, Twilio, PostgreSQL

- `docker-compose up --build`- **Middleware**: JWT interceptor (diğer servislere taşınabilir)

- Uses `.env.docker`: `DATABASE_URL=...@postgres:5432/...`

### Chat Service (Geliştirme Aşamasında 🚧)

### Protocol Buffer Code Generation- Chat odası yönetimi

- Kullanıcı üyelik yönetimi

Regenerate if you modified proto files:- Grup desteği



```powershell### Realtime Service (Geliştirme Aşamasında 🚧)

# For auth.proto- WebSocket bağlantı yönetimi

protoc --go_out=. --go_opt=paths=source_relative `- Gerçek zamanlı mesaj iletimi

       --go-grpc_out=. --go-grpc_opt=paths=source_relative `- Presence (çevrimiçi durum) yönetimi

       proto/auth.proto

### Status Service (Geliştirme Aşamasında 🚧)

# For all proto files- Kullanıcı hikayeleri (stories)

Get-ChildItem proto\*.proto | ForEach-Object {- Durum güncellemeleri

    protoc --go_out=. --go_opt=paths=source_relative `

           --go-grpc_out=. --go-grpc_opt=paths=source_relative `### Message Worker (Geliştirme Aşamasında 🚧)

           $_.FullName- Kafka consumer

}- Asenkron mesaj işleme

```

### API Gateway (Geliştirme Aşamasında 🚧)

### Running Tests- İstemci istekleri için tek giriş noktası

- gRPC client'lar ile servislere yönlendirme

```powershell

# All tests## 🔐 Kimlik Doğrulama Modeli

go test ./...

Sistem her login için iki JWT token üretir:

# Specific package

go test ./internal/auth/service/...### Token Tipleri



# Verbose output| Token Type | Geçerlilik Süresi | Kullanım Amacı | Audience |

go test -v ./pkg/jwt/...|------------|-------------------|----------------|----------|

```| **Access Token** | 15 minutes | API access | `my-app-client` |

| **Refresh Token** | 7 days | New token generation | `my-auth-service` |

## 🐛 Troubleshooting

### Token Claims (pkg/jwt/token.go)

### "unknown driver postgres"

✅ **Solution**: `github.com/lib/pq` is already blank imported in `pkg/database/postgres.go`.```go

type CustomClaims struct {

### "relation '...' does not exist"    Type string `json:"type"` // "access" or "refresh"

✅ **Solution**: Ensure migrations are applied to the correct database:    jwt.RegisteredClaims

```powershell}

.\scripts\setup-db.ps1```

```

**Registered Claims:**

### "function uuid_generate_v4() does not exist"- `sub` (Subject): User ID

✅ **Solution**: Enable UUID extension:- `jti` (JWT ID): Unique token ID

```sql- `iss` (Issuer): `my-auth-service`

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";- `aud` (Audience): Based on token type

```- `iat` (Issued At): Creation time

Or use `pgcrypto` with `gen_random_uuid()`.- `exp` (Expires At): Expiration time



### Postman: "Message violates its Protobuf type definition"### Refresh Token Rotation

✅ **Solution**: 

- Ensure field names match the proto file (`access_token` not `accessToken`)When a refresh token is used:

- Re-import `proto/auth.proto` file1. Token is validated and hash is checked in DB

2. Current session is revoked (`revoked_at` is set)

### JWT "Unauthenticated" error3. New access + refresh token pair is generated

✅ **Solution**:4. New refresh token hash is stored in DB

- Ensure `JWT_SECRET` is the same in both environments5. **Previous refresh token cannot be reused** (one-time use)

- Don't use access token at refresh endpoint (token type check exists)

- Ensure token hasn't expired### Session Revocation



### Docker Postgres connection error**Single device logout:**

✅ **Solution**:```

```powershellRevokeCurrentDevice(refresh_token) → user_devices.revoked_at = NOW()

# Check container status```

docker-compose ps

**All devices logout:**

# View logs```

docker-compose logs postgresLogoutAllDevices(access_token) → Revoke all user's devices

```

# Check healthcheck status

docker inspect goapp-postgres-1 | Select-String -Pattern "Health"### JWT Interceptor Middleware

```

`internal/auth/middleware/auth_interceptor.go`:

## 🚀 Roadmap- gRPC unary interceptor

- Validates access token from `Authorization: Bearer <token>` header

### Short Term- Injects `user_id` into context

- [x] Auth Service gRPC API- AuthService methods are exempted by default (public)

- [x] JWT token management (access + refresh)

- [x] Refresh token rotation**Usage in other services:**

- [x] Session revocation (single device / all devices)

- [x] JWT interceptor middleware```go

- [ ] Auth Service unit teststokenManager, _ := jwt.NewTokenManager(jwtSecret, 15*time.Minute, 7*24*time.Hour)

- [ ] API Gateway implementationserver := grpc.NewServer(

- [ ] Chat Service basic API    grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor(tokenManager)),

)

### Medium Term```

- [ ] Realtime Service WebSocket connections

- [ ] Status Service story features## 🗄️ Database Schema

- [ ] Kafka integration and Message Worker

- [ ] Message encryption (E2E)### Migrations

- [ ] File upload and storage

- [ ] Push notification integration| File | Description |

|------|-------------|

### Long Term| `0000_enable_uuid.sql` | Enable UUID extension |

- [ ] Metrics and monitoring (Prometheus)| `0001_initial_schema.up.sql` | Base schema (users, contacts, etc.) |

- [ ] Distributed tracing (Jaeger)| `0002_user_devices_revocation.up.sql` | Device management and refresh token revocation |

- [ ] Rate limiting and DDoS protection| `0003_chat_schema.up.sql` | Chat rooms and messages |

- [ ] Multi-region deployment| `0004_add_group_support.up.sql` | Group chat support |

- [ ] CI/CD pipeline

### Main Tables

## 📚 More Information

**users**

For more detailed backend documentation: [`docs/Backend.md`](docs/Backend.md)- `id` (UUID, PK)

- `phone_number` (UNIQUE)

## 📄 License- `display_name`

- `profile_picture_url`

This project is for educational purposes.- `created_at`, `updated_at`



## 👨‍💻 Contributing**user_devices**

- `id` (UUID, PK)

1. Fork the repository- `user_id` (FK → users)

2. Create a feature branch (`git checkout -b feature/amazing-feature`)- `refresh_token_hash` (SHA-256 hash)

3. Commit your changes (`git commit -m 'Add amazing feature'`)- `device_name`, `device_type`

4. Push to the branch (`git push origin feature/amazing-feature`)- `push_notification_token`

5. Create a Pull Request- `last_login_at`

- `revoked_at` (session revocation için)

---- **UNIQUE constraint**: `(user_id, refresh_token_hash)`



**Note**: This project is under active development. Some features are not yet complete.**chat_rooms**

- `id` (UUID, PK)
- `name`
- `is_group`
- `created_at`, `updated_at`

**messages**
- `id` (UUID, PK)
- `chat_room_id` (FK → chat_rooms)
- `sender_id` (FK → users)
- `content`
- `message_type` (text, image, video, etc.)
- `status` (sent, delivered, read)
- `created_at`

## 🛠️ Technology Stack

### Backend
- **Language**: Go 1.24
- **gRPC**: Inter-service communication
- **Protocol Buffers**: Data serialization

### Database
- **PostgreSQL 13**: Main data store
- **UUID**: For primary keys

### Authentication
- **JWT**: golang-jwt/jwt v5
- **Twilio Verify**: OTP verification

### Messaging (Planned)
- **Apache Kafka**: Asynchronous message queue

### Configuration & Logging
- **Viper**: Configuration management
- **Zap**: Structured logging
- **godotenv**: Environment variables

### Containerization
- **Docker**: Container runtime
- **Docker Compose**: Multi-service orchestration

## 📝 Development Notes

### Local vs Docker Running

**Local (Recommended - For Quick Testing):**
- Application runs locally, connects to Dockerized PostgreSQL
- Uses `.env.local`: `DATABASE_URL=...@localhost:5433/...`
- Start: `go run ./cmd/auth_service`

**Docker (Production-like):**
- Both application and PostgreSQL run in containers
- `docker-compose up --build`
- Uses `.env.docker`: `DATABASE_URL=...@postgres:5432/...`

### Protocol Buffer Code Generation

Regenerate if you modified proto files:

```powershell
# For auth.proto
protoc --go_out=. --go_opt=paths=source_relative `
       --go-grpc_out=. --go-grpc_opt=paths=source_relative `
       proto/auth.proto

# For all proto files
Get-ChildItem proto\*.proto | ForEach-Object {
    protoc --go_out=. --go_opt=paths=source_relative `
           --go-grpc_out=. --go-grpc_opt=paths=source_relative `
           $_.FullName
}
```

### Running Tests

```powershell
# All tests
go test ./...

# Specific package
go test ./internal/auth/service/...

# Verbose output
go test -v ./pkg/jwt/...
```

## 🐛 Troubleshooting

### "unknown driver postgres"
✅ **Solution**: `github.com/lib/pq` is already blank imported in `pkg/database/postgres.go`.

### "relation '...' does not exist"
✅ **Solution**: Ensure migrations are applied to the correct database:
```powershell
.\scripts\setup-db.ps1
```

### "function uuid_generate_v4() does not exist"
✅ **Solution**: Enable UUID extension:
```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```
Or use `pgcrypto` with `gen_random_uuid()`.

### Postman: "Message violates its Protobuf type definition"
✅ **Solution**: 
- Ensure field names match the proto file (`access_token` not `accessToken`)
- Re-import `proto/auth.proto` file

### JWT "Unauthenticated" error
✅ **Solution**:
- Ensure `JWT_SECRET` is the same in both environments
- Don't use access token at refresh endpoint (token type check exists)
- Ensure token hasn't expired

### Docker Postgres connection error
✅ **Solution**:
```powershell
# Check container status
docker-compose ps

# View logs
docker-compose logs postgres

# Check healthcheck status
docker inspect goapp-postgres-1 | Select-String -Pattern "Health"
```

## 🚀 Roadmap

### Short Term
- [x] Auth Service gRPC API
- [x] JWT token management (access + refresh)
- [x] Refresh token rotation
- [x] Session revocation (single device / all devices)
- [x] JWT interceptor middleware
- [ ] Auth Service unit tests
- [ ] API Gateway implementation
- [ ] Chat Service basic API

### Medium Term
- [ ] Realtime Service WebSocket connections
- [ ] Status Service story features
- [ ] Kafka integration and Message Worker
- [ ] Message encryption (E2E)
- [ ] File upload and storage
- [ ] Push notification integration

### Long Term
- [ ] Metrics and monitoring (Prometheus)
- [ ] Distributed tracing (Jaeger)
- [ ] Rate limiting and DDoS protection
- [ ] Multi-region deployment
- [ ] CI/CD pipeline

## 📚 More Information

For more detailed backend documentation: [`docs/Backend.md`](docs/Backend.md)

## 📄 License

This project is for educational purposes.

## 👨‍💻 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Create a Pull Request

---

**Note**: This project is under active development. Some features are not yet complete.
