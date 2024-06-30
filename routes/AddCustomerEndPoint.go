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

func AddCustomerEndPoint(c *fiber.Ctx) error {
	var body models.Endpoint
	_ = c.BodyParser(&body)

	CustomerEndpointId := utils.UUIDGenerator("CE-")
	fmt.Println("CustomerEndpointId ID:", CustomerEndpointId)
	err := CreateEndpointDb(body, CustomerEndpointId)
	if err != nil {
		fmt.Println("Error while inserting data into endpoints table", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error while inserting data into endpoints table",
			"details": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Endpoint Mapped successfully",
		"endpoint": fiber.Map{
			"id":         CustomerEndpointId,
			"webhook_id": body.WebhookID,
			"url":        body.URL,
			"status":     "active",
			"created_at": time.Now(),
		},
	})

}

func CreateEndpointDb(endpoint models.Endpoint, CustomerEndpointId string) error {
	fmt.Println("CustomerEndpointId ID:", CustomerEndpointId)
	sqlStatement := `INSERT INTO endpoints (id, webhook_id, status, url, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := Database.DBConn.Exec(context.Background(), sqlStatement, CustomerEndpointId, endpoint.WebhookID, "active", endpoint.URL, time.Now())
	if err != nil {
		fmt.Println("Error while inserting data into endpoints table", err)
		return err
	}
	return nil
}
