package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/dykethecreator/GoApp/internal/auth/middleware"
	"github.com/dykethecreator/GoApp/internal/realtime"
	realtimeHandler "github.com/dykethecreator/GoApp/internal/realtime/handler"
	"github.com/dykethecreator/GoApp/pkg/jwt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	envFile := ".env." + env
	if env == "local" {
		envFile = ".env.local"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: .env file not found (%s), using system environment", envFile)
	}

	// Get port
	port := os.Getenv("REALTIME_SERVICE_GRPC_PORT")
	if port == "" {
		port = "50053"
	}

	// Setup JWT for auth interceptor
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}

	tokenManager, err := jwt.NewTokenManager(jwtSecret, 15*time.Minute, 7*24*time.Hour)
	if err != nil {
		log.Fatalf("Failed to create token manager: %v", err)
	}

	// Use global hub (shared with other services in-process)
	hub := realtime.GetGlobalHub()

	// Start hub event loop in background
	go hub.Run()

	// Setup gRPC server with auth interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor(tokenManager)),
		grpc.StreamInterceptor(middleware.StreamAuthInterceptor(tokenManager)),
	)

	// Register realtime handler
	handler := realtimeHandler.NewRealtimeHandler(hub)
	handler.Register(grpcServer)

	// Enable reflection for grpcurl/Postman
	reflection.Register(grpcServer)

	// Start listening
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("realtime_service listening on :%s (env=%s)", port, env)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
