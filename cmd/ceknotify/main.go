package main

import (
	"fmt"
	"log"

	"ceknotify/internal/ntfy"
)

func main() {
	fmt.Println("Hello, World!")

	var serverURL, topic string
	fmt.Print("Enter ntfy server URL (default: https://ntfy.sh): ")
	fmt.Scanln(&serverURL)
	if serverURL == "" {
		serverURL = "https://ntfy.sh"
	}
	fmt.Print("Enter topic to subscribe to (default: update-info-bar): ")
	fmt.Scanln(&topic)
	if topic == "" {
		topic = "update-info-bar"
	}

	client := ntfy.NewNtfyClient(serverURL, topic)

	// Prompt user to choose connection method
	fmt.Println("\nChoose connection method:")
	fmt.Println("1. JSON streaming")
	fmt.Println("2. Server-Sent Events (SSE)")
	fmt.Print("Enter your choice (1 or 2): ")

	var choice int
	fmt.Scanln(&choice)

	fmt.Printf("Connecting to topic '%s' on server '%s'...\n", topic, serverURL)
	switch choice {
	case 1:
		fmt.Println("📊 Using JSON streaming...")
		if err := client.Listen(); err != nil {
			log.Fatalf("❌ Error listening for notifications: %v", err)
		}
	case 2:
		fmt.Println("📡 Using Server-Sent Events (SSE)...")
		if err := client.ListenSSE(); err != nil {
			log.Fatalf("❌ Error listening for notifications: %v", err)
		}
	default:
		fmt.Println("⚠️  Invalid choice, defaulting to JSON streaming...")
		if err := client.Listen(); err != nil {
			log.Fatalf("❌ Error listening for notifications: %v", err)
		}
	}
	fmt.Println("Exiting...")
	fmt.Println("Goodbye!")
}
