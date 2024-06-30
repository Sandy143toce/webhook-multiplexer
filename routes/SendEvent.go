package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"context"

	Database "github.com/Sandy143toce/webhook-multiplexer/database"
	"github.com/Sandy143toce/webhook-multiplexer/models"
	"github.com/Sandy143toce/webhook-multiplexer/utils"
	"github.com/gofiber/fiber/v2"
)

func SendEvent(c *fiber.Ctx) error {
	var body models.SendEvent
	_ = c.BodyParser(&body)
	fmt.Println("C is", c)
	protocol := "http"
	if c.Protocol() == "https" {
		protocol = "https"
	}
	host := c.Hostname()
	path := c.Path()
	fullURL := fmt.Sprintf("%s://%s%s", protocol, host, path)
	fmt.Println("Full URL is", fullURL) // This will print the full URL and can be used to fetch corresponding mapped endpoints

	WebhookData, err := FetchWebhookAndEndpoints(fullURL, "active")
	if err != nil {
		fmt.Println("Error while fetching webhook data", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error while fetching webhook data",
		})
	}
	fmt.Println("Webhook Data:", WebhookData)
	if len(WebhookData) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Message": "Event sent successfully",
		})
	}
	results := ProcessEndpointsConcurrently(WebhookData, body)
	fmt.Println("Results:", results)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Event sent successfully",
	})
}

func FetchWebhookAndEndpoints(webhook_url string, status string) ([]models.WebhookAndEndpoints, error) {
	sqlStatement := `
	SELECT w.id, 
	ce.url,
	ce.id AS endpoint_id
	FROM webhooks w
	INNER JOIN endpoints ce ON w.id = ce.webhook_id
	WHERE w.url=$1 AND w.status=$2 AND ce.status=$2`

	rows, err := Database.DBConn.Query(context.Background(), sqlStatement, webhook_url, status)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch webhook and endpoints: %v", err)
	}
	defer rows.Close()

	var webhookAndEndpoints []models.WebhookAndEndpoints

	for rows.Next() {
		var webhookAndEndpoint models.WebhookAndEndpoints
		err := rows.Scan(&webhookAndEndpoint.WebhookID, &webhookAndEndpoint.URL, &webhookAndEndpoint.EndPointId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan webhook and endpoints: %v", err)
		}
		webhookAndEndpoints = append(webhookAndEndpoints, webhookAndEndpoint)
	}
	return webhookAndEndpoints, nil
}

func ProcessEndpointsConcurrently(endpoints []models.WebhookAndEndpoints, payload models.SendEvent) map[string]string {
	var wg sync.WaitGroup
	results := make(map[string]string)
	resultsMutex := &sync.Mutex{}

	for _, endpoint := range endpoints {
		wg.Add(1)
		go func(e models.WebhookAndEndpoints) {
			defer wg.Done()
			result := forwardWebhook(e.URL, payload)

			resultsMutex.Lock()
			results[e.EndPointId] = result
			resultsMutex.Unlock()

			// Store the result for this endpoint
			err := storeWebhookResult(e, result, payload)
			if err != nil {
				// Log the error, but continue processing other endpoints
				fmt.Printf("Error storing result for endpoint %s: %v\n", e.EndPointId, err)
			}
		}(endpoint)
	}

	wg.Wait()
	return results
}

func forwardWebhook(url string, payload models.SendEvent) string {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Sprintf("Error marshalling payload: %v", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Sprintf("Error sending webhook: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response: %v", err)
	}

	return fmt.Sprintf("Status: %d, Body: %s", resp.StatusCode, string(body))
}

func storeWebhookResult(endpoint models.WebhookAndEndpoints, result string, payload models.SendEvent) error {
	payloadstr, err := StringifyPayload(payload)
	if err != nil {
		fmt.Errorf("failed to stringify payload: %v", err)
		payloadstr = "NA"
	}
	ResponseId := utils.UUIDGenerator("RP-")
	sqlStatement := `INSERT INTO response (id, webhook_id, endpoint_id, body, result, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = Database.DBConn.Exec(context.Background(), sqlStatement, ResponseId, endpoint.WebhookID, endpoint.EndPointId, payloadstr, result, time.Now())

	if err != nil {
		return fmt.Errorf("failed to store webhook result: %v", err)
	}

	return nil
}

func StringifyPayload(payload models.SendEvent) (string, error) {
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}
	return string(jsonBytes), nil
}
