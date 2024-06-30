package main

import (
	dbConn "github.com/Sandy143toce/webhook-multiplexer/database"
	setup "github.com/Sandy143toce/webhook-multiplexer/setup"
	utils "github.com/Sandy143toce/webhook-multiplexer/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog/log"
	// other imports...
)

func main() {
	// Initialize the database connection pool
	if err := initPg(); err != nil {
		panic(err)
	}
	app := fiber.New()

	app.Use(requestid.New())
	app.Use(cors.New())
	// Default middleware config
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression, // 1
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setup.SetupRoutes(app)

	port := "8000"
	if port == "" {
		port = "3000" // Provide a default port if one isn't supplied
	}
	err := app.Listen(":" + port)
	if err != nil {
		log.Error().Str("Error", err.Error())
	}
	log.Info().Msgf("App listening to %s", port)
}

func initPg() error {
	var err error
	dbConn.DBConn, err = utils.InitDB()
	if err != nil {
		return err
	}

	return nil
}
