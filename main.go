package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	engine := html.New("./views", ".html")
	// If you want other engine, just replace with following
	// Create a new engine with django
	// engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	app.Get("/buttonClicked", func(c *fiber.Ctx) error {
		return c.Render("other", fiber.Map{})
	})

	log.Fatal(app.Listen(":3000"))
}
