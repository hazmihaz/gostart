package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"

	"github.com/hazmihaz/gostart/internal/user"
)

func main() {
	fmt.Println("Starting Server")

	app := fiber.New()

	app.Settings.ErrorHandler = errorHandler

	app.Use(middleware.Recover())

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", handler)
	v1.Get("/hello", helloHandler)

	user.RegisterHandlers(v1)

	log.Fatal(app.Listen(3300))
}

func handler(c *fiber.Ctx) {
	c.Send("Go Rest Starter API V1")
}

func errorHandler(c *fiber.Ctx, err error) {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Return HTTP response
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	c.Status(code).SendString(err.Error())
}

func helloHandler(c *fiber.Ctx) {
	c.Send("Hello World!")
}
