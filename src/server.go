package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("fomo", fiber.Map{
			"Title": "FOMO, App!",
		})
	})

	log.Info("Server is running on port 8080 🚀")
	app.Listen(":8080")
}
