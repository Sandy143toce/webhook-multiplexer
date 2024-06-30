package setup

import (
	"github.com/Sandy143toce/webhook-multiplexer/routes"
	"github.com/Sandy143toce/webhook-multiplexer/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/webhook-multiplexer")
	api.Get("/metrics", monitor.New(monitor.Config{Title: "webhook_multiplexer API Metrics"}))

	// Register the routes
	api.Post("/create-webhook",
		utils.AuthorizeRequest("CreateWebhookRequest"),
		routes.CreateWebhook)
	api.Post("/add-customer-endpoint",
		utils.AuthorizeRequest("AddCustomerEndpointRequest"),
		routes.AddCustomerEndPoint)
	api.Post("/send-event",
		utils.AuthorizeRequest("SendEventRequest"),
		routes.SendEvent)
}
