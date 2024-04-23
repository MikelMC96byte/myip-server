package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Add server header to all responses
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Server", "myip/server")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// It returns the IP address of the request as a JSON object
		// {"ip": "127.0.0.1"}

		return c.JSON(fiber.Map{"ip": c.IP()})
	})

	app.Listen(":3000")
}
