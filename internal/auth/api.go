package auth

import (
	"log"

	"github.com/gofiber/fiber"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(g fiber.Router) {
	g.Post("/auth/login", login())
}

func login() fiber.Handler {
	return func(c *fiber.Ctx) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		c.BodyParser(&req)
		log.Println(req)

		c.JSON(&req)
	}
}
