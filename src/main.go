package main

import (
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/oschwald/geoip2-golang"
)

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

	app.Get("/geo", func(c *fiber.Ctx) error {
		db, err := geoip2.Open("/data/GeoLite2-City.mmdb")
		if err != nil {
			return c.JSON(fiber.Map{"error": err.Error()})
		}
		defer db.Close()

		record, err := db.City(net.ParseIP(c.IP()))

		if err != nil {
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(record)
	})

	app.Listen(":3000")
}
