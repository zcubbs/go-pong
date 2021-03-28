package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Pong 👋!",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
