package api

import "github.com/gofiber/fiber/v2"

// Exported function Main (needs to be uppercase to be exported)
func Routes() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/help", func(c *fiber.Ctx) error {
		return c.SendString("help pages")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about pages")
	})

	app.Listen(":3000")
}
