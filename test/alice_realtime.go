package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/dykethecreator/GoApp/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run realtime_test.go <USER_NAME> <ACCESS_TOKEN> <CONVERSATION_ID>")
		fmt.Println("Example: go run realtime_test.go Alice eyJhbGc... conv-123-456")
		os.Exit(1)
	}

	userName := os.Args[1]
	accessToken := os.Args[2]
	conversationID := os.Args[3]

	fmt.Printf("🚀 Starting Realtime Client for %s\n", userName)
	fmt.Printf("📡 Connecting to localhost:50050...\n")

	// gRPC connection
	conn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRealtimeServiceClient(conn)

	// Send token via metadata
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken)

	// Open bidirectional stream
	stream, err := client.Connect(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to connect stream: %v", err)
	}

	fmt.Printf("✅ Connected to Realtime Service as %s\n", userName)
	fmt.Println("─────────────────────────────────────────")

	// Listen for incoming messages (goroutine)
	go func() {
		for {
			event, err := stream.Recv()
			if err == io.EOF {
				log.Println("🔌 Stream closed by server")
				return
			}
			if err != nil {
				log.Printf("❌ Error receiving: %v", err)
				return
			}

			// Process based on event type
			timestamp := time.Now().Format("15:04:05")
			switch e := event.Event.(type) {
			case *pb.ServerEvent_Pong:
				fmt.Printf("[%s] ✅ PONG received\n", timestamp)
			case *pb.ServerEvent_NewMessage:
				fmt.Printf("[%s] 📨 NEW MESSAGE from %s:\n", timestamp, e.NewMessage.SenderId[:8])
				fmt.Printf("     Content: %s\n", e.NewMessage.Content)
				fmt.Printf("     Conversation: %s\n", e.NewMessage.ConversationId[:8])
			case *pb.ServerEvent_Typing:
				status := "typing..."
				if !e.Typing.IsTyping {
					status = "stopped typing"
				}
				fmt.Printf("[%s] ⌨️  %s is %s\n", timestamp, e.Typing.UserId[:8], status)
			case *pb.ServerEvent_Presence:
				fmt.Printf("[%s] 👤 %s is now %s\n", timestamp, e.Presence.UserId[:8], e.Presence.Status)
			case *pb.ServerEvent_Delivered:
				fmt.Printf("[%s] ✓✓ Message %s delivered\n", timestamp, e.Delivered.MessageId[:8])
			}
		}
	}()

	// Send initial ping
	err = stream.Send(&pb.ClientEvent{
		Event: &pb.ClientEvent_Ping{Ping: &pb.Ping{Timestamp: time.Now().Unix()}},
	})
	if err != nil {
		log.Fatalf("❌ Failed to send ping: %v", err)
	}

	// Heartbeat (ping every 30 seconds)
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			err := stream.Send(&pb.ClientEvent{
				Event: &pb.ClientEvent_Ping{Ping: &pb.Ping{Timestamp: time.Now().Unix()}},
			})
			if err != nil {
				log.Printf("❌ Error sending ping: %v", err)
				return
			}
		}
	}()

	// Listen for commands (for typing indicator)
	fmt.Println("\n💡 Commands:")
	fmt.Println("   't' + ENTER = Send typing indicator")
	fmt.Println("   's' + ENTER = Stop typing indicator")
	fmt.Println("   'q' + ENTER = Quit")
	fmt.Println("─────────────────────────────────────────\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		switch command {
		case "t":
			err := stream.Send(&pb.ClientEvent{
				Event: &pb.ClientEvent_Typing{
					Typing: &pb.TypingIndicator{
						ConversationId: conversationID,
						UserId:         userName,
						IsTyping:       true,
					},
				},
			})
			if err != nil {
				log.Printf("❌ Error sending typing: %v", err)
			} else {
				fmt.Printf("⌨️  You are typing...\n")
			}
		case "s":
			err := stream.Send(&pb.ClientEvent{
				Event: &pb.ClientEvent_Typing{
					Typing: &pb.TypingIndicator{
						ConversationId: conversationID,
						UserId:         userName,
						IsTyping:       false,
					},
				},
			})
			if err != nil {
				log.Printf("❌ Error sending stop typing: %v", err)
			} else {
				fmt.Printf("⌨️  You stopped typing\n")
			}
		case "q":
			fmt.Println("👋 Disconnecting...")
			return
		default:
			fmt.Printf("❓ Unknown command: %s\n", command)
		}
	}
}
