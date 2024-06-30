package routes

import (
	"fmt"

	"context"
	"time"

	Database "github.com/Sandy143toce/webhook-multiplexer/database"
	"github.com/Sandy143toce/webhook-multiplexer/models"
	"github.com/Sandy143toce/webhook-multiplexer/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateWebhook(c *fiber.Ctx) error {
	var body models.Webhook
	_ = c.BodyParser(&body)
	WebhookId := utils.UUIDGenerator("WH-")
	fmt.Println("Webhook ID:", WebhookId)
	err := CreateWebhookDb(body, WebhookId)
	if err != nil {
		fmt.Println("Error while inserting data into webhooks table", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error while inserting data into webhooks table",
			"details": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Webhook created successfully",
		"webhook": fiber.Map{
			"id":         WebhookId,
			"name":       body.Name,
			"url":        body.URL,
			"status":     "active",
			"created_at": time.Now(),
		},
	})
}

func CreateWebhookDb(webhook models.Webhook, WebhookId string) error {
	sqlStatement := `INSERT INTO webhooks (id, name, url, status, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := Database.DBConn.Exec(context.Background(), sqlStatement, WebhookId, webhook.Name, webhook.URL, "active", time.Now())
	if err != nil {
		fmt.Println("Error while inserting data into webhooks table", err)
		return err
	}
	return nil
}
