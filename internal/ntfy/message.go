package ntfy

type NtfyMessage struct {
	ID       string   `json:"id"`
	Time     int64    `json:"time"`
	Event    string   `json:"event"`
	Topic    string   `json:"topic"`
	Message  string   `json:"message"`
	Title    string   `json:"title"`
	Tags     []string `json:"tags"`
	Priority int      `json:"priority"`
}
