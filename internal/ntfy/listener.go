package ntfy

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Listen using JSON streaming
func (c *NtfyClient) Listen() error {
	url := fmt.Sprintf("%s/%s/json", c.ServerURL, c.Topic)

	fmt.Printf("🔗 Connecting to JSON stream: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to connect to ntfy: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ntfy server returned status: %d", resp.StatusCode)
	}

	fmt.Printf("✅ Connected to topic '%s' via JSON\n", c.Topic)
	fmt.Println("🔔 Waiting for notifications...")

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var msg NtfyMessage
		if err := json.Unmarshal([]byte(line), &msg); err != nil {
			log.Printf("Failed to parse message: %v", err)
			continue
		}

		// Only process actual messages (skip keepalive)
		if msg.Event == "message" {
			c.displayNotification(msg)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading from ntfy stream: %w", err)
	}

	return nil
}

func (c *NtfyClient) ListenSSE() error {
	url := fmt.Sprintf("%s/%s/sse", c.ServerURL, c.Topic)

	fmt.Printf("🔗 Connecting to SSE stream: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers for SSE
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to ntfy SSE: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ntfy server returned status: %d", resp.StatusCode)
	}

	fmt.Printf("✅ Connected to topic '%s' via SSE\n", c.Topic)
	fmt.Println("🔔 Waiting for notifications...")

	scanner := bufio.NewScanner(resp.Body)
	var eventData string

	for scanner.Scan() {
		line := scanner.Text()

		// Handle SSE format
		if strings.HasPrefix(line, "data: ") {
			eventData = strings.TrimPrefix(line, "data: ")
		} else if line == "" && eventData != "" {
			// Empty line indicates end of event, process the data
			c.processSSEMessage(eventData)
			eventData = ""
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading from ntfy SSE stream: %w", err)
	}

	return nil
}

func (c *NtfyClient) processSSEMessage(data string) {
	if data == "" {
		return
	}

	var msg NtfyMessage
	if err := json.Unmarshal([]byte(data), &msg); err != nil {
		log.Printf("Failed to parse SSE message: %v", err)
		return
	}

	// Only process actual messages (skip keepalive)
	if msg.Event == "message" {
		c.displayNotification(msg)
	}
}
