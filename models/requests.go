package models

type Endpoint struct {
	WebhookID string `json:"webhook_id"`
	URL       string `json:"url"`
}

type Webhook struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type SendEvent struct {
	EventName string      `json:"event_name"`
	Metadata  interface{} `json:"metadata"`
}

// type PayLoad struct {
// 	EventName string      `json:"event_name"`
// 	Metadata  interface{} `json:"metadata"`
// }
