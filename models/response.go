package models

import "time"

type Response struct {
	ID         string `json:"id"`
	EndpointID string `json:"endpoint_id"`
	WebhookID  string `json:"webhook_id"`
	Body       string `json:"body"`
	Result     string `json:"result"`
	CreatedAt  string `json:"created_at"`
}

type EndpointResponse struct {
	ID        string    `json:"id"`
	WebhookID string    `json:"webhook_id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type WebhookResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type WebhookAndEndpoints struct {
	WebhookID  string `json:"webhook_id"`
	URL        string `json:"url"`
	EndPointId string `json:"endpoint_id"`
}
