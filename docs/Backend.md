# Backend Documentation

This document explains the backend architecture, services, technology stack, environment setup, authentication model, and operational notes for this repository.

## Overview

- Language/runtime: Go
- Communication: gRPC with Protocol Buffers (see `proto/`)
- Database: PostgreSQL
- Messaging: Kafka (planned wiring under `pkg/eventbus` and `internal/worker`)
- Auth: JWT (access + refresh) with device-scoped refresh token revocation
- Container: Docker & Docker Compose

Key folders:
- `cmd/`: service entrypoints
- `internal/`: private app logic per domain/service
- `pkg/`: shared libraries (config, DB, JWT, logging, domain models)
- `proto/`: protobuf definitions and generated Go code
- `migrations/`: SQL migrations

## Services (current state)

- Auth Service (`cmd/auth_service`): implements OTP verification (via Twilio), JWT issuance/validation, refresh token rotation, session revocation, and gRPC methods.
- API Gateway, Chat, Status, Realtime, Message Worker: stubs/placeholders in this repo and will be expanded progressively.

## gRPC & Protobuf

- Service definitions live in `proto/*.proto`.
- Generated code lives alongside (`proto/auth.pb.go`, `proto/auth_grpc.pb.go`).
- Regeneration (optional, for contributors):
  - You need `protoc` and the Go plugins installed.
  - Windows PowerShell example (adjust paths if needed):
    - `protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/auth.proto`

## Configuration and environments

The auth service loads environment files based on context:

- Local default: `.env.local`
- In containers (Docker): `.env.docker` (and `RUNNING_IN_DOCKER=true` is set by Compose)
- Base `.env` may be loaded last for overrides

Important environment variables:
- `DATABASE_URL`: PostgreSQL connection string
- `JWT_SECRET`: HMAC secret for JWT signing (minimum 32 bytes)
- `AUTH_SERVICE_GRPC_PORT`: default `50051`
- Twilio: `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, `TWILIO_VERIFY_SERVICE_SID`

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
