package ntfy

type NtfyClient struct {
	ServerURL string
	Topic     string
}

func NewNtfyClient(serverURL, topic string) *NtfyClient {
	return &NtfyClient{
		ServerURL: serverURL,
		Topic:     topic,
	}
}
