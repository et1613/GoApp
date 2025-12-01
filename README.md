# WhatsApp Clone - Microservices Backend# WhatsApp Clone Backend# WhatsApp Clone Backend# WhatsApp Clone Backend# WhatsApp Clone Backend



[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)

[![gRPC](https://img.shields.io/badge/gRPC-Protocol%20Buffers-green.svg)](https://grpc.io)

[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue.svg)](https://www.postgresql.org)[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)



A production-ready WhatsApp-like messaging platform backend built with Go, featuring microservices architecture, JWT-based authentication with device management, real-time messaging capabilities, and gRPC inter-service communication.[![gRPC](https://img.shields.io/badge/gRPC-Protocol%20Buffers-green.svg)](https://grpc.io)



## 📋 Table of Contents[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue.svg)](https://www.postgresql.org)[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)



- [Features](#-features)

- [Architecture](#-architecture)

- [Project Structure](#-project-structure)A WhatsApp-like application backend built with Go and microservices architecture. Features JWT-based authentication, real-time messaging, and gRPC communication protocol.[![gRPC](https://img.shields.io/badge/gRPC-Protocol%20Buffers-green.svg)](https://grpc.io)

- [Technology Stack](#-technology-stack)

- [Getting Started](#-getting-started)

- [Microservices](#-microservices)

- [Authentication](#-authentication)## 🏗️ Project Structure[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue.svg)](https://www.postgresql.org)[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)This project is a Go-based backend for a WhatsApp-like application, built with a microservices architecture.

- [Database Schema](#-database-schema)

- [API Documentation](#-api-documentation)

- [Development](#-development)

- [Troubleshooting](#-troubleshooting)The project follows Clean Architecture and microservices pattern:

- [Roadmap](#-roadmap)



## ✨ Features

```A WhatsApp-like application backend built with Go and microservices architecture. Features JWT-based authentication, real-time messaging, and gRPC communication protocol.[![gRPC](https://img.shields.io/badge/gRPC-Protocol%20Buffers-green.svg)](https://grpc.io)

### Implemented ✅

- **JWT Authentication**: Access and refresh token system with rotationGoApp/

- **Device Management**: Multi-device login support with device-specific sessions

- **Session Revocation**: Logout single device or all devices at once├── cmd/                    # Service entry points

- **OTP Verification**: Phone number verification via Twilio (with dev mode bypass)

- **gRPC Communication**: High-performance inter-service communication│   ├── auth_service/      # Authentication service

- **PostgreSQL Database**: Reliable data persistence with UUID primary keys

- **Clean Architecture**: Separation of concerns with layers (handler, service, repository, store)│   ├── chat_service/      # Chat management service## 🏗️ Project Structure[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue.svg)](https://www.postgresql.org)## Project Structure

- **Configuration Management**: Viper-based config with environment-specific loading

- **Structured Logging**: Zap logger integration│   ├── realtime_service/  # WebSocket real-time communication

- **Docker Support**: Containerized services with Docker Compose

│   ├── status_service/    # User status (stories) service

### In Development 🚧

- **Chat Service**: Group and direct messaging│   ├── message_worker/    # Message queue processor

- **Realtime Service**: WebSocket connections for live messaging

- **Status Service**: User stories/status updates│   └── api_gateway/       # API Gateway (in development)The project follows Clean Architecture and microservices pattern:

- **Message Worker**: Kafka-based asynchronous message processing

- **API Gateway**: Unified entry point for client requests├── internal/              # Private application and business logic



## 🏗 Architecture│   ├── auth/             # Auth domain logic



This project follows a **microservices architecture** with **Clean Architecture** principles:│   │   ├── handler/      # gRPC handlers



```│   │   ├── service/      # Business logic layer```Go tabanlı, mikroservis mimarisi ile geliştirilmiş bir WhatsApp benzeri uygulama backend'idir. JWT tabanlı kimlik doğrulama, gerçek zamanlı mesajlaşma ve gRPC iletişim protokolü içerir.The project follows a clean architecture and microservices pattern.

┌─────────────┐

│ API Gateway │ (HTTP/REST) - Entry point for clients│   │   ├── repository/   # Data access interfaces

└──────┬──────┘

       │ gRPC│   │   ├── store/        # PostgreSQL implementationsGoApp/

       ├─────────────┬──────────────┬──────────────┬─────────────┐

       │             │              │              │             ││   │   └── middleware/   # JWT interceptor

   ┌───▼────┐   ┌───▼────┐   ┌─────▼─────┐   ┌───▼────┐   ┌───▼────┐

   │  Auth  │   │  Chat  │   │ Realtime  │   │ Status │   │ Worker ││   ├── chat/             # Chat domain logic├── cmd/                    # Service entry points

   │Service │   │Service │   │  Service  │   │Service │   │ (Kafka)│

   └───┬────┘   └───┬────┘   └─────┬─────┘   └───┬────┘   └───┬────┘│   ├── realtime/         # WebSocket hub and handlers

       │            │              │              │             │

       └────────────┴──────────────┴──────────────┴─────────────┘│   └── worker/           # Kafka consumer│   ├── auth_service/      # Authentication service

                                   │

                            ┌──────▼──────┐├── pkg/                   # Shared libraries

                            │ PostgreSQL  │

                            └─────────────┘│   ├── config/           # Configuration management (Viper)│   ├── chat_service/      # Chat management service## 🏗️ Proje Yapısı- `/cmd`: Entry points for each service (`main.go`).

```

│   ├── database/         # PostgreSQL connection management

### Design Principles

│   ├── domain/           # Domain models (User, Message, etc.)│   ├── realtime_service/  # WebSocket real-time communication

- **Domain-Driven Design**: Each service owns its domain logic

- **Clean Architecture**: Dependency inversion, testable business logic│   ├── jwt/              # JWT token management

- **Protocol Buffers**: Strongly-typed service contracts

- **Database-per-Service**: Each service can have isolated data (currently shared DB)│   ├── logger/           # Zap logger│   ├── status_service/    # User status (stories) service- `/internal`: Private application and business logic for each service.



## 📁 Project Structure│   └── eventbus/         # Kafka client



```├── proto/                 # gRPC Protocol Buffer definitions│   ├── message_worker/    # Message queue processor

GoApp/

├── cmd/                          # Service entry points (main.go files)├── migrations/            # PostgreSQL schema migrations

│   ├── all_in_one/              # Monolithic version (deprecated)

│   ├── api_gateway/             # API Gateway service├── scripts/              # Helper PowerShell scripts│   └── api_gateway/       # API Gateway (in development)Proje Clean Architecture ve mikroservis desenini takip eder:- `/pkg`: Shared libraries and domain types used across services.

│   ├── auth_service/            # Authentication service ✅

│   ├── chat_service/            # Chat management service 🚧└── docker-compose.yml    # Docker Compose configuration

│   ├── message_worker/          # Kafka message consumer 🚧

│   ├── realtime_service/        # WebSocket service 🚧```├── internal/              # Private application and business logic

│   └── status_service/          # Status/stories service 🚧

│

├── internal/                     # Private application code

│   ├── auth/                    # Auth domain## 🚀 Quick Start│   ├── auth/             # Auth domain logic- `/proto`: gRPC protocol definitions for inter-service communication.

│   │   ├── handler/             # gRPC request handlers

│   │   ├── middleware/          # JWT interceptor

│   │   ├── repository/          # Data access interfaces

│   │   ├── service/             # Business logic### Prerequisites│   │   ├── handler/      # gRPC handlers

│   │   └── store/               # PostgreSQL implementations

│   ├── chat/                    # Chat domain

│   ├── realtime/                # WebSocket hub & handlers

│   └── worker/                  # Kafka consumer logic- **Docker & Docker Compose** (for PostgreSQL)│   │   ├── service/      # Business logic layer```- `/migrations`: Database schema migrations.

│

├── pkg/                         # Public shared libraries- **Go 1.24 or higher**

│   ├── config/                  # Viper configuration

│   ├── database/                # PostgreSQL connection pool- **PowerShell** (for running scripts on Windows)│   │   ├── repository/   # Data access interfaces

│   ├── domain/                  # Domain models (User, Message, etc.)

│   ├── eventbus/                # Kafka client- **Postman** (for gRPC testing, optional)

│   ├── jwt/                     # JWT token management

│   └── logger/                  # Zap logger│   │   ├── store/        # PostgreSQL implementationsGoApp/- `/configs`: Configuration files for different environments.

│

├── proto/                       # Protocol Buffer definitions### 1. Environment Variables Setup

│   ├── auth.proto              # Auth service contract

│   ├── chat.proto              # Chat service contract│   │   └── middleware/   # JWT interceptor

│   ├── realtime.proto          # Realtime service contract

│   └── *.pb.go                 # Generated Go codeThe project automatically loads .env files based on context:

│

├── migrations/                  # Database migrations- Local run: `.env.local`│   ├── chat/             # Chat domain logic├── cmd/                    # Servis giriş noktaları

│   ├── 0000_enable_uuid.sql

│   ├── 0001_initial_schema.up.sql- Docker run: `.env.docker` (with RUNNING_IN_DOCKER=true)

│   ├── 0002_user_devices_revocation.up.sql

│   ├── 0003_chat_schema.up.sql│   ├── realtime/         # WebSocket hub and handlers

│   └── 0004_add_group_support.up.sql

│Copy the example file to get started:

├── scripts/                     # Helper scripts

│   └── setup-db.ps1            # Database setup automation│   └── worker/           # Kafka consumer│   ├── auth_service/      # Kimlik doğrulama servisi## Getting Started

│

├── docs/                        # Documentation```powershell

│   └── Backend.md              # Detailed backend docs

│Copy-Item .env.example .env.local├── pkg/                   # Shared libraries

├── docker-compose.yml           # Docker orchestration

├── go.mod                       # Go dependencies```

└── go.work                      # Go workspace

```│   ├── config/           # Configuration management (Viper)│   ├── chat_service/      # Chat yönetim servisi



## 🛠 Technology Stack**Important variables:**



### Backend- `DATABASE_URL`: PostgreSQL connection string│   ├── database/         # PostgreSQL connection management

| Technology | Purpose | Version |

|------------|---------|---------|- `JWT_SECRET`: Strong secret for JWT signing (minimum 32 bytes)

| **Go** | Primary language | 1.24 |

| **gRPC** | Inter-service communication | Latest |- `AUTH_DEV_MODE=true`: Twilio bypass (OTP code is always `123456`)│   ├── domain/           # Domain models (User, Message, etc.)│   ├── realtime_service/  # WebSocket gerçek zamanlı iletişim### Prerequisites

| **Protocol Buffers** | Service contracts & serialization | v3 |

- Twilio (for Production): `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, `TWILIO_VERIFY_SERVICE_SID`

### Database & Storage

| Technology | Purpose | Version |│   ├── jwt/              # JWT token management

|------------|---------|---------|

| **PostgreSQL** | Primary database | 13 |### 2. Start PostgreSQL

| **UUID Extension** | Primary key generation | - |

│   ├── logger/           # Zap logger│   ├── status_service/    # Kullanıcı durumu (hikaye) servisi

### Authentication & Security

| Technology | Purpose | Version |Start only PostgreSQL with Docker Compose:

|------------|---------|---------|

| **JWT** | Token-based auth | golang-jwt/jwt v5 |│   └── eventbus/         # Kafka client

| **Twilio Verify** | OTP verification | v1.28.4 |

```powershell

### Messaging (Planned)

| Technology | Purpose | Version |docker-compose up -d postgres├── proto/                 # gRPC Protocol Buffer definitions│   ├── message_worker/    # Mesaj kuyruğu işleyici- Docker and Docker Compose

|------------|---------|---------|

| **Apache Kafka** | Event streaming | TBD |```



### Configuration & Utilities├── migrations/            # PostgreSQL schema migrations

| Technology | Purpose | Version |

|------------|---------|---------|PostgreSQL will run on port `5433` (to avoid conflicts with local PostgreSQL).

| **Viper** | Configuration management | v1.21.0 |

| **Zap** | Structured logging | v1.27.0 |├── scripts/              # Helper PowerShell scripts│   └── api_gateway/       # API Gateway (geliştirme aşamasında)- Go 1.18 or higher

| **godotenv** | Environment variables | v1.5.1 |

### 3. Apply Database Migrations

### DevOps

| Technology | Purpose | Version |└── docker-compose.yml    # Docker Compose configuration

|------------|---------|---------|

| **Docker** | Containerization | Latest |**Automated method (recommended):**

| **Docker Compose** | Multi-container orchestration | Latest |

```├── internal/              # Özel uygulama ve iş mantığı

## 🚀 Getting Started

```powershell

### Prerequisites

Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass

- **Docker & Docker Compose** (for PostgreSQL)

- **Go 1.24 or higher**.\scripts\setup-db.ps1

- **PowerShell** (for Windows scripts)

- **Postman** (recommended for gRPC testing)```## 🚀 Quick Start│   ├── auth/             # Auth domain logic### Running the application

- **protoc** (optional, for regenerating proto files)



### 1. Clone the Repository

**Manual method:**

```powershell

git clone https://github.com/et1613/GoApp.git

cd GoApp

``````powershell### Prerequisites│   │   ├── handler/      # gRPC handlers



### 2. Environment Setup# Enable UUID extension



The application loads environment files based on context:Get-Content migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev

- **Local development**: `.env.local`

- **Docker containers**: `.env.docker` (with `RUNNING_IN_DOCKER=true`)



Create your local environment file:# Main schema- **Docker & Docker Compose** (for PostgreSQL)│   │   ├── service/      # İş mantığı katmanı1.  **Start the infrastructure:**



```powershellGet-Content migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

Copy-Item .env.example .env.local

```- **Go 1.24 or higher**



Edit `.env.local` with your configuration:# Device management and revocation



```envGet-Content migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1- **PowerShell** (for running scripts on Windows)│   │   ├── repository/   # Veri erişim arayüzleri    ```bash

# Database Configuration

DATABASE_URL=postgres://user:password@localhost:5433/whatsapp_clone_dev?sslmode=disable



# JWT Configuration# Chat schema- **Postman** (for gRPC testing, optional)

JWT_SECRET=your-super-secret-key-minimum-32-characters-long

Get-Content migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

# Auth Service Configuration

AUTH_SERVICE_GRPC_PORT=50051│   │   ├── store/        # PostgreSQL implementasyonları    docker-compose up -d

AUTH_DEV_MODE=true  # Bypass Twilio (OTP is always "123456")

# Group support

# Twilio Configuration (Production)

# TWILIO_ACCOUNT_SID=your_account_sidGet-Content migrations\0004_add_group_support.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1### 1. Environment Variables Setup

# TWILIO_AUTH_TOKEN=your_auth_token

# TWILIO_VERIFY_SERVICE_SID=your_verify_service_sid```

```

│   │   └── middleware/   # JWT interceptor    ```

**Important Variables:**

- `DATABASE_URL`: PostgreSQL connection string### 4. Run Auth Service

- `JWT_SECRET`: Minimum 32 characters for HMAC-SHA256 signing

- `AUTH_DEV_MODE=true`: Bypass Twilio, use OTP code `123456`The project automatically loads .env files based on context:

- Twilio credentials: Only needed for production OTP verification

From the project root directory:

### 3. Start PostgreSQL

- Local run: `.env.local`│   ├── chat/             # Chat domain logic

Start only PostgreSQL with Docker Compose:

```powershell

```powershell

docker-compose up -d postgresgo run ./cmd/auth_service- Docker run: `.env.docker` (with RUNNING_IN_DOCKER=true)

```

```

PostgreSQL will be available on port `5433` (to avoid conflicts with local installations).

│   ├── realtime/         # WebSocket hub ve handlers2.  **Run database migrations:**

Verify it's running:

Output: `auth_service listening on :50051 (env=local)`

```powershell

docker-compose psCopy the example file to get started:

```

### 5. Test with Postman gRPC

### 4. Apply Database Migrations

│   └── worker/           # Kafka consumer    You'll need a migration tool like `golang-migrate/migrate`.

**Option A: Automated (Recommended)**

1. **Create a new gRPC Request in Postman**

```powershell

Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass2. **Server URL**: `localhost:50051` (plaintext/TLS off)```powershell

.\scripts\setup-db.ps1

```3. **Import proto file**: `proto/auth.proto`



**Option B: Manual**4. **Test AuthService methods**:Copy-Item .env.example .env.local├── pkg/                   # Paylaşılan kütüphaneler    ```bash



```powershell

# Enable UUID extension

Get-Content migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev**SendOTP example:**```



# Apply base schema```json

Get-Content migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

{│   ├── config/           # Yapılandırma yönetimi (Viper)    migrate -database "postgres://postgres:Fbtex1967.@localhost:5432/whatsapp_clone_dev?sslmode=disable" -path migrations up

# Apply device management schema

Get-Content migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1  "phone_number": "+905551234567"



# Apply chat schema}**Important variables:**

Get-Content migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

```

# Apply group support

Get-Content migrations\0004_add_group_support.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1- `DATABASE_URL`: PostgreSQL connection string│   ├── database/         # PostgreSQL bağlantı yönetimi    ```

```

**VerifyOTP example:**

### 5. Run Auth Service

```json- `JWT_SECRET`: Strong secret for JWT signing (minimum 32 bytes)

From the project root:

{

```powershell

go run ./cmd/auth_service  "phone_number": "+905551234567",- `AUTH_DEV_MODE=true`: Twilio bypass (OTP code is always `123456`)│   ├── domain/           # Domain modelleri (User, Message, vb.)

```

  "otp_code": "123456"

Expected output:

```}- Twilio (for Production): `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, `TWILIO_VERIFY_SERVICE_SID`

{"level":"info","ts":...,"msg":"auth_service listening on :50051 (env=local)"}

``````



### 6. Test with Postman (gRPC)│   ├── jwt/              # JWT token yönetimi3.  **Run the services:**



1. **Create a new gRPC Request** in PostmanSuccessful response returns `access_token` and `refresh_token`.

2. **Server URL**: `localhost:50051` (disable TLS)

3. **Import proto file**: Navigate to `proto/auth.proto`### 2. Start PostgreSQL

4. **Select Service**: `auth.AuthService`

**ValidateToken example:**

#### Test Endpoints

```json│   ├── logger/           # Zap logger    Navigate to each service's directory and run it.

**SendOTP**

```json{

{

  "phone_number": "+905551234567"  "access_token": "<ACCESS_TOKEN>"Start only PostgreSQL with Docker Compose:

}

```}



**VerifyOTP** (use `123456` when `AUTH_DEV_MODE=true`)```│   └── eventbus/         # Kafka client    ```bash

```json

{

  "phone_number": "+905551234567",

  "otp_code": "123456"**RefreshToken example:**```powershell

}

``````json

Response includes `access_token` and `refresh_token`.

{docker-compose up -d postgres├── proto/                 # gRPC Protocol Buffer tanımları    go run ./cmd/api_gateway/

**ValidateToken**

```json  "refresh_token": "<REFRESH_TOKEN>"

{

  "access_token": "eyJhbGciOiJIUzI1NiIs..."}```

}

``````



**RefreshToken** (implements rotation)├── migrations/            # PostgreSQL şema migrasyonları    go run ./cmd/auth_service/

```json

{**RevokeCurrentDevice example:**

  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."

}```jsonPostgreSQL will run on port `5433` (to avoid conflicts with local PostgreSQL).

```

{

**RevokeCurrentDevice**

```json  "refresh_token": "<REFRESH_TOKEN>"├── scripts/              # Yardımcı PowerShell scriptleri    # ... and so on for other services

{

  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."}

}

``````### 3. Apply Database Migrations



**LogoutAllDevices**

```json

{**LogoutAllDevices example:**└── docker-compose.yml    # Docker Compose yapılandırması    ```

  "access_token": "eyJhbGciOiJIUzI1NiIs..."

}```json

```

{**Automated method (recommended):**

## 📦 Microservices

  "access_token": "<ACCESS_TOKEN>"

### Auth Service ✅ (Active)

}```

**Port**: 50051 (gRPC)

```

**Responsibilities**:

- Phone number verification with Twilio OTP```powershell

- JWT token generation (access + refresh)

- Token validation and refresh with rotation## 📦 Microservices

- Device-based session management

- Single device and multi-device logoutSet-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass## Quickstart: AuthService + Postman gRPC (Windows)

- User registration and profile management

### Auth Service (Active ✅)

**Technologies**: gRPC, JWT, Twilio Verify, PostgreSQL

- **Port**: 50051 (gRPC).\scripts\setup-db.ps1

**Key Features**:

- **Refresh Token Rotation**: One-time use refresh tokens- **Responsibilities**:

- **Device Sessions**: Track and manage multiple devices

- **Session Revocation**: Logout specific device or all devices  - Phone number verification with Twilio OTP```## 🚀 Hızlı Başlangıç

- **JWT Interceptor**: Portable middleware for other services

  - JWT (access + refresh token) generation

**Domain Logic**:

```  - Token validation and refresh (rotation)

internal/auth/

├── handler/grpc.go           # gRPC endpoint implementations  - Device-based session management

├── service/service.go        # Business logic

├── repository/repository.go  # Data access interfaces  - Single device or all devices logout**Manual method:**End-to-end minimum setup to test OTP and JWT issuance via Postman using gRPC.

├── store/postgres.go         # PostgreSQL implementation

└── middleware/               # JWT validation interceptor- **Technologies**: gRPC, JWT, Twilio, PostgreSQL

```

- **Middleware**: JWT interceptor (portable to other services)

### Chat Service 🚧 (In Development)



**Planned Features**:

- Direct messaging (1-on-1 chats)### Chat Service (In Development 🚧)```powershell### Ön Gereksinimler

- Group messaging

- Message history and pagination- Chat room management

- Read receipts and delivery status

- Typing indicators- User membership management# Enable UUID extension

- Message search

- Group support

**Database Tables**:

- `chat_rooms`: Chat room metadataGet-Content migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev### 1) Environment variables

- `chat_room_members`: User memberships

- `messages`: Message content and metadata### Realtime Service (In Development 🚧)



### Realtime Service 🚧 (In Development)- WebSocket connection management



**Planned Features**:- Real-time message delivery

- WebSocket connection management

- Real-time message delivery- Presence (online status) management# Main schema- **Docker & Docker Compose** (PostgreSQL için)

- Online/offline presence

- Typing indicators broadcast

- Connection pooling and scaling

### Status Service (In Development 🚧)Get-Content migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

**Components**:

- WebSocket Hub: Connection management- User stories

- Message routing

- Presence tracking- Status updates- **Go 1.24 veya üstü**This service auto-loads environment files based on context:



### Status Service 🚧 (In Development)



**Planned Features**:### Message Worker (In Development 🚧)# Device management and revocation

- User stories (24-hour expiry)

- Status updates- Kafka consumer

- View tracking

- Media attachments- Asynchronous message processingGet-Content migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1- **PowerShell** (Windows için script çalıştırma)



### Message Worker 🚧 (In Development)



**Planned Features**:### API Gateway (In Development 🚧)

- Kafka consumer for async processing

- Message persistence- Single entry point for client requests

- Push notification dispatch

- Media processing pipeline- Routing to services via gRPC clients# Chat schema- **Postman** (gRPC test için, opsiyonel)- Local runs: loads `.env.local` (we added one with sensible defaults)



### API Gateway 🚧 (In Development)



**Planned Features**:## 🔐 Authentication ModelGet-Content migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

- HTTP/REST to gRPC translation

- Request routing

- Rate limiting

- API versioningThe system generates two JWT tokens per login:- Docker runs: `docker-compose` passes `.env.docker` and also sets `RUNNING_IN_DOCKER=true` (the app attempts to load `.env.docker` but works even if the file isn't baked into the image)

- Client authentication



## 🔐 Authentication

### Token Types# Group support

### Token Architecture



The system uses a dual-token JWT approach:

| Token Type | Validity | Purpose | Audience |Get-Content migrations\0004_add_group_support.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1### 1. Ortam Değişkenlerini Ayarlama- Base `.env` is also loaded last if present (for overrides)

| Token Type | Lifetime | Purpose | Audience | Storage |

|------------|----------|---------|----------|---------||------------|----------|---------|----------|

| **Access Token** | 15 minutes | API access | `my-app-client` | Memory |

| **Refresh Token** | 7 days | Token renewal | `my-auth-service` | Database (hashed) || **Access Token** | 15 minutes | API access | `my-app-client` |```



### JWT Claims| **Refresh Token** | 7 days | New token generation | `my-auth-service` |



Claims are defined in `pkg/jwt/token.go`:



```go### Token Claims (pkg/jwt/token.go)

type CustomClaims struct {

    Type string `json:"type"` // "access" or "refresh"### 4. Run Auth Service

    jwt.RegisteredClaims

}```go

```

type CustomClaims struct {Proje farklı ortamlar için otomatik .env dosyası yükler:Example files provided:

**Registered Claims**:

- `sub` (Subject): User UUID    Type TokenType `json:"type"` // "access" or "refresh"

- `jti` (JWT ID): Unique token identifier

- `iss` (Issuer): `my-auth-service`    jwt.RegisteredClaimsFrom the project root directory:

- `aud` (Audience): Token type specific

- `iat` (Issued At): Token creation timestamp}

- `exp` (Expires At): Token expiration timestamp

```- Local çalıştırma: `.env.local`

### Refresh Token Rotation



Security feature to prevent token replay attacks:

**Registered Claims:**```powershell

1. Client sends refresh token to `RefreshToken` endpoint

2. Server validates token and checks hash in `user_devices` table- `sub` (Subject): User ID

3. Current device session is revoked (`revoked_at = NOW()`)

4. New access + refresh token pair is generated- `jti` (JWT ID): Unique token IDgo run ./cmd/auth_service- Docker içinde: `.env.docker` (RUNNING_IN_DOCKER=true ile)- `.env.example` — connects to Postgres at `localhost:5432`, sets a sample `JWT_SECRET`, and enables `AUTH_DEV_MODE=true` (Twilio bypass: OTP code is `123456`).

5. New refresh token hash is stored in database

6. **Old refresh token cannot be reused** (one-time use)- `iss` (Issuer): `my-auth-service`



### Session Management- `aud` (Audience): Based on token type```



**Device Tracking**:- `iat` (Issued At): Creation time

- Each login creates a `user_device` record

- Stores SHA-256 hash of refresh token- `exp` (Expires At): Expiration time- `.env.docker.example` — same but using `postgres:5432` for Compose and `RUNNING_IN_DOCKER=true`.

- Tracks device name, type, and push notification token

- Records last login timestamp



**Session Revocation**:### Refresh Token RotationOutput: `auth_service listening on :50051 (env=local)`



**Single Device Logout**:

```

RevokeCurrentDevice(refresh_token) When a refresh token is used:Başlamak için örnek dosyayı kopyalayın:

→ SET revoked_at = NOW() WHERE refresh_token_hash = SHA256(token)

```1. Token is validated and hash is checked in DB



**All Devices Logout**:2. Current session is revoked (`revoked_at` is set)### 5. Test with Postman gRPC

```

LogoutAllDevices(access_token)3. New access + refresh token pair is generated

→ SET revoked_at = NOW() WHERE user_id = token.sub

```4. New refresh token hash is stored in DBCreate a copy for your environment:



### JWT Interceptor Middleware5. **Previous refresh token cannot be reused** (one-time use)



Located in `internal/auth/middleware/auth_interceptor.go`:1. **Create a new gRPC Request in Postman**



**Functionality**:### Session Revocation

- Validates `Authorization: Bearer <token>` header

- Extracts and validates access token2. **Server URL**: `localhost:50051` (plaintext/TLS off)```powershell

- Injects `user_id` into gRPC context

- Exempts specified methods (e.g., auth endpoints)**Single device logout:**



**Usage in Other Services**:```3. **Import proto file**: `proto/auth.proto`



```goRevokeCurrentDevice(refresh_token) → user_devices.revoked_at = NOW()

import (

    "github.com/dykethecreator/GoApp/pkg/jwt"```4. **Test AuthService methods**:Copy-Item .env.example .env.local```powershell

    "github.com/dykethecreator/GoApp/internal/auth/middleware"

)



tokenManager, _ := jwt.NewTokenManager(**All devices logout:**

    jwtSecret, 

    15*time.Minute,  // access token duration```

    7*24*time.Hour,  // refresh token duration

)LogoutAllDevices(access_token) → Revoke all user's devices**SendOTP example:**```Copy-Item .env.example .env



grpcServer := grpc.NewServer(```

    grpc.UnaryInterceptor(

        middleware.UnaryAuthInterceptor(tokenManager),```json

    ),

)### JWT Interceptor Middleware

```

{```

## 🗄 Database Schema

`internal/auth/middleware/auth_interceptor.go`:

### Migrations

- gRPC unary interceptor  "phone_number": "+905551234567"

| File | Description |

|------|-------------|- Validates access token from `Authorization: Bearer <token>` header

| `0000_enable_uuid.sql` | Enable `uuid-ossp` extension |

| `0001_initial_schema.up.sql` | Users, contacts, calls base schema |- Injects `user_id` into context}**Önemli değişkenler:**

| `0002_user_devices_revocation.up.sql` | Device sessions and token revocation |

| `0003_chat_schema.up.sql` | Chat rooms and messages |- AuthService methods are exempted by default (public)

| `0004_add_group_support.up.sql` | Group chat enhancements |

```

### Core Tables

**Usage in other services:**

#### users

```sql- `DATABASE_URL`: PostgreSQL bağlantı dizesiEdit `.env` as needed:

id UUID PRIMARY KEY DEFAULT uuid_generate_v4()

phone_number VARCHAR(20) UNIQUE NOT NULL```go

display_name VARCHAR(255)

profile_picture_url TEXTtokenManager, _ := jwt.NewTokenManager(jwtSecret, 15*time.Minute, 7*24*time.Hour)**VerifyOTP example:**

bio TEXT

created_at TIMESTAMP DEFAULT NOW()server := grpc.NewServer(

updated_at TIMESTAMP DEFAULT NOW()

```    grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor(tokenManager)),```json- `JWT_SECRET`: JWT imzalama için güçlü bir secret (minimum 32 byte)



#### user_devices)

```sql

id UUID PRIMARY KEY DEFAULT uuid_generate_v4()```{

user_id UUID REFERENCES users(id) ON DELETE CASCADE

refresh_token_hash VARCHAR(64) NOT NULL

device_name VARCHAR(255)

device_type VARCHAR(50)## 🗄️ Database Schema  "phone_number": "+905551234567",- `AUTH_DEV_MODE=true`: Twilio bypass (OTP kodu her zaman `123456`)- Set `JWT_SECRET` to a strong value

push_notification_token TEXT

last_login_at TIMESTAMP DEFAULT NOW()

revoked_at TIMESTAMP

created_at TIMESTAMP DEFAULT NOW()### Migrations  "otp_code": "123456"



UNIQUE(user_id, refresh_token_hash)

INDEX idx_user_devices_user_id ON user_devices(user_id)

INDEX idx_user_devices_revocation ON user_devices(user_id, revoked_at)| File | Description |}- Twilio (Production için): `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, `TWILIO_VERIFY_SERVICE_SID`- Add real Twilio credentials if you want to use OTP verification against Twilio (otherwise keep `AUTH_DEV_MODE=true`).

```

|------|-------------|

#### chat_rooms

```sql| `0000_enable_uuid.sql` | Enable UUID extension |```

id UUID PRIMARY KEY DEFAULT uuid_generate_v4()

name VARCHAR(255)| `0001_initial_schema.up.sql` | Base schema (users, contacts, etc.) |

is_group BOOLEAN DEFAULT false

created_by UUID REFERENCES users(id)| `0002_user_devices_revocation.up.sql` | Device management and refresh token revocation |

created_at TIMESTAMP DEFAULT NOW()

updated_at TIMESTAMP DEFAULT NOW()| `0003_chat_schema.up.sql` | Chat rooms and messages |

```

| `0004_add_group_support.up.sql` | Group chat support |Successful response returns `access_token` and `refresh_token`.

#### messages

```sql

id UUID PRIMARY KEY DEFAULT uuid_generate_v4()

chat_room_id UUID REFERENCES chat_rooms(id) ON DELETE CASCADE### Main Tables### 2. PostgreSQL Başlatma### 2) Start PostgreSQL (Docker)

sender_id UUID REFERENCES users(id) ON DELETE SET NULL

content TEXT

message_type VARCHAR(20) DEFAULT 'text'

status VARCHAR(20) DEFAULT 'sent'**users****ValidateToken example:**

created_at TIMESTAMP DEFAULT NOW()

```- `id` (UUID, PK)



## 📚 API Documentation- `phone_number` (UNIQUE)```json



### Auth Service (gRPC)- `display_name`



**Service**: `auth.AuthService`  - `profile_picture_url`{

**Port**: 50051  

**Proto**: `proto/auth.proto`- `created_at`, `updated_at`



#### SendOTP  "access_token": "<ACCESS_TOKEN>"Docker Compose ile sadece PostgreSQL'i başlatın:Run only Postgres in the background:

```protobuf

rpc SendOTP(SendOTPRequest) returns (SendOTPResponse);**user_devices**

```

Sends OTP to phone number via Twilio.- `id` (UUID, PK)}



#### VerifyOTP- `user_id` (FK → users)

```protobuf

rpc VerifyOTP(VerifyOTPRequest) returns (AuthResponse);- `refresh_token_hash` (SHA-256 hash)```

```

Verifies OTP and returns JWT tokens. Creates user if doesn't exist.- `device_name`, `device_type`



#### ValidateToken- `push_notification_token`

```protobuf

rpc ValidateToken(ValidateTokenRequest) returns (ValidateTokenResponse);- `last_login_at`

```

Validates access token and returns user ID.- `revoked_at` (for session revocation)**RefreshToken example:**```powershell```powershell



#### RefreshToken- **UNIQUE constraint**: `(user_id, refresh_token_hash)`

```protobuf

rpc RefreshToken(RefreshTokenRequest) returns (RefreshTokenResponse);```json

```

Rotates refresh token and returns new token pair.**chat_rooms**



#### RevokeCurrentDevice- `id` (UUID, PK){docker-compose up -d postgresdocker-compose up -d postgres

```protobuf

rpc RevokeCurrentDevice(RevokeDeviceRequest) returns (RevokeDeviceResponse);- `name`

```

Revokes specific device session.- `is_group`  "refresh_token": "<REFRESH_TOKEN>"



#### LogoutAllDevices- `created_at`, `updated_at`

```protobuf

rpc LogoutAllDevices(LogoutAllDevicesRequest) returns (LogoutAllDevicesResponse);}``````

```

Revokes all user device sessions.**messages**



## 💻 Development- `id` (UUID, PK)```



### Local vs Docker- `chat_room_id` (FK → chat_rooms)



**Local Development (Recommended)**:- `sender_id` (FK → users)

- App runs on host machine

- Connects to Dockerized PostgreSQL- `content`

- Uses `.env.local`

- Fast iteration and debugging- `message_type` (text, image, video, etc.)**RevokeCurrentDevice example:**

- Command: `go run ./cmd/auth_service`

- `status` (sent, delivered, read)

**Docker Development**:

- Both app and database in containers- `created_at````jsonPostgreSQL `5433` portunda çalışacaktır (yerel PostgreSQL ile çakışmayı önlemek için).### 3) Apply database schema (migrations)

- Uses `.env.docker`

- Production-like environment

- Command: `docker-compose up --build`

## 🛠️ Technology Stack{

### Running Tests



```powershell

# Run all tests### Backend  "refresh_token": "<REFRESH_TOKEN>"

go test ./...

- **Language**: Go 1.24

# Run specific package tests

go test ./internal/auth/service/...- **gRPC**: Inter-service communication}



# Run with verbose output- **Protocol Buffers**: Data serialization

go test -v ./pkg/jwt/...

```### 3. Veritabanı Migrasyonlarını UygulamaOption A: Use the automated script (recommended):

# Run with coverage

go test -cover ./...### Database



# Run specific test- **PostgreSQL 13**: Main data store

go test -v -run TestTokenManager_GenerateAccessToken ./pkg/jwt/...

```- **UUID**: For primary keys



### Protocol Buffer Generation**LogoutAllDevices example:**



When you modify `.proto` files, regenerate Go code:### Authentication



```powershell- **JWT**: golang-jwt/jwt v5```json

# Regenerate auth.proto

protoc --go_out=. --go_opt=paths=source_relative `- **Twilio Verify**: OTP verification

       --go-grpc_out=. --go-grpc_opt=paths=source_relative `

       proto/auth.proto{**Otomatik yöntem (önerilen):**```powershell



# Regenerate all proto files### Messaging (Planned)

Get-ChildItem proto\*.proto | ForEach-Object {

    protoc --go_out=. --go_opt=paths=source_relative `- **Apache Kafka**: Asynchronous message queue  "access_token": "<ACCESS_TOKEN>"

           --go-grpc_out=. --go-grpc_opt=paths=source_relative `

           $_.FullName

}

```### Configuration & Logging}Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass



### Code Organization- **Viper**: Configuration management



**Clean Architecture Layers**:- **Zap**: Structured logging```



1. **Handler** (`internal/*/handler/`): gRPC endpoint implementations- **godotenv**: Environment variables

2. **Service** (`internal/*/service/`): Business logic

3. **Repository** (`internal/*/repository/`): Data access interfaces```powershell.\scripts\setup-db.ps1

4. **Store** (`internal/*/store/`): Database implementations

### Containerization

**Dependency Flow**: `Handler → Service → Repository ← Store`

- **Docker**: Container runtime## 📦 Microservices

### Adding a New Service

- **Docker Compose**: Multi-service orchestration

1. Create service directory structure:

   ```Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass```

   cmd/new_service/main.go

   internal/new_service/## 📝 Development Notes

   ├── handler/

   ├── service/### Auth Service (Active ✅)

   ├── repository/

   └── store/### Local vs Docker Running

   ```

- **Port**: 50051 (gRPC).\scripts\setup-db.ps1

2. Define proto contract in `proto/new_service.proto`

3. Generate Go code: `protoc --go_out=. --go-grpc_out=. proto/new_service.proto`**Local (Recommended - For Quick Testing):**

4. Implement layers following existing patterns

5. Add to `docker-compose.yml` if needed- Application runs locally, connects to Dockerized PostgreSQL- **Responsibilities**:



## 🐛 Troubleshooting- Uses `.env.local`: `DATABASE_URL=...@localhost:5433/...`



### "unknown driver postgres"- Start: `go run ./cmd/auth_service`  - Phone number verification with Twilio OTP```Option B: Manual migration:



**Symptom**: `sql: unknown driver "postgres" (forgotten import?)`



**Solution**: The driver is imported in `pkg/database/postgres.go`. Ensure your service imports this package:**Docker (Production-like):**  - JWT (access + refresh token) generation

```go

import _ "github.com/dykethecreator/GoApp/pkg/database"- Both application and PostgreSQL run in containers

```

- `docker-compose up --build`  - Token validation and refresh (rotation)

### "relation '...' does not exist"

- Uses `.env.docker`: `DATABASE_URL=...@postgres:5432/...`

**Symptom**: Database table not found

  - Device-based session management

**Solution**: Apply migrations in order:

```powershell### Protocol Buffer Code Generation

.\scripts\setup-db.ps1

```  - Single device or all devices logout**Manuel yöntem:**```powershell



Verify database:Regenerate if you modified proto files:

```powershell

docker compose exec postgres psql -U user -d whatsapp_clone_dev -c "\dt"- **Technologies**: gRPC, JWT, Twilio, PostgreSQL

```

```powershell

### "function uuid_generate_v4() does not exist"

# For auth.proto- **Middleware**: JWT interceptor (portable to other services)Get-Content -Raw migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -f -

**Symptom**: UUID function not available

protoc --go_out=. --go_opt=paths=source_relative `

**Solution**: Enable UUID extension:

```sql       --go-grpc_out=. --go-grpc_opt=paths=source_relative `

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

```       proto/auth.proto



Or use `pgcrypto` alternative:### Chat Service (In Development 🚧)```powershellGet-Content -Raw migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -

```sql

CREATE EXTENSION IF NOT EXISTS "pgcrypto";# For all proto files

-- Then use gen_random_uuid() instead

```Get-ChildItem proto\*.proto | ForEach-Object {- Chat room management



### Postman: "Message violates its Protobuf type definition"    protoc --go_out=. --go_opt=paths=source_relative `



**Symptom**: gRPC request fails with type error           --go-grpc_out=. --go-grpc_opt=paths=source_relative `- User membership management# UUID extension'ı etkinleştirGet-Content -Raw migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -



**Solution**:           $_.FullName

- Ensure field names match proto exactly (`access_token`, not `accessToken`)

- Re-import `proto/auth.proto` in Postman}- Group support

- Check for missing required fields

- Verify data types (string vs int)```



### JWT "Unauthenticated" errorGet-Content migrations\0000_enable_uuid.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_devGet-Content -Raw migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1 -f -



**Symptom**: Token validation fails### Running Tests



**Solutions**:### Realtime Service (In Development 🚧)

- Verify `JWT_SECRET` is identical in all environments

- Check token hasn't expired (access: 15min, refresh: 7 days)```powershell

- Don't use access token at refresh endpoints

- Ensure token type matches endpoint (`access` vs `refresh`)# All tests- WebSocket connection management```

- Check `Authorization: Bearer <token>` header format

go test ./...

### Docker Postgres Connection Failed

- Real-time message delivery

**Symptom**: Cannot connect to PostgreSQL

# Specific package

**Solutions**:

go test ./internal/auth/service/...- Presence (online status) management# Ana şema

Check container status:

```powershell

docker-compose ps

```# Verbose output



View logs:go test -v ./pkg/jwt/...

```powershell

docker-compose logs postgres```### Status Service (In Development 🚧)Get-Content migrations\0001_initial_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1If you previously ran into an inline index syntax error for `call_logs`, this repository already fixes it by creating the index separately.

```



Verify health:

```powershell## 🐛 Troubleshooting- User stories

docker inspect goapp-postgres-1 | Select-String -Pattern "Health"

```



Test connection:### "unknown driver postgres"- Status updates

```powershell

docker compose exec postgres psql -U user -d whatsapp_clone_dev -c "SELECT 1;"✅ **Solution**: `github.com/lib/pq` is already blank imported in `pkg/database/postgres.go`.

```



### Port Already in Use

### "relation '...' does not exist"

**Symptom**: `bind: address already in use`

✅ **Solution**: Ensure migrations are applied to the correct database:### Message Worker (In Development 🚧)# Cihaz yönetimi ve revocation### 4) Run Auth Service locally

**Solutions**:

```powershell

Check what's using the port:

```powershell.\scripts\setup-db.ps1- Kafka consumer

# Check port 5433 (PostgreSQL)

netstat -ano | findstr :5433```



# Check port 50051 (gRPC)- Asynchronous message processingGet-Content migrations\0002_user_devices_revocation.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1

netstat -ano | findstr :50051

```### "function uuid_generate_v4() does not exist"



Kill process or change port in configuration.✅ **Solution**: Enable UUID extension:



## 🗺 Roadmap```sql



### Phase 1: Core Authentication ✅ (Complete)CREATE EXTENSION IF NOT EXISTS "uuid-ossp";### API Gateway (In Development 🚧)From the project root:

- [x] JWT-based authentication

- [x] Twilio OTP verification```

- [x] Refresh token rotation

- [x] Device session managementOr use `pgcrypto` with `gen_random_uuid()`.- Single entry point for client requests

- [x] Session revocation (single + all devices)

- [x] JWT interceptor middleware



### Phase 2: Messaging Foundation 🚧 (Current)### Postman: "Message violates its Protobuf type definition"- Routing to services via gRPC clients# Chat şeması

- [ ] Chat Service gRPC API

- [ ] Direct messaging (1-on-1)✅ **Solution**: 

- [ ] Group messaging support

- [ ] Message persistence- Ensure field names match the proto file (`access_token` not `accessToken`)

- [ ] Read receipts

- [ ] API Gateway implementation- Re-import `proto/auth.proto` file



### Phase 3: Real-time Features 🔜 (Next)## 🔐 Authentication ModelGet-Content migrations\0003_chat_schema.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1```powershell

- [ ] WebSocket connection management

- [ ] Real-time message delivery### JWT "Unauthenticated" error

- [ ] Online/offline presence

- [ ] Typing indicators✅ **Solution**:

- [ ] Message delivery status updates

- Ensure `JWT_SECRET` is the same in both environments

### Phase 4: Advanced Features 📅 (Planned)

- [ ] Status/Stories service- Don't use access token at refresh endpoint (token type check exists)The system generates two JWT tokens per login:go run ./cmd/auth_service

- [ ] Voice/video calling (signaling)

- [ ] Media upload and storage (S3/MinIO)- Ensure token hasn't expired

- [ ] End-to-end encryption

- [ ] Push notifications (FCM/APNS)

- [ ] Message search and indexing

### Docker Postgres connection error

### Phase 5: Scalability & Operations 🎯 (Future)

- [ ] Kafka event streaming integration✅ **Solution**:### Token Types# Grup desteği```

- [ ] Message worker implementation

- [ ] Metrics and monitoring (Prometheus)```powershell

- [ ] Distributed tracing (Jaeger)

- [ ] Rate limiting and DDoS protection# Check container status

- [ ] Database sharding strategy

- [ ] Redis caching layerdocker-compose ps

- [ ] CI/CD pipeline

- [ ] Kubernetes deployment| Token Type | Validity | Purpose | Audience |Get-Content migrations\0004_add_group_support.up.sql | docker compose exec -T postgres psql -U user -d whatsapp_clone_dev -v ON_ERROR_STOP=1



### Phase 6: Quality & Security 🔒 (Ongoing)# View logs

- [ ] Comprehensive unit tests (80%+ coverage)

- [ ] Integration testsdocker-compose logs postgres|------------|----------|---------|----------|

- [ ] Load testing and benchmarks

- [ ] Security audit

- [ ] API documentation (Swagger/OpenAPI)

- [ ] Developer documentation# Check healthcheck status| **Access Token** | 15 minutes | API access | `my-app-client` |```You should see a log similar to: `auth_service listening on :50051 (env=local)`.

- [ ] Deployment guides

docker inspect goapp-postgres-1 | Select-String -Pattern "Health"

## 📖 Additional Documentation

```| **Refresh Token** | 7 days | New token generation | `my-auth-service` |

For more detailed information, see:

- [Backend Documentation](docs/Backend.md) - Deep dive into architecture

- [Proto Definitions](proto/) - gRPC service contracts

- [Migration Files](migrations/) - Database schema evolution## 📚 More Information



## 📄 License



This project is for **educational purposes only**. For more detailed backend documentation: [`docs/Backend.md`](docs/Backend.md)### Token Claims (pkg/jwt/token.go)



## 🤝 Contributing



Contributions are welcome! Please follow these steps:## 📄 License### 4. Auth Service'i Çalıştırma### 5) Test via Postman (gRPC)



1. Fork the repository

2. Create a feature branch (`git checkout -b feature/amazing-feature`)

3. Commit your changes (`git commit -m 'Add amazing feature'`)This project is for educational purposes.```go

4. Push to the branch (`git push origin feature/amazing-feature`)

5. Open a Pull Request



### Contribution Guidelines## 👨‍💻 Contributingtype CustomClaims struct {



- Follow Go best practices and conventions

- Write tests for new features

- Update documentation as needed1. Fork the repository    Type string `json:"type"` // "access" or "refresh"

- Use meaningful commit messages

- Keep PRs focused and atomic2. Create a feature branch (`git checkout -b feature/amazing-feature`)



## 👥 Authors3. Commit your changes (`git commit -m 'Add amazing feature'`)    jwt.RegisteredClaimsProje kök dizininden:1. In Postman, create a new gRPC Request.



- **Project Lead** - Initial work and architecture4. Push to the branch (`git push origin feature/amazing-feature`)



## 🙏 Acknowledgments5. Create a Pull Request}



- Go gRPC team for excellent tooling

- Twilio for OTP verification services

- PostgreSQL community for robust database---```2. Server URL: `localhost:50051` (plaintext/TLS off).

- Clean Architecture principles by Robert C. Martin



## 📞 Support

**Note**: This project is under active development. Some features are not yet complete.

For questions or issues:

- Open an issue on GitHub

- Check existing documentation**Registered Claims:**```powershell3. Import `proto/auth.proto`.

- Review troubleshooting section

- `sub` (Subject): User ID

---

- `jti` (JWT ID): Unique token IDgo run ./cmd/auth_service4. Select `auth.AuthService` and call:

**Note**: This project is under active development. Some features are still being implemented. Check the roadmap for current status and upcoming features.

- `iss` (Issuer): `my-auth-service`

**Last Updated**: December 2025

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

##  More Information

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
