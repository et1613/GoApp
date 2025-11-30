package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/dykethecreator/GoApp/internal/auth/handler"
	authMiddleware "github.com/dykethecreator/GoApp/internal/auth/middleware"
	authService "github.com/dykethecreator/GoApp/internal/auth/service"
	authStore "github.com/dykethecreator/GoApp/internal/auth/store"
	chatHandler "github.com/dykethecreator/GoApp/internal/chat/handler"
	chatService "github.com/dykethecreator/GoApp/internal/chat/service"
	chatStore "github.com/dykethecreator/GoApp/internal/chat/store"
	"github.com/dykethecreator/GoApp/internal/realtime"
	realtimeHandler "github.com/dykethecreator/GoApp/internal/realtime/handler"
	"github.com/dykethecreator/GoApp/pkg/database"
	appjwt "github.com/dykethecreator/GoApp/pkg/jwt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("🚀 Starting All-In-One Service (Auth + Chat + Realtime)")

	// Load env
	_ = godotenv.Load(".env.local")

	// Database
	db, err := database.NewDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// JWT Token Manager
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}
	tokenManager, err := appjwt.NewTokenManager(jwtSecret, 15*time.Minute, 7*24*time.Hour)
	if err != nil {
		log.Fatalf("Failed to create token manager: %v", err)
	}

	// Realtime Hub (shared in-process)
	hub := realtime.GetGlobalHub()
	go hub.Run()
	log.Println("✅ Realtime Hub started")

	// Auth Components
	userRepo := authStore.NewUserStore(db.DB)
	deviceRepo := authStore.NewUserDeviceStore(db.DB)
	authSvc := authService.NewAuthService(userRepo, deviceRepo)
	authHandler := handler.NewAuthHandler(authSvc)

	// Chat Components
	chatRepo := chatStore.NewChatStore(db.DB)
	chatSvc := chatService.NewChatService(chatRepo)
	chatHdlr := chatHandler.NewChatHandler(chatSvc)

	// Realtime Handler
	realtimeHdlr := realtimeHandler.NewRealtimeHandler(hub)

	// gRPC Server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authMiddleware.UnaryAuthInterceptor(tokenManager)),
		grpc.StreamInterceptor(authMiddleware.StreamAuthInterceptor(tokenManager)),
	)

	// Register all services
	authHandler.Register(grpcServer)
	chatHdlr.Register(grpcServer)
	realtimeHdlr.Register(grpcServer)

	reflection.Register(grpcServer)

	// Listen on single port
	port := "50050"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf(`
╔═══════════════════════════════════════╗
║   ALL-IN-ONE SERVICE READY            ║
╠═══════════════════════════════════════╣
║ Port: %s                              ║
║ Services:                             ║
║   ✅ AuthService                       ║
║   ✅ ChatService                       ║
║   ✅ RealtimeService                   ║
║ Realtime Hub: SHARED ✓                ║
╚═══════════════════════════════════════╝
`, port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
