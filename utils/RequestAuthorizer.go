package utils

import (
	"context"
	"net/url"

	Database "github.com/Sandy143toce/webhook-multiplexer/database"
	"github.com/Sandy143toce/webhook-multiplexer/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthorizeRequest(structName string) fiber.Handler {
	// Authorize the request
	return func(c *fiber.Ctx) error {
		errBody, valid := RequestValidation(c, structName)
		if !valid {
			return c.Status(fiber.StatusBadRequest).JSON(errBody)
		}
		return c.Next()
	}
}

func RequestValidation(c *fiber.Ctx, structName string) (models.ErrorResponse, bool) {

	switch structName {
	case "CreateWebhookRequest":
		var reqBody models.Webhook
		if err := c.BodyParser(&reqBody); err != nil {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: err.Error(),
				Details: "Invalid request body",
			}, false
		}
		if reqBody.URL == "" {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "URL is required",
			}, false
		}
		if reqBody.Name == "" {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "Name is required",
			}, false
		}
		if !isValidURL(reqBody.URL) {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "Invalid URL",
			}, false
		}
		if WebhookURLAlreadyExists(reqBody.URL, Database.DBConn) {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "Webhook URL already exists",
			}, false
		}
		return models.ErrorResponse{}, true
	case "AddCustomerEndpointRequest":
		var reqBody models.Endpoint
		if err := c.BodyParser(&reqBody); err != nil {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: err.Error(),
				Details: "Invalid request body",
			}, false
		}
		if reqBody.URL == "" {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "URL is required",
			}, false
		}
		if !isValidURL(reqBody.URL) {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "Invalid URL",
			}, false
		}
		if reqBody.WebhookID == "" {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "Webhook ID is required",
			}, false
		}
		if !isValidWebhookID(reqBody.WebhookID, Database.DBConn) {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "Invalid Webhook ID",
			}, false
		}
		if isEndpointAlreadyMapped(&reqBody, Database.DBConn) {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "Endpoint is already mapped",
			}, false
		}
		return models.ErrorResponse{}, true
	case "SendEventRequest":
		var reqBody models.SendEvent
		if err := c.BodyParser(&reqBody); err != nil {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: err.Error(),
				Details: "Invalid request body",
			}, false
		}
		if reqBody.EventName == "" {
			return models.ErrorResponse{
				Code:    400,
				Key:     "invalid_payload",
				Message: "invalid_payload",
				Details: "EventName is required",
			}, false
		}
		return models.ErrorResponse{}, true
	}
	return models.ErrorResponse{}, true
}

func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func isValidWebhookID(webhookID string, DBConn *pgxpool.Pool) bool {
	// Check if the webhook ID is a valid UUID
	sqlStatement := `SELECT id FROM webhooks WHERE id = $1`
	var id string
	err := DBConn.QueryRow(context.Background(), sqlStatement, webhookID).Scan(&id)
	if err == nil {
		return true
	} else {
		return false
	}
}

func isEndpointAlreadyMapped(body *models.Endpoint, DBConn *pgxpool.Pool) bool {
	// Check if the webhook ID is a valid UUID
	sqlStatement := `SELECT id FROM endpoints WHERE webhook_id = $1 AND url = $2`
	var id string
	err := DBConn.QueryRow(context.Background(), sqlStatement, body.WebhookID, body.URL).Scan(&id)
	if err == nil {
		return true
	} else {
		return false
	}
}

func WebhookURLAlreadyExists(url string, DBConn *pgxpool.Pool) bool {
	sqlStatement := `SELECT id FROM webhooks WHERE url = $1`
	var id string
	err := DBConn.QueryRow(context.Background(), sqlStatement, url).Scan(&id)
	if err == nil {
		return true
	} else {
		return false
	}
}
