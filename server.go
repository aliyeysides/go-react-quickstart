// Go React Quickstart is a starter using a fiber backend and React frontend.
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Static("/", "client/dist")

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Listen(":3000")
}
