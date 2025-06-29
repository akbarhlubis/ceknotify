package ntfy

import (
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
)

// displayNotification shows notification in terminal
func (c *NtfyClient) displayNotification(msg NtfyMessage) {
	fmt.Printf("\n📨 New notification received:\n")
	fmt.Printf("   Title: %s\n", msg.Title)
	fmt.Printf("   Message: %s\n", msg.Message)
	fmt.Printf("   Time: %s\n", time.Unix(msg.Time, 0).Format("15:04:05"))
	if len(msg.Tags) > 0 {
		fmt.Printf("   Tags: %v\n", msg.Tags)
	}
	fmt.Println("---")

	// Desktop notification
	if err := c.showDesktopNotification(msg); err != nil {
		log.Printf("⚠️ Failed to show desktop notification: %v", err)
	}
}

func (c *NtfyClient) showDesktopNotification(msg NtfyMessage) error {
	title := msg.Title
	if title == "" {
		title = "Notification from " + c.Topic
	}

	message := msg.Message

	// Send the notification
	return beeep.Alert(title, message, "")
}
