package auth

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "auth api login"})
}

func Register(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "auth api register"})
}
