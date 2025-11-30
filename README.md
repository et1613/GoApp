# WhatsApp Clone Backend# WhatsApp Clone Backend



[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)This project is a Go-based backend for a WhatsApp-like application, built with a microservices architecture.

[![gRPC](https://img.shields.io/badge/gRPC-Protocol%20Buffers-green.svg)](https://grpc.io)

[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue.svg)](https://www.postgresql.org)## Project Structure



Go tabanlı, mikroservis mimarisi ile geliştirilmiş bir WhatsApp benzeri uygulama backend'idir. JWT tabanlı kimlik doğrulama, gerçek zamanlı mesajlaşma ve gRPC iletişim protokolü içerir.The project follows a clean architecture and microservices pattern.



## 🏗️ Proje Yapısı- `/cmd`: Entry points for each service (`main.go`).

- `/internal`: Private application and business logic for each service.

Proje Clean Architecture ve mikroservis desenini takip eder:- `/pkg`: Shared libraries and domain types used across services.

- `/proto`: gRPC protocol definitions for inter-service communication.

```- `/migrations`: Database schema migrations.

GoApp/- `/configs`: Configuration files for different environments.

├── cmd/                    # Servis giriş noktaları

│   ├── auth_service/      # Kimlik doğrulama servisi## Getting Started

│   ├── chat_service/      # Chat yönetim servisi

│   ├── realtime_service/  # WebSocket gerçek zamanlı iletişim### Prerequisites

│   ├── status_service/    # Kullanıcı durumu (hikaye) servisi

│   ├── message_worker/    # Mesaj kuyruğu işleyici- Docker and Docker Compose

│   └── api_gateway/       # API Gateway (geliştirme aşamasında)- Go 1.18 or higher

├── internal/              # Özel uygulama ve iş mantığı

│   ├── auth/             # Auth domain logic### Running the application

│   │   ├── handler/      # gRPC handlers

│   │   ├── service/      # İş mantığı katmanı1.  **Start the infrastructure:**

│   │   ├── repository/   # Veri erişim arayüzleri    ```bash

│   │   ├── store/        # PostgreSQL implementasyonları    docker-compose up -d

│   │   └── middleware/   # JWT interceptor    ```

│   ├── chat/             # Chat domain logic

│   ├── realtime/         # WebSocket hub ve handlers2.  **Run database migrations:**

│   └── worker/           # Kafka consumer    You'll need a migration tool like `golang-migrate/migrate`.

├── pkg/                   # Paylaşılan kütüphaneler    ```bash

│   ├── config/           # Yapılandırma yönetimi (Viper)    migrate -database "postgres://postgres:Fbtex1967.@localhost:5432/whatsapp_clone_dev?sslmode=disable" -path migrations up

│   ├── database/         # PostgreSQL bağlantı yönetimi    ```

│   ├── domain/           # Domain modelleri (User, Message, vb.)

│   ├── jwt/              # JWT token yönetimi3.  **Run the services:**

│   ├── logger/           # Zap logger    Navigate to each service's directory and run it.

│   └── eventbus/         # Kafka client    ```bash

├── proto/                 # gRPC Protocol Buffer tanımları    go run ./cmd/api_gateway/

├── migrations/            # PostgreSQL şema migrasyonları    go run ./cmd/auth_service/

├── scripts/              # Yardımcı PowerShell scriptleri    # ... and so on for other services

└── docker-compose.yml    # Docker Compose yapılandırması    ```

```

## Quickstart: AuthService + Postman gRPC (Windows)

## 🚀 Hızlı Başlangıç

End-to-end minimum setup to test OTP and JWT issuance via Postman using gRPC.

### Ön Gereksinimler

### 1) Environment variables

- **Docker & Docker Compose** (PostgreSQL için)

- **Go 1.24 veya üstü**This service auto-loads environment files based on context:

- **PowerShell** (Windows için script çalıştırma)

- **Postman** (gRPC test için, opsiyonel)- Local runs: loads `.env.local` (we added one with sensible defaults)

- Docker runs: `docker-compose` passes `.env.docker` and also sets `RUNNING_IN_DOCKER=true` (the app attempts to load `.env.docker` but works even if the file isn't baked into the image)

### 1. Ortam Değişkenlerini Ayarlama- Base `.env` is also loaded last if present (for overrides)



Proje farklı ortamlar için otomatik .env dosyası yükler:Example files provided:

- Local çalıştırma: `.env.local`

- Docker içinde: `.env.docker` (RUNNING_IN_DOCKER=true ile)- `.env.example` — connects to Postgres at `localhost:5432`, sets a sample `JWT_SECRET`, and enables `AUTH_DEV_MODE=true` (Twilio bypass: OTP code is `123456`).

- `.env.docker.example` — same but using `postgres:5432` for Compose and `RUNNING_IN_DOCKER=true`.

Başlamak için örnek dosyayı kopyalayın:

Create a copy for your environment:

```powershell

Copy-Item .env.example .env.local```powershell

```Copy-Item .env.example .env

```

**Önemli değişkenler:**

- `DATABASE_URL`: PostgreSQL bağlantı dizesiEdit `.env` as needed:

- `JWT_SECRET`: JWT imzalama için güçlü bir secret (minimum 32 byte)

- `AUTH_DEV_MODE=true`: Twilio bypass (OTP kodu her zaman `123456`)- Set `JWT_SECRET` to a strong value

- Twilio (Production için): `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, `TWILIO_VERIFY_SERVICE_SID`- Add real Twilio credentials if you want to use OTP verification against Twilio (otherwise keep `AUTH_DEV_MODE=true`).



### 2. PostgreSQL Başlatma### 2) Start PostgreSQL (Docker)



Docker Compose ile sadece PostgreSQL'i başlatın:Run only Postgres in the background:



```powershell```powershell

docker-compose up -d postgresdocker-compose up -d postgres

``````



PostgreSQL `5433` portunda çalışacaktır (yerel PostgreSQL ile çakışmayı önlemek için).### 3) Apply database schema (migrations)



### 3. Veritabanı Migrasyonlarını UygulamaOption A: Use the automated script (recommended):



**Otomatik yöntem (önerilen):**```powershell

Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass

```powershell.\scripts\setup-db.ps1

Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass```

.\scripts\setup-db.ps1

```Option B: Manual migration:



**Manuel yöntem:**```powershell

Get-Content -Raw migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -f -

```powershellGet-Content -Raw migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

# UUID extension'ı etkinleştirGet-Content -Raw migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

Get-Content migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_devGet-Content -Raw migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

```

# Ana şema

Get-Content migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1If you previously ran into an inline index syntax error for `call_logs`, this repository already fixes it by creating the index separately.



# Cihaz yönetimi ve revocation### 4) Run Auth Service locally

Get-Content migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

From the project root:

# Chat şeması

Get-Content migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1```powershell

go run ./cmd/auth_service

# Grup desteği```

Get-Content migrations\0004_add_group_support.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

```You should see a log similar to: `auth_service listening on :50051 (env=local)`.



### 4. Auth Service'i Çalıştırma### 5) Test via Postman (gRPC)



Proje kök dizininden:1. In Postman, create a new gRPC Request.

2. Server URL: `localhost:50051` (plaintext/TLS off).

```powershell3. Import `proto/auth.proto`.

go run ./cmd/auth_service4. Select `auth.AuthService` and call:

```     - `SendOTP` with body `{ "phone_number": "+9055xxxxxxx" }`

     - `VerifyOTP` with body `{ "phone_number": "+9055xxxxxxx", "otp_code": "123456" }`

Çıktı: `auth_service listening on :50051 (env=local)`

On success, `VerifyOTP` returns `access_token` and `refresh_token`.

### 5. Postman ile gRPC Test

5a) Token utilities via gRPC

1. **Postman'de yeni gRPC Request oluşturun**

2. **Server URL**: `localhost:50051` (plaintext/TLS kapalı)- `ValidateToken` with body `{ "access_token": "<ACCESS>" }` → returns `{ is_valid, user_id }` (no error for invalid; just `is_valid=false`).

3. **Proto dosyasını import edin**: `proto/auth.proto`- `RefreshToken` with body `{ "refresh_token": "<REFRESH>" }` → returns `{ access_token, refresh_token }` (refresh token rotasyonu etkin).

4. **AuthService metodlarını test edin**:

5b) Revoke sessions via gRPC

**SendOTP örneği:**

```json- `RevokeCurrentDevice` with body `{ "refresh_token": "<REFRESH>" }` → returns `{ success: true }`. Afterwards, the same refresh token can no longer be used.

{- `LogoutAllDevices` with body `{ "access_token": "<ACCESS>" }` → returns `{ success: true }`. Afterwards, any existing refresh tokens for that user are invalidated (server checks DB-stored hashes and sees they are revoked).

  "phone_number": "+905551234567"

}## Notes: Local vs Docker run

```

- Local app run (recommended for quick testing):

**VerifyOTP örneği:**    - App connects to Dockerized Postgres via `localhost:5432`.

```json    - `.env.local` already contains `DATABASE_URL=...@localhost:5432/...`.

{    - Start with `go run ./cmd/auth_service`.

  "phone_number": "+905551234567",

  "otp_code": "123456"- Docker app run (Compose):

}    - App connects via the Compose network using host `postgres`.

```    - `docker-compose up --build` will build and run `auth_service` against `postgres`.

    - `docker-compose` passes `.env.docker` to the container.

Başarılı yanıt `access_token` ve `refresh_token` döner.

## Troubleshooting

**ValidateToken örneği:**

```json- `unknown driver "postgres"`: Make sure the project has `github.com/lib/pq` and the driver is blank-imported (already added in `pkg/database/postgres.go`).

{- `relation "..." does not exist`: Ensure you applied migrations to the exact database your service is using.

  "access_token": "<ACCESS_TOKEN>"- `function uuid_generate_v4() does not exist`: Run `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` on the target database, or switch to `pgcrypto` + `gen_random_uuid()` in schema.

}

```## Services



**RefreshToken örneği:**- **API Gateway**: The single entry point for all client requests.

```json- **Auth Service**: Handles user authentication, registration, and session management.

{- **Chat Service**: Manages chat rooms and user memberships.

  "refresh_token": "<REFRESH_TOKEN>"- **Message Worker**: Asynchronously processes and stores messages from a queue.

}- **Realtime Service**: Manages WebSocket connections for real-time communication.

```- **Status Service**: Handles user status updates (stories).



**RevokeCurrentDevice örneği:**- Revocation and middleware (overview)

```json    - Refresh tokens: We added a migration (`0002_user_devices_revocation.up.sql`) to support storing refresh token hashes and revocation timestamps in `user_devices`.

{    - Middleware: `internal/auth/middleware/auth_interceptor.go` bir gRPC unary interceptor sağlar; `authorization: Bearer <token>` başlığından access token doğrular, `user_id`’yi context’e ekler.

  "refresh_token": "<REFRESH_TOKEN>"    - AuthService içinde interceptor varsayılan olarak AuthService RPC’lerini muaf tutar (public uçlar). Diğer servislerde enable etmek için:

}        - Aynı `JWT_SECRET` ile bir `TokenManager` oluşturun (örn. 15dk access / 7gün refresh).

```        - gRPC server’ı `grpc.UnaryInterceptor(UnaryAuthInterceptor(tm))` ile oluşturun.

        - Gerekirse `info.FullMethod` ile public uçları muaf tutabilirsiniz.

**LogoutAllDevices örneği:**

```json    - Applying the revocation migration:

{        ```powershell

  "access_token": "<ACCESS_TOKEN>"        # Local Postgres

}        psql -U postgres -h localhost -d whatsapp_clone_dev -f .\migrations\0002_user_devices_revocation.up.sql

```        ```

    - Next steps to fully wire revocation:

## 📦 Mikroservisler        - On VerifyOTP: hash the refresh token (e.g., SHA-256), store in `user_devices` with `last_login_at=NOW()`.

        - On RefreshToken: compute the hash and ensure an active (revoked_at IS NULL) record exists for that user and hash; otherwise return Unauthenticated.

### Auth Service (Aktif ✅)        - Optionally rotate refresh token and update storage; return the new refresh token in the response (requires proto change).

- **Port**: 50051 (gRPC)
- **Sorumluluklar**:
  - Twilio OTP ile telefon numarası doğrulama
  - JWT (access + refresh token) üretimi
  - Token validasyonu ve yenileme (rotation)
  - Cihaz bazlı session yönetimi
  - Tek cihaz veya tüm cihazlar için logout
- **Teknolojiler**: gRPC, JWT, Twilio, PostgreSQL
- **Middleware**: JWT interceptor (diğer servislere taşınabilir)

### Chat Service (Geliştirme Aşamasında 🚧)
- Chat odası yönetimi
- Kullanıcı üyelik yönetimi
- Grup desteği

### Realtime Service (Geliştirme Aşamasında 🚧)
- WebSocket bağlantı yönetimi
- Gerçek zamanlı mesaj iletimi
- Presence (çevrimiçi durum) yönetimi

### Status Service (Geliştirme Aşamasında 🚧)
- Kullanıcı hikayeleri (stories)
- Durum güncellemeleri

### Message Worker (Geliştirme Aşamasında 🚧)
- Kafka consumer
- Asenkron mesaj işleme

### API Gateway (Geliştirme Aşamasında 🚧)
- İstemci istekleri için tek giriş noktası
- gRPC client'lar ile servislere yönlendirme

## 🔐 Kimlik Doğrulama Modeli

Sistem her login için iki JWT token üretir:

### Token Tipleri

| Token Type | Geçerlilik Süresi | Kullanım Amacı | Audience |
|------------|-------------------|----------------|----------|
| **Access Token** | 15 dakika | API erişimi için | `my-app-client` |
| **Refresh Token** | 7 gün | Yeni token üretimi için | `my-auth-service` |

### Token Claims (pkg/jwt/token.go)

```go
type CustomClaims struct {
    Type string `json:"type"` // "access" veya "refresh"
    jwt.RegisteredClaims
}
```

**Registered Claims:**
- `sub` (Subject): User ID
- `jti` (JWT ID): Benzersiz token ID
- `iss` (Issuer): `my-auth-service`
- `aud` (Audience): Token tipine göre
- `iat` (Issued At): Oluşturma zamanı
- `exp` (Expires At): Son kullanma tarihi

### Refresh Token Rotation

Refresh token kullanıldığında:
1. Token validasyonu yapılır ve DB'de hash kontrol edilir
2. Mevcut session revoke edilir (`revoked_at` set edilir)
3. Yeni access + refresh token çifti üretilir
4. Yeni refresh token hash'i DB'ye kaydedilir
5. **Önceki refresh token artık kullanılamaz** (tek kullanımlık)

### Session Revocation

**Tek cihaz logout:**
```
RevokeCurrentDevice(refresh_token) → user_devices.revoked_at = NOW()
```

**Tüm cihazlar logout:**
```
LogoutAllDevices(access_token) → Kullanıcının tüm cihazlarını revoke et
```

### JWT Interceptor Middleware

`internal/auth/middleware/auth_interceptor.go`:
- gRPC unary interceptor
- `Authorization: Bearer <token>` header'ından access token doğrular
- `user_id`'yi context'e ekler
- AuthService metodları varsayılan olarak muaf tutulur (public)

**Diğer servislerde kullanım:**

```go
tokenManager, _ := jwt.NewTokenManager(jwtSecret, 15*time.Minute, 7*24*time.Hour)
server := grpc.NewServer(
    grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor(tokenManager)),
)
```

## 🗄️ Veritabanı Şeması

### Migrasyonlar

| Dosya | Açıklama |
|-------|----------|
| `0000_enable_uuid.sql` | UUID extension'ı etkinleştir |
| `0001_initial_schema.up.sql` | Temel şema (users, contacts, vb.) |
| `0002_user_devices_revocation.up.sql` | Cihaz yönetimi ve refresh token revocation |
| `0003_chat_schema.up.sql` | Chat odaları ve mesajlar |
| `0004_add_group_support.up.sql` | Grup chat desteği |

### Temel Tablolar

**users**
- `id` (UUID, PK)
- `phone_number` (UNIQUE)
- `display_name`
- `profile_picture_url`
- `created_at`, `updated_at`

**user_devices**
- `id` (UUID, PK)
- `user_id` (FK → users)
- `refresh_token_hash` (SHA-256 hash)
- `device_name`, `device_type`
- `push_notification_token`
- `last_login_at`
- `revoked_at` (session revocation için)
- **UNIQUE constraint**: `(user_id, refresh_token_hash)`

**chat_rooms**
- `id` (UUID, PK)
- `name`
- `is_group`
- `created_at`, `updated_at`

**messages**
- `id` (UUID, PK)
- `chat_room_id` (FK → chat_rooms)
- `sender_id` (FK → users)
- `content`
- `message_type` (text, image, video, vb.)
- `status` (sent, delivered, read)
- `created_at`

## 🛠️ Teknoloji Yığını

### Backend
- **Dil**: Go 1.24
- **gRPC**: Servisler arası iletişim
- **Protocol Buffers**: Veri serileştirme

### Veritabanı
- **PostgreSQL 13**: Ana veri deposu
- **UUID**: Birincil anahtarlar için

### Kimlik Doğrulama
- **JWT**: golang-jwt/jwt v5
- **Twilio Verify**: OTP doğrulaması

### Mesajlaşma (Planlı)
- **Apache Kafka**: Asenkron mesaj kuyruğu

### Yapılandırma & Logging
- **Viper**: Yapılandırma yönetimi
- **Zap**: Yapılandırılmış logging
- **godotenv**: Ortam değişkenleri

### Containerization
- **Docker**: Container runtime
- **Docker Compose**: Çoklu servis orkestasyonu

## 📝 Geliştirme Notları

### Local vs Docker Çalıştırma

**Local (Önerilen - Hızlı Test İçin):**
- Uygulama local'de çalışır, Dockerized PostgreSQL'e bağlanır
- `.env.local` kullanılır: `DATABASE_URL=...@localhost:5433/...`
- Başlatma: `go run ./cmd/auth_service`

**Docker (Production-like):**
- Hem uygulama hem PostgreSQL container'da çalışır
- `docker-compose up --build`
- `.env.docker` kullanılır: `DATABASE_URL=...@postgres:5432/...`

### Protocol Buffer Kod Üretimi

Proto dosyalarını değiştirdiyseniz yeniden üretin:

```powershell
# auth.proto için
protoc --go_out=. --go_opt=paths=source_relative `
       --go-grpc_out=. --go-grpc_opt=paths=source_relative `
       proto/auth.proto

# Tüm proto dosyaları için
Get-ChildItem proto\*.proto | ForEach-Object {
    protoc --go_out=. --go_opt=paths=source_relative `
           --go-grpc_out=. --go-grpc_opt=paths=source_relative `
           $_.FullName
}
```

### Test Çalıştırma

```powershell
# Tüm testler
go test ./...

# Spesifik paket
go test ./internal/auth/service/...

# Verbose output
go test -v ./pkg/jwt/...
```

## 🐛 Sorun Giderme

### "unknown driver postgres"
✅ **Çözüm**: `github.com/lib/pq` zaten `pkg/database/postgres.go`'da blank import edilmiş.

### "relation '...' does not exist"
✅ **Çözüm**: Migrasyonları doğru veritabanına uyguladığınızdan emin olun:
```powershell
.\scripts\setup-db.ps1
```

### "function uuid_generate_v4() does not exist"
✅ **Çözüm**: UUID extension'ı etkinleştirin:
```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```
veya `pgcrypto` kullanıp `gen_random_uuid()` tercih edin.

### Postman: "Message violates its Protobuf type definition"
✅ **Çözüm**: 
- Field isimlerinin proto dosyası ile eşleştiğinden emin olun (`access_token` değil `accessToken`)
- `proto/auth.proto` dosyasını tekrar import edin

### JWT "Unauthenticated" hatası
✅ **Çözüm**:
- `JWT_SECRET` değerinin her iki ortamda da aynı olduğundan emin olun
- Access token'ı refresh endpoint'inde kullanmayın (token type kontrolü var)
- Token süresinin dolmadığından emin olun

### Docker Postgres bağlantı hatası
✅ **Çözüm**:
```powershell
# Container durumunu kontrol edin
docker-compose ps

# Logları inceleyin
docker-compose logs postgres

# Healthcheck durumunu kontrol edin
docker inspect goapp-postgres-1 | Select-String -Pattern "Health"
```

## 🚀 Yol Haritası

### Kısa Vadeli
- [x] Auth Service gRPC API
- [x] JWT token yönetimi (access + refresh)
- [x] Refresh token rotation
- [x] Session revocation (tek cihaz / tüm cihazlar)
- [x] JWT interceptor middleware
- [ ] Auth Service unit testleri
- [ ] API Gateway implementasyonu
- [ ] Chat Service temel API

### Orta Vadeli
- [ ] Realtime Service WebSocket bağlantıları
- [ ] Status Service hikaye özellikleri
- [ ] Kafka entegrasyonu ve Message Worker
- [ ] Mesaj şifreleme (E2E)
- [ ] Dosya yükleme ve storage
- [ ] Push notification entegrasyonu

### Uzun Vadeli
- [ ] Metrikler ve monitoring (Prometheus)
- [ ] Distributed tracing (Jaeger)
- [ ] Rate limiting ve DDoS koruması
- [ ] Multi-region deployment
- [ ] CI/CD pipeline

## 📚 Daha Fazla Bilgi

Daha detaylı backend dokümantasyonu için: [`docs/Backend.md`](docs/Backend.md)

## 📄 Lisans

Bu proje eğitim amaçlıdır.

## 👨‍💻 Katkıda Bulunma

1. Fork yapın
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Değişikliklerinizi commit edin (`git commit -m 'Add amazing feature'`)
4. Branch'inizi push edin (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

---

**Not**: Bu proje aktif geliştirme aşamasındadır. Bazı özellikler henüz tamamlanmamıştır.
