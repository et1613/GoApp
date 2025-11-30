package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/dykethecreator/GoApp/internal/auth/middleware"
	"github.com/dykethecreator/GoApp/internal/chat/handler"
	"github.com/dykethecreator/GoApp/internal/chat/service"
	"github.com/dykethecreator/GoApp/internal/chat/store"
	"github.com/dykethecreator/GoApp/internal/realtime"
	"github.com/dykethecreator/GoApp/pkg/database"
	appjwt "github.com/dykethecreator/GoApp/pkg/jwt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Environment loading strategy (same as auth_service)
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		if os.Getenv("RUNNING_IN_DOCKER") != "" {
			appEnv = "docker"
		} else {
			appEnv = "local"
		}
	}

	_ = godotenv.Load(".env." + appEnv)
	_ = godotenv.Load()

	// Start global hub for realtime messaging (in-process)
	hub := realtime.GetGlobalHub()
	go hub.Run()
	log.Println("Realtime hub started in chat_service")

	// Database connection
	db, err := database.NewDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// gRPC port (default 50052)
	grpcPort := os.Getenv("CHAT_SERVICE_GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50052"
	}
	listenAddr := fmt.Sprintf(":%s", grpcPort)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Build TokenManager for interceptor (same secret as auth_service)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}
	tm, err := appjwt.NewTokenManager(jwtSecret, 15*time.Minute, 7*24*time.Hour)
	if err != nil {
		log.Fatalf("failed to init token manager: %v", err)
	}

	// Create gRPC server with auth interceptor
	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor(tm)),
	)

	// DI: ChatStore → ChatService → ChatHandler
	chatStore := store.NewChatStore(db.DB)
	chatService := service.NewChatService(chatStore)
	chatHandler := handler.NewChatHandler(chatService)

	// Register handler
	chatHandler.Register(s)

	log.Printf("chat_service listening on %s (env=%s)", listenAddr, os.Getenv("APP_ENV"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
