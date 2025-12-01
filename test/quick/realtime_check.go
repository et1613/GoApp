package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/dykethecreator/GoApp/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// Hard-coded test values - only change these!
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoiYWNjZXNzIiwiaXNzIjoibXktYXV0aC1zZXJ2aWNlIiwic3ViIjoiZDk0YmIyYWUtMTg1Mi00OGY5LWFkZmMtMmMyZDgxZjc4MTk5IiwiYXVkIjpbIm15LWFwcC1jbGllbnQiXSwiZXhwIjoxNzY0NTEwOTQzLCJpYXQiOjE3NjQ1MTAwNDMsImp0aSI6ImJmYWE0ODg1LTRlNGItNGM2OS05ODYwLTczNjA0YjM4YWVkNyJ9.yziT0zxRzdMFI3Nrtuvf1LXM_2uakYkDqGJDRvyRq1s"
	conversationID := "df396d61-6059-4545-b317-acd1f51a99cd"

	fmt.Println("🧪 Quick Realtime Test")
	fmt.Println("═══════════════════════════════════════════")

	// Connect
	conn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Connection failed: %v", err)
	}
	defer conn.Close()

	client := pb.NewRealtimeServiceClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+accessToken)

	stream, err := client.Connect(ctx)
	if err != nil {
		log.Fatalf("❌ Stream failed: %v", err)
	}

	fmt.Printf("✅ Connected to Realtime Service\n")
	fmt.Printf("   Monitoring conversation: %s\n", conversationID)

	// Listen for events
	go func() {
		for {
			event, err := stream.Recv()
			if err == io.EOF {
				log.Println("🔌 Stream closed")
				return
			}
			if err != nil {
				log.Printf("❌ Receive error: %v", err)
				return
			}

			switch e := event.Event.(type) {
			case *pb.ServerEvent_Pong:
				fmt.Printf("  ✅ PONG received\n")
			case *pb.ServerEvent_NewMessage:
				fmt.Printf("  📨 NEW MESSAGE: %s\n", e.NewMessage.Content)
			case *pb.ServerEvent_Presence:
				fmt.Printf("  👤 User %s is %s\n", e.Presence.UserId, e.Presence.Status)
			case *pb.ServerEvent_Typing:
				fmt.Printf("  ⌨️  Typing: %v\n", e.Typing.IsTyping)
			}
		}
	}()

	// Send ping
	fmt.Println("\n📤 Sending PING...")
	err = stream.Send(&pb.ClientEvent{
		Event: &pb.ClientEvent_Ping{Ping: &pb.Ping{Timestamp: time.Now().Unix()}},
	})
	if err != nil {
		log.Fatalf("❌ Ping failed: %v", err)
	}

	// Wait for messages
	fmt.Println("⏳ Listening for 60 seconds...")
	fmt.Println("   (Now send SendMessage from Postman!)")
	fmt.Println("═══════════════════════════════════════════")
	fmt.Println()

	time.Sleep(60 * time.Second)
	fmt.Println("\n✅ Test completed!")
}
