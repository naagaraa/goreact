package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Exported function Main (needs to be uppercase to be exported)
func Routes() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/help", func(c *fiber.Ctx) error {
		return c.SendString("help pages")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about pages")
	})

	// Define your routes
	app.Get("/api/v1/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello from Go Fiber!"})
	})

	app.Listen(":3000")
}
